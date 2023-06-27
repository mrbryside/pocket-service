package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type EchoServer struct {
	handler Handler
}

func NewEchoServer(h Handler) EchoServer {
	return EchoServer{handler: h}
}
func (ec EchoServer) Start() {
	e := echo.New()

	// route
	e.POST("/pockets", ec.handler.CreatePockets)
	e.POST("/pockets/transactions", ec.handler.CreateTransaction)
	e.GET("/pockets", ec.handler.GetById)

	e.Use(middleware.Logger())

	e.Logger.Fatal(e.Start(":8080"))
}
