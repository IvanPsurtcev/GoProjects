package http

import (
	"context"
	"encoding/json"
	"metamask-auth/internal/game"
	"metamask-auth/internal/request"
	"metamask-auth/pkg/fhttp"

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

type requestHandlers struct {
	logger    *zap.Logger
	requestUC request.UseCase
}

func NewHandlers(logger *zap.Logger, requestUC request.UseCase) request.Handlers {
	return &requestHandlers{
		logger:    logger,
		requestUC: requestUC,
	}
}

func (h *requestHandlers) CreateRequest() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		var params request.CreateRequest
		if err := json.Unmarshal(ctx.PostBody(), &params); err != nil {
			h.logger.Error("CreateRequest err:", zap.Error(err))
			fhttp.ErrProcessRequest(ctx)
			return
		}
		res, err := h.requestUC.CreateRequest(context.TODO(), &params)
		switch err {
		case request.ErrReqAlreadyExists:
			ctx.Error(request.ErrReqAlreadyExists.Error(), fasthttp.StatusBadRequest)
		case nil:
			fhttp.RenderJson(ctx, fasthttp.StatusCreated, res)
		default:
			fhttp.ErrProcessRequest(ctx)
		}
	}
}

func (h *requestHandlers) CreateRandRequest() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		var params request.CreateRequest
		if err := json.Unmarshal(ctx.PostBody(), &params); err != nil {
			h.logger.Error("CreateGame err:", zap.Error(err))
			fhttp.ErrProcessRequest(ctx)
			return
		}
		res, err := h.requestUC.CreateRandRequest(context.TODO(), &params)
		switch err {
		case game.ErrGameAlreadyExists:
			ctx.Error(game.ErrGameAlreadyExists.Error(), fasthttp.StatusBadRequest)
		case nil:
			fhttp.RenderJson(ctx, fasthttp.StatusCreated, res)
		default:
			fhttp.ErrProcessRequest(ctx)
		}
	}
}

func (h *requestHandlers) GetUserActiveInvites() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		userAddress := ctx.UserValue("userAddress").(string)
		params := request.GetInvitesRequest{
			Address: userAddress,
		}
		res, err := h.requestUC.GetUserActiveInvites(context.TODO(), &params)
		switch err {
		case request.ErrUserNotFound:
			ctx.Error(request.ErrUserNotFound.Error(), fasthttp.StatusNotFound)
		case nil:
			fhttp.RenderJson(ctx, fasthttp.StatusCreated, res)
		default:
			fhttp.ErrProcessRequest(ctx)
		}
	}
}

func (h *requestHandlers) DeclineRequest() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		var params request.DeclineRequest
		if err := json.Unmarshal(ctx.PostBody(), &params); err != nil {
			h.logger.Error("DeclineRequest err:", zap.Error(err))
			fhttp.ErrProcessRequest(ctx)
			return
		}
		res, err := h.requestUC.DeclineRequest(context.TODO(), &params)
		if err != nil {
			fhttp.ErrProcessRequest(ctx)
			return
		}
		switch err {
		case request.ErrNotFound:
			ctx.Error(request.ErrNotFound.Error(), fasthttp.StatusNotFound)
		case nil:
			fhttp.RenderJson(ctx, fasthttp.StatusOK, res)
		default:
			fhttp.ErrProcessRequest(ctx)
		}
	}
}

func (h *requestHandlers) AcceptRequest() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		var params request.AcceptRequest
		if err := json.Unmarshal(ctx.PostBody(), &params); err != nil {
			h.logger.Error("AcceptRequest err:", zap.Error(err))
			fhttp.ErrProcessRequest(ctx)
			return
		}
		res, err := h.requestUC.AcceptRequest(context.TODO(), &params)
		if err != nil {
			fhttp.ErrProcessRequest(ctx)
			return
		}
		switch err {
		case request.ErrNotFound:
			ctx.Error(request.ErrNotFound.Error(), fasthttp.StatusNotFound)
		case nil:
			fhttp.RenderJson(ctx, fasthttp.StatusOK, res)
		default:
			fhttp.ErrProcessRequest(ctx)
		}
	}
}
