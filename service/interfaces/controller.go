package interfaces

import (
	"fmt"
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
		PrintAddError("name and deadline are required")
		return nil
	}

	// deadline /　priority must be numeric
	dl, err := strconv.Atoi(deadline)
	if err != nil {
		PrintAddError(err.Error())
		return nil
	}

	pri := 1
	if priority != "" {
		pri, err = strconv.Atoi(priority)
		if err != nil {
			PrintAddError(err.Error())
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
		PrintDeleteError("taskID is required")
		return nil
	}

	i, err := strconv.Atoi(id)
	if err != nil {
		PrintDeleteError(err.Error())
		return nil
	}

	if err = con.usecase.Delete(i); err != nil {
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

// staus priority deadline name
func (con *Controller) Change(c *cli.Context) error {
	id := c.Args().Get(0)
	column := c.Args().Get(1)
	data := c.Args().Get(2)

	i, err := strconv.Atoi(id)
	if err != nil {
		PrintChangeError(err.Error())
		return nil
	}

	if column == "" || data == "" {
		PrintChangeError("column abd data are required")
		return nil
	}

	if err := con.usecase.Change(i, column, data); err != nil {
		//log
		fmt.Print(fmt.Sprintf("faild to change a task: %s", err.Error()))
		return err
	}

	fmt.Println(fmt.Sprintf("Chnage task(id: %d ) %s=%s ", i, column, data))
	return nil
}

func PrintAddError(message string) {
	fmt.Println(fmt.Sprintf("ERROR : %s", message))
	fmt.Println("-----------------------------------")
	fmt.Println("usage : task add [name] [deadline] [priority]")
	fmt.Println("        REQUIRED [name] : name a task")
	fmt.Println("        REQUIRED [deadline] : due in X days")
	fmt.Println("        OPTIONAL [priority] : 0:low, 1:normal(default), 2:high")
}

func PrintDeleteError(message string) {
	fmt.Println(fmt.Sprintf("ERROR : %s", message))
	fmt.Println("-----------------------------------")
	fmt.Println("usage : task add [name] [deadline] [priority]")
	fmt.Println("        REQUIRED [name] : name a task")
	fmt.Println("        REQUIRED [deadline] : due in X days")
	fmt.Println("        OPTIONAL [priority] : 0:low, 1:normal(default), 2:high")
}

func PrintChangeError(message string) {
	fmt.Println(fmt.Sprintf("ERROR : %s", message))
	fmt.Println("-----------------------------------")
	fmt.Println("usage : task add [name] [deadline] [priority]")
	fmt.Println("        REQUIRED [name] : name a task")
	fmt.Println("        REQUIRED [deadline] : due in X days")
	fmt.Println("        OPTIONAL [priority] : 0:low, 1:normal(default), 2:high")
}
