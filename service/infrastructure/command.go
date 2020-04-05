package infrastructure

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/yukinooz/go_task/service/interfaces"
	"os"
)

type handler struct {
	controller *interfaces.Controller
}

var h *handler

// NewHandler aa
func NewHandler(cnt *interfaces.Controller) {
	h = &handler{
		controller: cnt,
	}
}

func Action() {
	//todo 関数にする
	app := &cli.App{
		Name:                 "task",
		Usage:                "manage your tasks",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:    "list",
				Usage:   "list all the uncompleted tasks",
				Action:  list,
			},
			{
				Name:    "add",
				Usage:   "add a task",
				Action:  add,
			},
			{
				Name:    "change",
				Usage:   "change status of the task",
				Action:  change,
			},
			{
				Name:    "delete",
				Usage:   "add a task",
				Action:  remove,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

func list(c *cli.Context) error {
	return h.controller.List()
}

func add(c *cli.Context) error {
	return h.controller.Add(c)
}

func remove(c *cli.Context) error {
	return h.controller.Delete(c)
}

func change(c *cli.Context) error {
	return h.controller.Change(c)
}
