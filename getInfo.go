package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func getOsInfo() map[string]string {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		fmt.Println("Error opening a file.", err)
		return nil
	}

	os := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), "=")
		os[text[0]] = text[1]
	}
	return os
}

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

func getProcInfo() map[string]string {
	data, err := os.Open("/proc/cpuinfo")
	if err != nil {
		fmt.Println("Error opening a file", err)
		return nil
	}

	proc := make(map[string]string)
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		text := strings.Split(strings.ReplaceAll(scanner.Text(), " ", ""), ":")
		proc[text[0]] = text[1]
	}
	return proc
}

func getDiskInfo() [][]string {
	cmd := exec.Command("df")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing command", err)
		return nil
	}
	outputSlice := strings.Split(string(output), "\n")
	outputSlice = outputSlice[:len(outputSlice) - 1]
	disk := make([][]string, len(outputSlice))
	for i, val := range outputSlice {
		arr := strings.Split(val, " ")
		disk[i] = arr
	}
	for _, x := range disk {
		fmt.Println(x)
	}
	return disk
}