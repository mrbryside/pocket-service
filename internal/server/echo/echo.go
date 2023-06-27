package echo

import (
	"github.com/labstack/echo/v4"
)

type EchoServer struct {
	handler Handler
}

func NewEchoServer(h Handler) EchoServer {
	return EchoServer{handler: h}
}
func (ec EchoServer) Start() {
	e := echo.New()
	e.POST("/pockets", ec.handler.CreatePockets)
	e.POST("/pockets/transactions", ec.handler.CreateTransaction)

	e.Logger.Fatal(e.Start(":8080"))
}
