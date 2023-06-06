package usecase

import (
	"context"
	"go.uber.org/zap"
	"math"
	"math/rand"
	"metamask-auth/internal/auth"
	"metamask-auth/internal/contract"
	"metamask-auth/internal/game"
	"metamask-auth/pkg/reservoir"
	"strings"
)

type gameUseCase struct {
	logger     *zap.Logger
	gameRepo   game.Repository
	authRepo   auth.Repository
	reservoir  reservoir.Client
	contractUC contract.UseCase
}

func NewUseCase(
	logger *zap.Logger,
	gameRepo game.Repository,
	reservoir reservoir.Client,
	contractUC contract.UseCase,
	authRepo auth.Repository) game.UseCase {
	return &gameUseCase{
		logger:     logger,
		gameRepo:   gameRepo,
		authRepo:   authRepo,
		reservoir:  reservoir,
		contractUC: contractUC,
	}
}

func (u *gameUseCase) CreateGame(ctx context.Context, params *game.CreateGameRequest) (*game.CreateGameResponse, error) {
	gameId, err := u.contractUC.CreateGame(ctx)
	if err != nil {
		return nil, err
	}
	err = u.gameRepo.CreateGame(ctx, gameId, params.ApplicantAddr, params.ReceiverAddr)
	if err != nil {
		return nil, err
	}
	return &game.CreateGameResponse{
		Id: gameId,
	}, nil
}

func (u *gameUseCase) CreateRandGame(ctx context.Context, params *game.CreateGameRequest) (*game.CreateGameResponse, error) {
	randUser, err := u.authRepo.GetRandUser(ctx, params.ApplicantAddr)
	if err != nil {
		return nil, err
	}
	params.ReceiverAddr = randUser.Address
	return u.CreateGame(ctx, params)
}

func (u *gameUseCase) SetApplicantBet(ctx context.Context, params *game.SetBetRequest) error {
	if err := u.gameRepo.SetArg(ctx, params.GameId, ColApplicantNft, params.Contract+" "+params.Token); err != nil {
		return err
	}
	return nil
}

func (u *gameUseCase) SetReceiverBet(ctx context.Context, params *game.SetBetRequest) error {
	if err := u.gameRepo.SetArg(ctx, params.GameId, ColReceiverNft, params.Contract+" "+params.Token); err != nil {
		return err
	}
	return nil
}

func (u *gameUseCase) SetApplicantStatus(ctx context.Context, params *game.SetPlayerStatusRequest) error {
	if err := u.gameRepo.SetArg(ctx, params.GameId, ColApplicantReady, params.Status); err != nil {
		return err
	}
	return nil
}

func (u *gameUseCase) SetReceiverStatus(ctx context.Context, params *game.SetPlayerStatusRequest) error {
	if err := u.gameRepo.SetArg(ctx, params.GameId, ColReceiverReady, params.Status); err != nil {
		return err
	}
	return nil
}

func (u *gameUseCase) CancelGame(ctx context.Context, params *game.SetGameStatusRequest) error {
	if err := u.gameRepo.SetArg(ctx, params.GameId, ColStatus, StatusCanceled); err != nil {
		return err
	}
	return nil
}

func (u *gameUseCase) GameDone(ctx context.Context, params *game.SetGameStatusRequest) error {
	if err := u.gameRepo.SetArg(ctx, params.GameId, ColStatus, StatusDone); err != nil {
		return err
	}
	return nil
}

func (u *gameUseCase) GetGame(ctx context.Context, params *game.BaseGameRequest) (*game.GetGameResponse, error) {
	g, err := u.gameRepo.GetGame(ctx, params.Id)
	if err != nil {
		return nil, err
	}
	applicantNft, err := u.parseNftString(g.ApplicantNft)
	if err != nil {
		return nil, err
	}
	receiverNft, err := u.parseNftString(g.ReceiverNft)
	if err != nil {
		return nil, err
	}
	return &game.GetGameResponse{
		Id:             g.Id,
		ApplicantAddr:  g.ApplicantAddr,
		ReceiverAddr:   g.ReceiverAddr,
		WinnerAddr:     g.WinnerAddr,
		Status:         g.Status,
		ApplicantReady: g.ApplicantReady,
		ReceiverReady:  g.ReceiverReady,
		CreateDate:     g.CreateDate,
		ApplicantNft:   applicantNft,
		ReceiverNft:    receiverNft,
	}, nil
}

