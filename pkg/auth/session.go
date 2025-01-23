package auth

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"

	"github.com/anthdm/hollywood/actor"

	"goblin/pkg/protocol"
)

type session struct {
	log  *slog.Logger
	conn net.Conn
}

func newSession(conn net.Conn) actor.Producer {
	return func() actor.Receiver {
		return &session{
			log: slog.Default().With(
				slog.String("actor", "auth_session"),
				slog.String("addr", conn.RemoteAddr().String()),
			),
			conn: conn,
		}
	}
}

func (s *session) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case actor.Initialized:

	case actor.Started:
		go s.readLoop(c)

	case actor.Stopped:
		s.conn.Close()

	default:
		s.log.Warn("unknown msg", "msg", msg)
	}
}

// const AuthHandler table[] =
// {
//     { AUTH_LOGON_CHALLENGE,     STATUS_CONNECTED, &AuthSocket::_HandleLogonChallenge    },
//     { AUTH_LOGON_PROOF,         STATUS_CONNECTED, &AuthSocket::_HandleLogonProof        },
//     { AUTH_RECONNECT_CHALLENGE, STATUS_CONNECTED, &AuthSocket::_HandleReconnectChallenge},
//     { AUTH_RECONNECT_PROOF,     STATUS_CONNECTED, &AuthSocket::_HandleReconnectProof    },
//     { REALM_LIST,               STATUS_AUTHED,    &AuthSocket::_HandleRealmList         },

// };

func (s *session) readLoop(c *actor.Context) {
	// TODO: MAX_AUTH_LOGON_CHALLENGES_IN_A_ROW 3
	buf := make([]byte, 4096)
	for {
		n, err := s.conn.Read(buf)
		if err != nil {
			if err == io.EOF || errors.Is(err, net.ErrClosed) {
				break
			}
			s.log.Error("Failed to read connection", slog.Any("err", err))
			break
		}

		// copy shared buffer, to prevent race conditions.
		raw := make([]byte, n)
		copy(raw, buf[:n])

		var opcode protocol.AuthCmd
		_, err = binary.Decode(raw[:n], binary.LittleEndian, &opcode)
		if err != nil {
			s.log.Error("Failed to decode message",
				"opcode", fmt.Sprintf("0x%02x", opcode),
				"raw", fmt.Sprintf("0x%02x", raw))
		}

		switch opcode {
		case protocol.AuthCmdChallenge:
			msg := protocol.C_MSGAuthChallenge{}
			msg.Read(raw)
			c.Send(c.PID(), msg)
		default:
			s.log.Error("Unknown opcode",
				"opcode", fmt.Sprintf("0x%02x", opcode))
		}
	}

	s.log.Debug("Client has disconnected")
	c.Engine().Poison(c.PID())
}
