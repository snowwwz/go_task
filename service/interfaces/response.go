package interfaces

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

func showError(message, cmd string) error {
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

func showSuccess(message string) error {
	_, err := fmt.Println(message)
	if err != nil {
		return err
	}
	return nil
}

func printAddUsage() error {
	fmt.Println("-----------------------------------")
	fmt.Println("usage : go_task add [name] [deadline] [priority]")
	fmt.Println("        [name]* : name a task")
	fmt.Println("        [deadline]* : due in X days")
	fmt.Println("        [priority] : 0:low, 1:normal(default), 2:high")
	return nil
}

func printDeleteUsage() error {
	fmt.Println("-----------------------------------")
	fmt.Println("usage : go_task delete [id]")
	fmt.Println("        [id]* : taskID")
	return nil
}

func printChangeUsage() error {
	fmt.Println("-----------------------------------")
	fmt.Println("usage : go_task change [id] [column] [data]")
	fmt.Println("        [id]* : taskID")
	fmt.Println("        [column]* : name/statis/priority/deadline")
	fmt.Println("        [data]* : new data")
	return nil
}

func showList(result [][]string) error {
	fmt.Println(fmt.Sprintf("( ¨̮ )　you have %d tasks", len(result)))
	fmt.Println("")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Status", "Priority", "Deadline", "Created",})
	table.AppendBulk(result)
	table.Render()
	return nil
}

func showJournal(result [][]string) error {
	fmt.Println(fmt.Sprintf("( ¨̮ )　%d tasks have been changed today", len(result)))
	fmt.Println("")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Number", "Name", "Status",})
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)
	table.AppendBulk(result)
	table.Render()
	return nil
}