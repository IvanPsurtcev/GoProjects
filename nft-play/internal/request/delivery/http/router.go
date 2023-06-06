package http

import (
	"metamask-auth/internal/middleware"
	"metamask-auth/internal/request"
	"metamask-auth/pkg/fhttp"

	"github.com/fasthttp/router"
)

func InitRouter(router *router.Group, h request.Handlers, mw middleware.MwManager) {
	mwHandler := mw.JwtValidate()
	router.POST("/create", fhttp.CombineHandlers(mwHandler, h.CreateRequest()))
	router.POST("/random", fhttp.CombineHandlers(mwHandler, h.CreateRandRequest()))
	router.GET("/invites", fhttp.CombineHandlers(mwHandler, h.GetUserActiveInvites()))
	router.PATCH("/decline", fhttp.CombineHandlers(mwHandler, h.DeclineRequest()))
	router.PATCH("/accept", fhttp.CombineHandlers(mwHandler, h.AcceptRequest()))
}
