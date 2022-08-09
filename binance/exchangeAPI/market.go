package exchangeAPI


type Dummy struct {

}

type ServerTime struct {
	ServerTime uint64 `json:"serverTime"`
}

type SymbolPriceTicker struct {
	Symbol string `json:"symbol"`
	Price string `json:"price"`
}

