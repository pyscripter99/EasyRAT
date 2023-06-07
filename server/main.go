package main

import (
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
	e.Logger.Fatal(e.Start(":" + fmt.Sprint(port)))
}
