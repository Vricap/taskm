package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func printTable(task []process) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"USER", "PID", "CPU", "MEMORY", "PROGRAM"})
	fmt.Printf("Total: %v\n", len(task))
	for _, val := range task[:25] {
		row := []string{val.USER, fmt.Sprint(val.PID), fmt.Sprintf("%.f%%", val.CPU), fmt.Sprintf("%.f mb", val.MEM), val.PROGRAM}
		table.Append(row)
	}
	table.Render()
}