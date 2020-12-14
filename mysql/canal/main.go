package main

import (
	"fmt"

	"github.com/siddontang/go-mysql/canal"
)

type MyEventHandler struct {
	canal.DummyEventHandler
}

func (h *MyEventHandler) OnRow(e *canal.RowsEvent) error {
	fmt.Printf("%s %v\n", e.Action, e.Rows)
	return nil
}

func (h *MyEventHandler) String() string {
	return "MyEventHandler"
}
func main() {
	cfg := canal.NewDefaultConfig()
	cfg.Addr = "10.0.0.2:3306"
	cfg.User = "canal"
	cfg.Password = "canal"
	// We only care table canal_test in test db
	cfg.Dump.TableDB = "test"
	cfg.Dump.Tables = []string{"canal_test"}

	c, err := canal.NewCanal(cfg)
	if err != nil {
		panic(err)
	}

	// Register a handler to handle RowsEvent
	c.SetEventHandler(&MyEventHandler{})

	// Start canal
	if err := c.Run(); err != nil {
		panic(err)
	}
}
