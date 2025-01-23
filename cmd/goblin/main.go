package main

import (
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/anthdm/hollywood/actor"
	"github.com/lmittmann/tint"

	"goblin/pkg/auth"
	"goblin/pkg/realm"
)

func init() {
	// set global logger with custom options
	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stderr, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	))
}

func main() {
	// TODO: load config file and others

	// some intro
	banner()

	// actors engine
	e, err := actor.NewEngine(actor.NewEngineConfig())
	if err != nil {
		panic(err)
	}

	// very simple game server setup
	authPID := e.Spawn(auth.NewActor(":3724"), "auth")
	realmPID := e.Spawn(realm.NewActor(), "realm")

	// graceful shutdown
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGINT, syscall.SIGTERM)
	<-sigch

	// wait till the server is gracefully shutdown by using a WaitGroup in the Poison call.
	wg := &sync.WaitGroup{}
	e.Poison(authPID, wg)
	e.Poison(realmPID, wg)
	wg.Wait()

	slog.Info("Halting process...")
}
