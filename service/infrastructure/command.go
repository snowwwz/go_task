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

// NewHandler create a new handler
func NewHandler(cnt *interfaces.Controller) {
	h = &handler{
		controller: cnt,
	}
}

// Run commands
func Run() {
	app := &cli.App{
		Name:                 "go_task",
		Usage:                "manage your tasks",
		EnableBashCompletion: true,
	}
	app.UseShortOptionHandling = true
	app.Commands = []*cli.Command{
		{
			Name:  "list",
			Usage: "list all the uncompleted tasks",
			Flags: []cli.Flag{
				&cli.BoolFlag{Name: "all", Aliases: []string{"a"}},
			},
			Action: list,
		},
		{
			Name:  "add",
			Usage: "add a task",
			Flags: []cli.Flag{
				&cli.IntFlag{Name: "name", Aliases: []string{"n"}},
				&cli.StringFlag{Name: "due", Aliases: []string{"d"}},
				&cli.StringFlag{Name: "priority", Aliases: []string{"p"}},
			},
			Action: add,
		},
		{
			Name:  "change",
			Usage: "change status of the task",
			Flags: []cli.Flag{
				&cli.IntFlag{Name: "id"},
				&cli.StringFlag{Name: "target", Aliases: []string{"t"}},
				&cli.StringFlag{Name: "data", Aliases: []string{"d"}},
			},
			Action: change,
		},
		{
			Name:  "delete",
			Usage: "add a task",
			Flags: []cli.Flag{
				&cli.IntFlag{Name: "id"},
			},
			Action: remove,
		},
		{
			Name:   "journal",
			Usage:  "list tasks which status have been changed",
			Action: journal,
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

func list(c *cli.Context) error {
	isAll := c.Bool("all")
	return h.controller.List(isAll)
}

func add(c *cli.Context) error {
	name := c.String("name")
	deadline := c.Int("due")
	priority := c.Int("priority")
	return h.controller.Add(name, deadline, priority)
}

func remove(c *cli.Context) error {
	id := c.Int("id")
	return h.controller.Delete(id)
}

func change(c *cli.Context) error {
	id := c.Int("id")
	target := c.String("target")
	data := c.String("data")
	return h.controller.Change(id, target, data)
}

func journal(c *cli.Context) error {
	return h.controller.Journal()
}
