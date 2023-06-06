package game

import "github.com/valyala/fasthttp"

type Handlers interface {
	CreateGame() fasthttp.RequestHandler
	CreateRandGame() fasthttp.RequestHandler
	GetGames() fasthttp.RequestHandler
	GetGame() fasthttp.RequestHandler
	GetUserNfts() fasthttp.RequestHandler
	SetApplicantBet() fasthttp.RequestHandler
	SetReceiverBet() fasthttp.RequestHandler
	SetApplicantStatus() fasthttp.RequestHandler
	SetReceiverStatus() fasthttp.RequestHandler
	CancelGame() fasthttp.RequestHandler
	GameDone() fasthttp.RequestHandler
	RunGame() fasthttp.RequestHandler
}
