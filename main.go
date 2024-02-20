package main

import (
	"os"
	"time"
)

func main() {	
	if len(os.Args) > 1 && os.Args[1] == "-v" {
		printInfoTable()
		return
	}

	InitTask()

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	go func ()  {
		for range ticker.C {
			InitTask()
		}
	}()

	select {}
}