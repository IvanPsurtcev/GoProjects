package contracts

import (
	"math/big"
	"time"
)

type CollectionCreatedEvent struct {
	Event      string    `json:"event_type"`
	Collection string    `json:"collection"`
	Name       string    `json:"name"`
	Symbol     string    `json:"symbol"`
	Timestamp  time.Time `json:"timestamp"`
}

type TokenMintedEvent struct {
	Event      string    `json:"event_type"`
	Collection string    `json:"collection"`
	Recipient  string    `json:"recipient"`
	TokenId    *big.Int  `json:"tokenId"`
	TokenURI   string    `json:"tokenURI"`
	Timestamp  time.Time `json:"timestamp"`
}
