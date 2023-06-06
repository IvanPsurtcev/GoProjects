package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"metamask-auth/internal/game"
	"metamask-auth/pkg/fhttp"
	"strconv"
)

type gameHandlers struct {
	logger *zap.Logger
	gameUC game.UseCase
}

func NewHandlers(logger *zap.Logger, gameUC game.UseCase) game.Handlers {
	return &gameHandlers{
		logger: logger,
		gameUC: gameUC,
	}
}

func (h *gameHandlers) CreateGame() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		var params game.CreateGameRequest
		if err := json.Unmarshal(ctx.PostBody(), &params); err != nil {
			h.logger.Error("CreateGame err:", zap.Error(err))
			fhttp.ErrProcessRequest(ctx)
			return
		}
		res, err := h.gameUC.CreateGame(context.TODO(), &params)
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

func (h *gameHandlers) CreateRandGame() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		var params game.CreateGameRequest
		if err := json.Unmarshal(ctx.PostBody(), &params); err != nil {
			h.logger.Error("CreateGame err:", zap.Error(err))
			fhttp.ErrProcessRequest(ctx)
			return
		}
		res, err := h.gameUC.CreateRandGame(context.TODO(), &params)
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

func (h *gameHandlers) SetApplicantBet() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		var params game.SetBetRequest
		if err := json.Unmarshal(ctx.PostBody(), &params); err != nil {
			h.logger.Error("SetApplicantBet err:", zap.Error(err))
			fhttp.ErrProcessRequest(ctx)
			return
		}
		err := h.gameUC.SetApplicantBet(context.TODO(), &params)
		switch err {
		case game.ErrGameNotFound:
			ctx.Error(game.ErrGameNotFound.Error(), fasthttp.StatusNotFound)
		case nil:
			ctx.SetStatusCode(fasthttp.StatusOK)
		default:
			fhttp.ErrProcessRequest(ctx)
		}
	}
}

func (h *gameHandlers) SetReceiverBet() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		var params game.SetBetRequest
		if err := json.Unmarshal(ctx.PostBody(), &params); err != nil {
			h.logger.Error("SetReceiverBet err:", zap.Error(err))
			fhttp.ErrProcessRequest(ctx)
			return
		}
		err := h.gameUC.SetReceiverBet(context.TODO(), &params)
		switch err {
		case game.ErrGameNotFound:
			ctx.Error(game.ErrGameNotFound.Error(), fasthttp.StatusNotFound)
		case nil:
			ctx.SetStatusCode(fasthttp.StatusOK)
		default:
			fhttp.ErrProcessRequest(ctx)
		}
	}
}

func (h *gameHandlers) SetApplicantStatus() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		var params game.SetPlayerStatusRequest
		if err := json.Unmarshal(ctx.PostBody(), &params); err != nil {
			h.logger.Error("SetApplicantStatus err:", zap.Error(err))
			fhttp.ErrProcessRequest(ctx)
			return
		}
		err := h.gameUC.SetApplicantStatus(context.TODO(), &params)
		switch err {
		case game.ErrGameNotFound:
			ctx.Error(game.ErrGameNotFound.Error(), fasthttp.StatusNotFound)
		case nil:
			ctx.SetStatusCode(fasthttp.StatusOK)
		default:
			fhttp.ErrProcessRequest(ctx)
		}
	}
}

func (h *gameHandlers) SetReceiverStatus() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		var params game.SetPlayerStatusRequest
		if err := json.Unmarshal(ctx.PostBody(), &params); err != nil {
			h.logger.Error("SetReceiverStatus err:", zap.Error(err))
			fhttp.ErrProcessRequest(ctx)
			return
		}
		err := h.gameUC.SetReceiverStatus(context.TODO(), &params)
		switch err {
		case game.ErrGameNotFound:
			ctx.Error(game.ErrGameNotFound.Error(), fasthttp.StatusNotFound)
		case nil:
			ctx.SetStatusCode(fasthttp.StatusOK)
		default:
			fhttp.ErrProcessRequest(ctx)
		}
	}
}

