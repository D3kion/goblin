package auth

import (
	"log/slog"
	"net"

	"github.com/anthdm/hollywood/actor"
)

func NewActor(listenAddr string) actor.Producer {
	return func() actor.Receiver {
		return &authActor{
			log:        slog.Default().With("actor", "auth"),
			listenAddr: listenAddr,
		}
	}
}

type authActor struct {
	log *slog.Logger

	listenAddr string
	ln         net.Listener
}

func (ac *authActor) Receive(c *actor.Context) {
	switch c.Message().(type) {
	case actor.Initialized:
		ln, err := net.Listen("tcp", ac.listenAddr)
		if err != nil {
			panic(err)
		}
		ac.ln = ln

	case actor.Started:
		ac.log.Info("Server started", slog.String("addr", ac.listenAddr))
		go ac.acceptLoop(c)

	case actor.Stopped:
		ac.log.Info("Server stopped")
	}
}

func (ac *authActor) acceptLoop(c *actor.Context) {
	for {
		conn, err := ac.ln.Accept()
		if err != nil {
			ac.log.Error("Failed to accept connection", slog.Any("err", err))
			break
		}
		ac.log.Debug("Accepted new client",
			slog.String("addr", conn.RemoteAddr().String()))

		c.SpawnChild(
			newSession(conn), "session",
			actor.WithID(conn.RemoteAddr().String()),
		)
	}
}
