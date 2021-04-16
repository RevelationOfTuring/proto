package golang

type CoinType int32

const (
	OIP10 CoinType = 0
	OIP20 CoinType = 1
)

type Token struct {
	Symbol  string       `json:"symbol,omitempty"`
	Decimal int32        `json:"decimal,omitempty"`
	Infos   []*TokenInfo `json:"infos,omitempty"`
}

type TokenInfo struct {
	TotalSupply int64    `json:"total_supply,omitempty"`
	CoinType    CoinType `json:"coin_type,omitempty"`
}

type TokenList struct {
	Tokens []*Token `json:"tokens,omitempty"`
}
