package realm

import (
	"log/slog"

	"github.com/anthdm/hollywood/actor"
)

func NewActor() actor.Producer {
	return func() actor.Receiver {
		return &realmServer{}
	}
}

type realmServer struct {
}

func (as *realmServer) Receive(c *actor.Context) {
	switch c.Message().(type) {
	case actor.Initialized:

	case actor.Started:
		slog.Info("Server started")

	case actor.Stopped:
		slog.Info("Server stopped")
	}
}
