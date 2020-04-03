package infrastructure

import (
	"fmt"
	"os"

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
		Commands: []*cli.Command{
			//
			// 	Name:  "add",
			// 	Usage: "add a task",
			// 	Action: h.controller.Add(*cli.Context),
			// },
			{
				Name:   "list",
				Usage:  "list all the uncompleted tasks",
				Action: list,
			},
			// {
			// 	Name:  "delete",
			// 	Usage: "delete a task",
			// 	Action: h.controller.Add(*cli.Context),
			// },
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}

func list(c *cli.Context) error {
	// fmt.Fprintf(c.App.Writer, ":wave: over here, eh\n")
	return h.controller.List()
}
