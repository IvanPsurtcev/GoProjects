package game

import "errors"

var (
	ErrGameNotFound      = errors.New("game with such id is not found")
	ErrGameAlreadyExists = errors.New("game already exists")
	ErrPlayerNotReady    = errors.New("player is not ready")
	ErrInvalidNft        = errors.New("invalid nft")
)
