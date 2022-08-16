package exchangeAPI

type Dummy struct {
}

type ServerTime struct {
	ServerTime uint64 `json:"serverTime"`
}

type SymbolPriceTicker struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type ExchangeInformation struct {
	Timezone   string `json:"timezone"`
	ServerTime uint64 `json:"serverTime"`
	// TODO: RateLimits
	// TODO: ExchangeFilters
	Symbols []ExchangeInformationSymbol `json:"symbols"`
}

type ExchangeInformationSymbol struct {
	Symbol                     string        `json:"symbol"`
	Status                     string        `json:"status"`
	BaseAsset                  string        `json:"baseAsset"`
	BaseAssetPrecision         int           `json:"baseAssetPrecision"`
	QuoteAsset                 string        `json:"quoteAsset"`
	QuotePrecision             int           `json:"quotePrecision"`
	QuoteAssetPrecision        int           `json:"quoteAssetPrecision"`
	OrderTypes                 []string      `json:"orderTypes"`
	IcebergAllowed             bool          `json:"icebergAllowed"`
	OcoAllowed                 bool          `json:"ocoAllowed"`
	QuoteOrderQtyMarketAllowed bool          `json:"quoteOrderQtyMarketAllowed"`
	AllowTrailingStop          bool          `json:"allowTrailingStop"`
	CancelReplaceAllowed       bool          `json:"cancelReplaceAllowed"`
	IsSpotTradingAllowed       bool          `json:"isSpotTradingAllowed"`
	IsMarginTradingAllowed     bool          `json:"isMarginTradingAllowed"`
	// TODO: Filters
	Permissions                []string      `json:"permissions"`
}
