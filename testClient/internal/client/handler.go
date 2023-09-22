package client

import "github.com/valyala/fasthttp"

type Handlers interface {
	ProcessItems() fasthttp.RequestHandler
	GetLimits() fasthttp.RequestHandler
}
