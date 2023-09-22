package http

import (
	"github.com/fasthttp/router"
	"testClient/internal/client"
)

func InitRouter(router *router.Group, h client.Handlers) {
	router.POST("/process", h.ProcessItems())
	router.GET("/limits", h.GetLimits())
}
