package daemon

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/SDur/ops-planner/db"
	"github.com/SDur/ops-planner/model"
	"github.com/SDur/ops-planner/ui"
)

type Config struct {
	ListenSpec string

	Db db.Config
	UI ui.Config
}

func Run(cfg *Config) error {
	log.Printf("Starting, HTTP on: %s\n", cfg.ListenSpec)

	db, err := db.InitDb(cfg.Db)
	if err != nil {
		log.Printf("Error initializing database: %v\n", err)
		return err
	}

	m := model.New(db)

	ui.Start(cfg.UI, m)

	waitForSignal()

	return nil
}

func waitForSignal() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	s := <-ch
	log.Printf("Got signal: %v, exiting.", s)
}
