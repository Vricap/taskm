package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func printInfoTable() {
	fmt.Println(`
***********
* OS INFO *
***********`)
	makeOsTable()
	fmt.Println()

	fmt.Println(`
************
* RAM INFO *
************`)
	makeRamTable()
	fmt.Println()

	fmt.Println(`
******************
* PROCESSOR INFO *
******************`)
	makeProcTable()
	fmt.Println()

	fmt.Println(`
*************
* DISK INFO *
*************`)
	makeDiskTable()
	fmt.Println()

	fmt.Println(`
****************
* NETWORK INFO *
****************`)
	makeNetTable()
	fmt.Println()

	fmt.Println(`
****************
* OTHER INFO *
****************`)
	makeOtherTable()
	fmt.Println()
}

func makeOsTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Build", "URL", "Documentation"})
	os := getOsInfo()
	row := []string{os["NAME"], os["BUILD_ID"], os["HOME_URL"], os["DOCUMENTATION_URL"]}
	table.Append(row)
	table.Render()
}

func makeRamTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Total", "Used", "Free", "Available"})
	ram := getRamInfo()
	row := []string{ram["Total"]+" mb", ram["Used"]+" mb", ram["Free"]+" mb", ram["Available"]+" mb"}
	table.Append(row)
	table.Render()
}

func makeProcTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Vendor", "Model", "Cores", "Clock Speed", "Cache Size"})
	proc := getProcInfo()
	row := []string{proc["vendor_id"], proc["model name"], proc["cpu cores"], proc["cpu MHz"], proc["cache size"]}
	table.Append(row)
	table.Render()
}

func makeDiskTable() {
	disk := getDiskInfo()
	table := tablewriter.NewWriter(os.Stdout)
	header := []string{}
	for i, row := range disk {
		for j, y := range row {
			if i == 0 && j != len(row) - 1 {
				header = append(header, y)
				continue
			}
		} 
			if i != 0 {
				table.Append(row)
			}
		}
		table.SetHeader(header)
	table.Render()
}

func makeNetTable() {
	table := tablewriter.NewWriter(os.Stdout)
	net := getNetInfo()
	if net["type"] == "" {
		table.SetHeader([]string{"You didn't connected to any network."})
		table.Render()
		return
	}
	if net["type"] == "Wifi" {
		fmt.Println("You're connected to an wireless network (wifi) with such configurations:")
		} else if net["type"] == "Ethernet" {
		fmt.Println("You're connected to an wired network (ethernet) with such configurations:")
	}
	table.SetHeader([]string{"Type", "Interface", "IP", "Broadcast", "Netmask"})
	row := []string{net["type"], net["interface"], net["ip"], net["broadcast"], net["netmask"]}
	table.Append(row)
	table.Render()
}

func makeOtherTable() {
	fmt.Println("Here are some other info about your computer.")
	logUser := getLoginUser()
	upTime := getUptime()
	gpu := getGpuInfo()
	
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Log In User", "Up Time", "GPU"})
	row := []string{
		fmt.Sprintf("You're %s login in %s on %s", logUser["user"], logUser["tty"], logUser["time"]), 
		fmt.Sprintf("Hours: %s Minutes: %s Seconds: %s", upTime["hours"], upTime["minutes"], upTime["seconds"]),
		gpu,
	}
	table.Append(row)
	table.Render()
}