package http

import (
	"context"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"testClient/internal/client"
	"testClient/pkg/fhttp"
)

type clientHandlers struct {
	logger   *zap.Logger
	clientUC client.UseCase
}

func NewHandlers(logger *zap.Logger, clientUC client.UseCase) client.Handlers {
	return &clientHandlers{
		logger:   logger,
		clientUC: clientUC,
	}
}

func (h *clientHandlers) ProcessItems() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		var params client.BatchRequest
		if err := json.Unmarshal(ctx.PostBody(), &params); err != nil {
			h.logger.Error("Process err:", zap.Error(err))
			fhttp.ErrProcessRequest(ctx)
			return
		}
		err := h.clientUC.ProcessItems(context.TODO(), &params)
		switch err {
		case nil:
			ctx.SetStatusCode(fasthttp.StatusCreated)
		default:
			fhttp.ErrProcessRequest(ctx)
		}
	}
}

func (h *clientHandlers) GetLimits() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		res, err := h.clientUC.GetLimits(context.TODO())
		switch err {
		case nil:
			fhttp.RenderJson(ctx, fasthttp.StatusOK, res)
		default:
			fhttp.ErrProcessRequest(ctx)
		}
	}
}
