package main

import (
	"os/exec"
	"strings"
)

// type Memory struct {
// 	Total string
// 	Used string
// 	Free string
// 	Available string
// }

func getRamInfo() map[string]string {
	output, err := exec.Command("free", "-m").CombinedOutput()
	if err != nil {
		return nil
	}
	outputSlice := strings.Split(string(output), "\n")
	
	field := strings.Fields(outputSlice[1])
	memory := map[string]string{
		"Total": field[1],
		"Used": field[2],
		"Free": field[3],
		"Available": field[6],
	}
	return memory
}