package http

import (
	"github.com/fasthttp/router"
	"metamask-auth/internal/game"
	"metamask-auth/internal/middleware"
	"metamask-auth/pkg/fhttp"
)

func InitRouter(router *router.Group, h game.Handlers, mw middleware.MwManager) {
	mwHandler := mw.JwtValidate()
	router.POST("/create", fhttp.CombineHandlers(mwHandler, h.CreateGame()))
	router.POST("/random", fhttp.CombineHandlers(mwHandler, h.CreateRandGame()))
	router.GET("/game", fhttp.CombineHandlers(mwHandler, h.GetGame()))
	router.GET("/nfts", fhttp.CombineHandlers(mwHandler, h.GetUserNfts()))
	router.GET("/games", fhttp.CombineHandlers(mwHandler, h.GetGames()))
	router.PATCH("/applicant_bet", fhttp.CombineHandlers(mwHandler, h.SetApplicantBet()))
	router.PATCH("/receiver_bet", fhttp.CombineHandlers(mwHandler, h.SetReceiverBet()))
	router.PATCH("/applicant_status", fhttp.CombineHandlers(mwHandler, h.SetApplicantStatus()))
	router.PATCH("/receiver_status", fhttp.CombineHandlers(mwHandler, h.SetReceiverStatus()))
	router.PATCH("/cancel", fhttp.CombineHandlers(mwHandler, h.CancelGame()))
	router.PATCH("/done", fhttp.CombineHandlers(mwHandler, h.GameDone()))
	router.POST("/run", fhttp.CombineHandlers(mwHandler, h.RunGame()))
}
