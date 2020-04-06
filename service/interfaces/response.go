package interfaces

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

func returnError(message, cmd string) error {
	_, err := fmt.Println(fmt.Sprintf("ERROR: %s", message))
	if err != nil {
		return err
	}

	// print usages
	switch cmd {
	case "add":
		return printAddUsage()
	case "delete":
		return printDeleteUsage()
	case "change":
		return printChangeUsage()
	}
	return nil
}

func returnSuccess(message string) error {
	_, err := fmt.Println(message)
	if err != nil {
		return err
	}
	return nil
}

func printAddUsage() error {
	fmt.Println("-----------------------------------")
	fmt.Println("usage : task add [name] [deadline] [priority]")
	fmt.Println("        [name]* : name a task")
	fmt.Println("        [deadline]* : due in X days")
	fmt.Println("        [priority] : 0:low, 1:normal(default), 2:high")
	return nil
}

func printDeleteUsage() error {
	fmt.Println("-----------------------------------")
	fmt.Println("usage : task delete [id]")
	fmt.Println("        [id]* : taskID")
	return nil
}

func printChangeUsage() error {
	fmt.Println("-----------------------------------")
	fmt.Println("usage : task change [id] [column] [data]")
	fmt.Println("        [id]* : taskID")
	fmt.Println("        [column]* : name/statis/priority/deadline")
	fmt.Println("        [priority] : 0:low, 1:normal(default), 2:high")
	return nil
}

func returnList(result [][]string) error {
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
