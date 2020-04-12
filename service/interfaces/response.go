package interfaces

import (
	"errors"
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
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

func showGraph() error {
	if err := ui.Init(); err != nil {
		return errors.New(fmt.Sprintf("failed to initialize termui: %v", err))
	}
	defer ui.Close()

	p := widgets.NewParagraph()
	p.Title = "Text Box"
	p.Text = "PRESS q TO QUIT DEMO"
	p.SetRect(0, 2, 40, 3)
	p.TextStyle.Fg = ui.ColorWhite
	p.BorderStyle.Fg = ui.ColorWhite

	g := widgets.NewGauge()
	g.Title = "Slim Gauge"
	g.SetRect(0, 6, 40, 9)
	g.Percent = 60
	g.BarColor = ui.ColorYellow
	g.LabelStyle = ui.NewStyle(ui.ColorBlue)
	g.BorderStyle.Fg = ui.ColorWhite

	l := widgets.NewList()
	l.Title = "List"
	l.Rows = []string{
		"[0] github.com/gizak/termui/v3",
		"[1] [你好，世界](fg:blue)",
		"[2] [こんにちは世界](fg:red)",
		"[3] [color](fg:white,bg:green) output",
	}
	l.TextStyle = ui.NewStyle(ui.ColorYellow)
	l.WrapText = false
	l.SetRect(43, 0, 80, 10)

	ui.Render(p,g,l)

	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
	return nil
}
