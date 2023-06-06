package reservoir

type TokenInfo struct {
	Contract string `json:"contract"`
	TokenId  string `json:"tokenId"`
	Image    string `json:"image"`
}

type ApiResponseFloorPrice struct {
	Tokens map[string]float64 `json:"tokens"`
}

type UserToken struct {
	Token TokenInfo `json:"token"`
}

type ApiResponseTokens struct {
	Tokens []UserToken `json:"tokens"`
}
