package game

import (
	"metamask-auth/pkg/reservoir"
	"time"
)

type Nft struct {
	TokenId  string   `json:"token_id"`
	Contract string   `json:"contract"`
	Price    *float64 `json:"price"`
}

type CreateGameRequest struct {
	ApplicantAddr string `json:"applicant_addr"`
	ReceiverAddr  string `json:"receiver_addr"`
}

type CreateGameResponse struct {
	Id int64 `json:"id"`
}

type SetBetRequest struct {
	GameId   int64  `json:"id"`
	Contract string `json:"contract"`
	Token    string `json:"token"`
}

type SetPlayerStatusRequest struct {
	GameId int64 `json:"id"`
	Status bool  `json:"status"`
}

type SetGameStatusRequest struct {
	GameId int64 `json:"id"`
	Status int64 `json:"status"`
}

type BaseGameRequest struct {
	Id int64 `json:"id"`
}

type GetGamesRequest struct {
	Status  int64
	Address string
}

type GetGamesResponse struct {
	Games []GetGameResponse `json:"games"`
}

type GetUserNftsRequest struct {
	Address string
}

type GetUserNftsResponse struct {
	Nfts []reservoir.TokenInfo `json:"nfts"`
}

type Game struct {
	Id             int64     `db:"id"`
	ApplicantAddr  string    `db:"applicant_addr"`
	ReceiverAddr   string    `db:"receiver_addr"`
	WinnerAddr     *string   `db:"winner_addr"`
	Status         int64     `db:"status"`
	ApplicantReady bool      `db:"applicant_ready"`
	ReceiverReady  bool      `db:"receiver_ready"`
	CreateDate     time.Time `db:"create_date"`
	ApplicantNft   *string   `db:"applicant_nft"`
	ReceiverNft    *string   `db:"receiver_nft"`
}

type GetGameResponse struct {
	Id             int64     `json:"id"`
	ApplicantAddr  string    `json:"applicant_addr"`
	ReceiverAddr   string    `json:"receiver_addr"`
	WinnerAddr     *string   `json:"winner_addr"`
	Status         int64     `json:"status"`
	ApplicantReady bool      `json:"applicant_ready"`
	ReceiverReady  bool      `json:"receiver_ready"`
	CreateDate     time.Time `json:"create_data"`
	ApplicantNft   *Nft      `json:"applicant_nft"`
	ReceiverNft    *Nft      `json:"receiver_nft"`
}
