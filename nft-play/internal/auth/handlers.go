package auth

import "github.com/valyala/fasthttp"

type Handlers interface {
	Authorize() fasthttp.RequestHandler
	GetUsers() fasthttp.RequestHandler
	GetUser() fasthttp.RequestHandler
}
