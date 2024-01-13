package main

import (
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

func InitTask() []process {
	cmd := exec.Command("ps", "aux")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil
	}
	info := parseProcOut(string(output))
	sort.SliceStable(info, func(i, j int) bool {
		return info[i].MEM > info[j].MEM
	}) 
	// for _, val := range info {
	// 	fmt.Println(val)
	// }
	return info
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