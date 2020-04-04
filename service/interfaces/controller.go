package interfaces

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/urfave/cli/v2"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/yukinooz/go_task/service/usecase"
)

// Controller struct
type Controller struct {
	usecase *usecase.Usecase
}

// NewController create a new controller
func NewController(u *usecase.Usecase) *Controller {
	return &Controller{
		usecase: u,
	}
}

// Add a task
func (con *Controller) Add(c *cli.Context) error {
	name := c.Args().Get(0)
	deadline := c.Args().Get(1)
	priority := c.Args().Get(2)

	// required args
	if name == "" || deadline == "" {
		PrintAddUsage("name and deadline are required")
		return nil
	}

	// deadline /　priority must be numeric
	dl, err := strconv.Atoi(deadline)
	if err != nil {
		PrintAddUsage(err.Error())
		return nil
	}

	pri := 1
	if priority != "" {
		pri, err = strconv.Atoi(priority)
		if err != nil {
			PrintAddUsage(err.Error())
			return nil
		}
	}

	if err = con.usecase.Add(name, pri, dl); err != nil {
		fmt.Println(fmt.Sprintf("faild to add a task: %s", err.Error()))
		return err
	}
	fmt.Println(fmt.Sprintf("Added a task \"%s\"", name))
	return nil
}

// Delete delete a task
func (con *Controller) Delete(c *cli.Context) error {
	id := c.Args().Get(0)

	if id == "" {
		PrintAddUsage("taskID is required")
		return nil
	}
	spew.Dump(id)

	i, err := strconv.Atoi(id)
	if err != nil {
		PrintAddUsage(err.Error())
		return nil
	}

	err = con.usecase.Delete(i)
	if err != nil {
		fmt.Println(fmt.Sprintf("faild to delete a task: %s", err.Error()))
		return err
	}
	fmt.Println("Deleted a task")
	return nil
}

// List lisy tasks
func (con *Controller) List() error {
	result, err := con.usecase.List()
	if err != nil {
		fmt.Print(fmt.Sprintf("faild to add a task: %s", err.Error()))
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"ID", "Name", "Status", "Priority", "Deadline", "Created",
	})

	for _, v := range result {
		table.Append(v)
	}
	table.Render()
	return nil
}

func PrintAddUsage(message string) {
	fmt.Println(fmt.Sprintf("ERROR : %s", message))
	fmt.Println("-----------------------------------")
	fmt.Println("usage : task add [name] [deadline] [priority]")
	fmt.Println("        REQUIRED [name] : name a task")
	fmt.Println("        REQUIRED [deadline] : due in X days")
	fmt.Println("        OPTIONAL [priority] : 0:low, 1:normal(default), 2:high")
}
