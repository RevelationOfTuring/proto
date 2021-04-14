package types

import "time"

const (
	OIP10 CoinType = 0
	OIP20 CoinType = 1
)

type CoinType int32

type TokenJSON struct {
	Symbol  string
	Decimal int32
	Infos   []TokenInfoJSON
	IcoTime time.Time
}

type TokenInfoJSON struct {
	TotalSupply int64
	Type        CoinType
}

type TokenListJSON struct {
	Tokens []TokenJSON
}
