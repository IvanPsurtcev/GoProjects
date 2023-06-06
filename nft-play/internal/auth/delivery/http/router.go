package http

import (
	"github.com/fasthttp/router"
	"metamask-auth/internal/auth"
	"metamask-auth/internal/middleware"
	"metamask-auth/pkg/fhttp"
)

func InitRouter(router *router.Group, h auth.Handlers, mw middleware.MwManager) {
	mwHandler := mw.JwtValidate()
	router.POST("/auth", h.Authorize())
	router.GET("/users", fhttp.CombineHandlers(mwHandler, h.GetUsers()))
	router.GET("/user", fhttp.CombineHandlers(mwHandler, h.GetUser()))
}
