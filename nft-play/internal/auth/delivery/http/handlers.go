package http

import (
	"context"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"metamask-auth/internal/auth"
	"metamask-auth/pkg/fhttp"
	"strconv"
)

type authHandlers struct {
	logger *zap.Logger
	authUC auth.UseCase
}

func NewHandlers(logger *zap.Logger, authUC auth.UseCase) auth.Handlers {
	return &authHandlers{
		logger: logger,
		authUC: authUC,
	}
}

func (h *authHandlers) Authorize() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		var params auth.AuthorizeRequest
		if err := json.Unmarshal(ctx.PostBody(), &params); err != nil {
			h.logger.Error("Authorize err:", zap.Error(err))
			fhttp.ErrProcessRequest(ctx)
			return
		}
		res, err := h.authUC.Authorize(context.TODO(), &params)
		if err != nil {
			fhttp.ErrProcessRequest(ctx)
			return
		}
		fhttp.RenderJson(ctx, fasthttp.StatusOK, res)
	}
}

func (h *authHandlers) GetUsers() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		offset, err := strconv.ParseInt(string(ctx.QueryArgs().Peek("offset")), 10, 64)
		if err != nil {
			offset = 0
		}
		limit, err := strconv.ParseInt(string(ctx.QueryArgs().Peek("limit")), 10, 64)
		if err != nil {
			limit = -1
		}
		params := auth.GetUsersRequest{
			Offset: offset,
			Limit:  limit,
		}
		res, err := h.authUC.GetUsers(context.TODO(), &params)
		if err != nil {
			fhttp.ErrProcessRequest(ctx)
			return
		}
		fhttp.RenderJson(ctx, fasthttp.StatusOK, res)
	}
}

func (h *authHandlers) GetUser() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		userAddress := ctx.UserValue("userAddress").(string)
		params := auth.GetUserRequest{
			Address: userAddress,
		}
		res, err := h.authUC.GetUser(context.TODO(), &params)
		if err != nil {
			fhttp.ErrProcessRequest(ctx)
			return
		}
		fhttp.RenderJson(ctx, fasthttp.StatusOK, res)
	}
}
