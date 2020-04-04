package infrastructure

import (
	"fmt"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
	"github.com/yukinooz/go_task/service/interfaces"
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
	app := &cli.App{
		EnableBashCompletion: true,
		Name:                 "task",
		Usage:                "manage your tasks",
		Commands: []*cli.Command{
			{
				Name:    "list",
				Usage:   "list all the uncompleted tasks",
				Aliases: []string{"l"},
				Action:  list,
			},
			{
				Name:    "add",
				Usage:   "add a task",
				Aliases: []string{"a"},
				Action:  add,
			},
			{
				Name:    "delete",
				Usage:   "add a task",
				Aliases: []string{"a"},
				Action:  remove,
			},
		},
	}

	sort.Sort(cli.CommandsByName(app.Commands))
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