func (u *gameUseCase) GetUserNfts(ctx context.Context, params *game.GetUserNftsRequest) (*game.GetUserNftsResponse, error) {
	res, err := u.reservoir.GetUserNFTs(params.Address)
	if err != nil {
		u.logger.Error("GetUserNfts reservoir err", zap.Error(err))
		return nil, err
	}
	return &game.GetUserNftsResponse{
		Nfts: res,
	}, nil
}

func (u *gameUseCase) RunGame(ctx context.Context, params *game.BaseGameRequest) (*game.GetGameResponse, error) {
	gameObj, err := u.GetGame(ctx, &game.BaseGameRequest{Id: params.Id})
	if err != nil {
		return nil, err
	}
	if !gameObj.ReceiverReady || !gameObj.ApplicantReady {
		return nil, game.ErrPlayerNotReady
	}
	if gameObj.ApplicantNft == nil ||
		gameObj.ReceiverNft == nil {
		return nil, game.ErrInvalidNft
	}

	var applicantProb float64
	if gameObj.ApplicantNft.Price == nil || gameObj.ReceiverNft.Price == nil {
		applicantProb = 0.5
	} else {
		wholePrice := *gameObj.ApplicantNft.Price + *gameObj.ReceiverNft.Price
		applicantProb = math.Max(*gameObj.ApplicantNft.Price/wholePrice, 0.1)
	}

	r := rand.Float64()
	var winner string
	if r > applicantProb {
		winner = gameObj.ReceiverAddr

	} else {
		winner = gameObj.ApplicantAddr
	}
	if err = u.gameRepo.SetWinner(ctx, gameObj.Id, winner); err != nil {
		return nil, err
	}
	if err = u.contractUC.SetWinner(ctx, &contract.WinnerGame{ID: gameObj.Id, Receiver: winner}); err != nil {
		return nil, err
	}
	gameObj.Status = 2
	gameObj.WinnerAddr = &winner
	return gameObj, nil
}

func (u *gameUseCase) parseNftString(nftString *string) (*game.Nft, error) {
	if nftString == nil {
		return nil, nil
	}
	nftRaw := strings.Split(*nftString, " ")
	var nftPrice *float64
	if len(nftRaw) == 2 {
		res, err := u.reservoir.GetFloorPriceNFT(&reservoir.TokenInfo{Contract: nftRaw[0], TokenId: nftRaw[1]})
		if err == nil {
			nftPrice = &res
		}
	} else {
		return nil, nil
	}
	res := game.Nft{
		Contract: nftRaw[0],
		TokenId:  nftRaw[1],
		Price:    nftPrice,
	}
	return &res, nil
}

func (u *gameUseCase) GetGames(ctx context.Context, params *game.GetGamesRequest) (*game.GetGamesResponse, error) {
	var activeGames []game.GetGameResponse
	games, err := u.gameRepo.GetGames(ctx, params.Address, params.Status)
	if err != nil {
		return nil, err
	}
	for _, activeGame := range games {
		g, err := u.gameRepo.GetGame(ctx, activeGame.Id)
		if err != nil {
			return nil, err
		}
		applicantNft, err := u.parseNftString(g.ApplicantNft)
		if err != nil {
			return nil, err
		}
		receiverNft, err := u.parseNftString(g.ReceiverNft)
		if err != nil {
			return nil, err
		}

		currentGame := game.GetGameResponse{
			Id:             g.Id,
			ApplicantAddr:  g.ApplicantAddr,
			ReceiverAddr:   g.ReceiverAddr,
			WinnerAddr:     g.WinnerAddr,
			Status:         g.Status,
			ApplicantReady: g.ApplicantReady,
			ReceiverReady:  g.ReceiverReady,
			CreateDate:     g.CreateDate,
			ApplicantNft:   applicantNft,
			ReceiverNft:    receiverNft,
		}
		activeGames = append(activeGames, currentGame)
	}
	return &game.GetGamesResponse{
		Games: activeGames,
	}, err
}
