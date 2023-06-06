package reservoir

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

type client struct {
	fasthttp *fasthttp.Client
	apiUrl   string
	logger   *zap.Logger
}

func NewClient(logger *zap.Logger, apiUrl string) Client {
	c := &client{
		apiUrl: apiUrl,
		logger: logger,
	}
	c.fasthttp = &fasthttp.Client{}

	return c
}

func (c *client) GetFloorPriceNFT(tokenInfo *TokenInfo) (float64, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(c.apiUrl + "/tokens/floor/v1?contract=" + tokenInfo.Contract)

	err := fasthttp.Do(req, resp)
	if err != nil {
		return 0, fmt.Errorf("failed to execute request: %v", err)
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		return 0, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	var apiResponse ApiResponseFloorPrice
	err = json.Unmarshal(resp.Body(), &apiResponse)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	if len(apiResponse.Tokens) == 0 {
		return -1, fmt.Errorf("no tokens found in the response")
	}

	priceNFT, ok := apiResponse.Tokens[tokenInfo.TokenId]
	if !ok {
		return -1, fmt.Errorf("token ID %s not found in the response", tokenInfo.TokenId)
	}

	return priceNFT, nil
}

func (c *client) GetUserNFTs(address string) ([]TokenInfo, error) {

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	queryParams := url.Values{
		"normalizeRoyalties":    {"false"},
		"sortBy":                {"acquiredAt"},
		"sortDirection":         {"desc"},
		"limit":                 {"20"},
		"includeTopBid":         {"false"},
		"includeAttributes":     {"false"},
		"includeLastSale":       {"false"},
		"includeRawData":        {"false"},
		"useNonFlaggedFloorAsk": {"false"},
	}
	req.SetRequestURI(c.apiUrl + "/users/" + address + "/tokens/v7?" + queryParams.Encode())

	err := fasthttp.Do(req, resp)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode())
	}

	var userTokensResponse ApiResponseTokens
	err = json.Unmarshal(resp.Body(), &userTokensResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	var tokens []TokenInfo
	for _, userToken := range userTokensResponse.Tokens {
		tokens = append(tokens, userToken.Token)
	}

	return tokens, nil
}
