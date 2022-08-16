package binance

import (
	api "github.com/pvzzle/crypto-connectors/binance/exchangeAPI"
)

// Public API

// FetchServerTime
func (c *connector) FetchServerTime() (*api.ServerTime, error) {
	res, fetchError := c.publicFetch("GET", "/api/v3/time", "")
	if fetchError != nil {
		return nil, fetchError
	}

	j := &api.ServerTime{}
	if parseError := parseBody(res, j); parseError != nil {
		return nil, parseError
	}

	return j, nil
}

// FetchExchangeInfo
func (c *connector) FetchExchangeInfo() (*api.ExchangeInformation, error) {
	// TODO: query parameters set
	res, fetchError := c.publicFetch("GET", "/api/v3/exchangeInfo", "")
	if fetchError != nil {
		return nil, fetchError
	}

	j := &api.ExchangeInformation{}
	if parseError := parseBody(res, j); parseError != nil {
		return nil, parseError
	}

	return j, nil
}

// Private API
