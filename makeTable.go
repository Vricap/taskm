package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func makeTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"USER", "PID", "CPU", "MEMORY", "PROGRAM"})
	taskManager := InitTask()
	fmt.Printf("Total: %v\n", len(taskManager))
	for _, val := range taskManager[:50] {
		row := []string{val.USER, fmt.Sprint(val.PID), fmt.Sprintf("%.f%%", val.CPU), fmt.Sprintf("%.f mb", val.MEM), val.PROGRAM}
		table.Append(row)
	}
	table.Render()
}