func (h *gameHandlers) RunGame() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		var params game.BaseGameRequest
		if err := json.Unmarshal(ctx.PostBody(), &params); err != nil {
			h.logger.Error("SetReceiverStatus err:", zap.Error(err))
			fhttp.ErrProcessRequest(ctx)
			return
		}
		res, err := h.gameUC.RunGame(context.TODO(), &params)
		switch err {
		case game.ErrGameNotFound:
			ctx.Error(game.ErrGameNotFound.Error(), fasthttp.StatusNotFound)
		case game.ErrPlayerNotReady:
			ctx.Error(game.ErrPlayerNotReady.Error(), fasthttp.StatusBadRequest)
		case game.ErrInvalidNft:
			ctx.Error(game.ErrInvalidNft.Error(), fasthttp.StatusBadRequest)
		case nil:
			fhttp.RenderJson(ctx, fasthttp.StatusOK, res)
		default:
			fhttp.ErrProcessRequest(ctx)
		}
	}
}

func (h *gameHandlers) CancelGame() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		var params game.SetGameStatusRequest
		if err := json.Unmarshal(ctx.PostBody(), &params); err != nil {
			h.logger.Error("CancelGame err:", zap.Error(err))
			fhttp.ErrProcessRequest(ctx)
			return
		}
		err := h.gameUC.CancelGame(context.TODO(), &params)
		switch err {
		case game.ErrGameNotFound:
			ctx.Error(game.ErrGameNotFound.Error(), fasthttp.StatusNotFound)
		case nil:
			ctx.SetStatusCode(fasthttp.StatusOK)
		default:
			fhttp.ErrProcessRequest(ctx)
		}
	}
}

func (h *gameHandlers) GameDone() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		var params game.SetGameStatusRequest
		if err := json.Unmarshal(ctx.PostBody(), &params); err != nil {
			h.logger.Error("GameDone err:", zap.Error(err))
			fhttp.ErrProcessRequest(ctx)
			return
		}
		err := h.gameUC.GameDone(context.TODO(), &params)
		switch err {
		case game.ErrGameNotFound:
			ctx.Error(game.ErrGameNotFound.Error(), fasthttp.StatusNotFound)
		case nil:
			ctx.SetStatusCode(fasthttp.StatusOK)
		default:
			fhttp.ErrProcessRequest(ctx)
		}
	}
}

func (h *gameHandlers) GetGame() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		gameId, err := strconv.ParseInt(string(ctx.QueryArgs().Peek("id")), 10, 64)
		if err != nil {
			h.logger.Error("GameDone err:", zap.Error(err))
			fhttp.ErrProcessRequest(ctx)
			return
		}
		res, err := h.gameUC.GetGame(context.TODO(), &game.BaseGameRequest{Id: gameId})
		switch err {
		case game.ErrGameNotFound:
			ctx.Error(game.ErrGameNotFound.Error(), fasthttp.StatusNotFound)
		case nil:
			fhttp.RenderJson(ctx, fasthttp.StatusOK, res)
		default:
			fhttp.ErrProcessRequest(ctx)
		}
	}
}

func (h *gameHandlers) GetUserNfts() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		userAddress := ctx.UserValue("userAddress").(string)
		fmt.Println(userAddress)
		res, err := h.gameUC.GetUserNfts(context.TODO(), &game.GetUserNftsRequest{Address: userAddress})
		switch err {
		case game.ErrGameNotFound:
			ctx.Error(game.ErrGameNotFound.Error(), fasthttp.StatusNotFound)
		case nil:
			fhttp.RenderJson(ctx, fasthttp.StatusOK, res)
		default:
			fhttp.ErrProcessRequest(ctx)
		}
	}
}

func (h *gameHandlers) GetGames() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		userAddress := ctx.UserValue("userAddress").(string)
		params := game.GetGamesRequest{
			Address: userAddress,
		}
		statusString := string(ctx.QueryArgs().Peek("status"))
		if statusString != "" {
			status, err := strconv.ParseInt(statusString, 10, 64)
			if err != nil {
				ctx.Error("invalid query param", fasthttp.StatusBadRequest)
				return
			}
			params.Status = status
		}
		res, err := h.gameUC.GetGames(context.TODO(), &params)
		switch err {
		case game.ErrGameNotFound:
			ctx.Error(game.ErrGameNotFound.Error(), fasthttp.StatusNotFound)
		case nil:
			fhttp.RenderJson(ctx, fasthttp.StatusOK, res)
		default:
			fhttp.ErrProcessRequest(ctx)
		}
	}
}
