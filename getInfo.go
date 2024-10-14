package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
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
		text := strings.Split(scanner.Text(), ":")
		proc[strings.TrimSpace(text[0])] = strings.TrimSpace(text[1])
	}
	return proc
}

func getDiskInfo() [][]string {
	cmd := exec.Command("df", "-h")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing command", err)
		return nil
	}
	outputSlice := strings.Split(string(output), "\n")
	outputSlice = outputSlice[:len(outputSlice) - 1]
	disk := make([][]string, len(outputSlice))
	for i, val := range outputSlice {
		regex := regexp.MustCompile(`\s+`)
		val = regex.ReplaceAllString(val, " ")
		arr := strings.Split(val, " ")
		disk[i] = arr
	}	
	return disk
}

func getNetInfo() map[string]string {
	// FIXME: somehow if you had docker installed, that is considered an connection
	cmd := exec.Command("ifconfig")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing command", err)
		return nil
	}

	net := strings.Split(string(output), "\n")
	y := []int{}
	for i, val := range net {
		if val == "" {
			y = append(y, i)
			if len(y) == 2 {
				break
			}
		}
	}
	wlNet := net[y[1] + 1:]
	enNet := net[:y[0]]

	netInfo := map[string]string{
		"type": "",
		"interface": "",
		"ip": "",
		"broadcast": "",
		"netmask": "",
	}

	if strings.Contains(enNet[1], "inet") {
		attribute := strings.Split(strings.TrimSpace(enNet[1]), " ")
		netInfo["type"] = "Ethernet"
		netInfo["interface"] = strings.Split(enNet[0], ":")[0]
		netInfo["ip"] = attribute[1]
		netInfo["broadcast"] = attribute[7]
		netInfo["netmask"] = attribute[4]
	} else if strings.Contains(wlNet[1], "inet") {
		attribute := strings.Split(strings.TrimSpace(wlNet[1]), " ")
		netInfo["type"] = "Wifi"
		netInfo["interface"] = strings.Split(wlNet[0], ":")[0]
		netInfo["ip"] = attribute[1]
		netInfo["broadcast"] = attribute[7]
		netInfo["netmask"] = attribute[4]
	}
	
	return netInfo
}

func getUptime() map[string]string {
	data, err := os.ReadFile("/proc/uptime")
	if err != nil {
		fmt.Println("Error reading a file", err)
		return nil
	}

	time, _ := strconv.ParseFloat(strings.Split(string(data), " ")[0], 64)
	return map[string]string{
		"hours": fmt.Sprint(int(time) / 3600),
		"minutes": fmt.Sprint(int(time) % 3600 / 60),
		"seconds": fmt.Sprint(int(time) % 60),
	}
}

func getLoginUser() map[string]string {
	cmd := exec.Command("who")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing command", err)
		return nil
	}	

	user := strings.Split(string(output), " ")
	return 	map[string]string{
		"user": user[0],
		"tty": user[3],
		"time": strings.TrimSpace(user[len(user) - 2]) + " " + strings.TrimSpace(user[len(user) - 1]),
	}
}

func getGpuInfo() string {
	cmd := exec.Command("sh", "-c", "lspci | grep -iE 'vga|3d|2d'")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing command", err)
		return ""
	}
	
	return strings.TrimSpace(string(output))
}
