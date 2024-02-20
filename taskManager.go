package main

import (
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

type process struct {
	USER string
	PID int
	CPU float64
	MEM float64
	PROGRAM string	
}

func InitTask() {
	// Clear the terminal first
	go clearTerminal()

	// Excecute the command to get processes
	cmd := exec.Command("ps", "aux")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return
	}

	// Parse the output from the command to process struct
	task := parseProcOut(string(output))

	// Sort the process asc based on memory usage
	sort.SliceStable(task, func(i, j int) bool {
		return task[i].MEM > task[j].MEM
	})

	// Print the table containing the processes
	printTmTable(task)
}

func parseProcOut(output string) []process {
	var processes []process
	proces := strings.Split(output, "\n")
	for i := 1; i < len(proces); i++ {
		field := strings.Fields(proces[i])
		if len(field) == 0 {
			continue
		}

		pid, _ := strconv.Atoi(field[1])
		cpuPercentage, _ := strconv.ParseFloat(field[2], 64)
		memory, _ := strconv.ParseFloat(field[5], 64)

		proc := process{
			USER: field[0],
			PID: pid,
			CPU: cpuPercentage,
			MEM: memory / 1024.0,
			PROGRAM: field[10],
		}
		processes = append(processes, proc)
	}
	return processes
}

func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}