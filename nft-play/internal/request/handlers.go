package request

import "github.com/valyala/fasthttp"

type Handlers interface {
	CreateRequest() fasthttp.RequestHandler
	GetUserActiveInvites() fasthttp.RequestHandler
	DeclineRequest() fasthttp.RequestHandler
	AcceptRequest() fasthttp.RequestHandler
	CreateRandRequest() fasthttp.RequestHandler
}
