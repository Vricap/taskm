package main

import (
	"time"
)

func main() {	
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