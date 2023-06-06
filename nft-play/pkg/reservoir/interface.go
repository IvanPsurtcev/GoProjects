package reservoir

type Client interface {
	GetFloorPriceNFT(tokenInfo *TokenInfo) (float64, error)
	GetUserNFTs(address string) ([]TokenInfo, error)
}
