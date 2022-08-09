package binance

import (
	api "github.com/pvzzle/crypto-connectors/binance/exchangeAPI"
)

// Public API

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


// Private API