package main

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"proto/performance/types/gogoprotobuf"
	"proto/performance/types/golang"
	"proto/performance/types/protobuf"

	"github.com/tendermint/go-amino"
	"google.golang.org/protobuf/proto"
)

const (
	decimal      = 18
	symbol1      = "okt"
	symbol2      = "okb"
	totalsupply1 = 10240000
	totalsupply2 = 20480000
)

var (
	tokenListGo        golang.TokenList
	tokenListProto     protobuf.TokenList
	tokenListGogoProto gogoprotobuf.TokenList
)

func init() {
	tokenListGo = golang.TokenList{
		Tokens: []*golang.Token{
			{
				Symbol:  symbol1,
				Decimal: decimal,
				Infos: []*golang.TokenInfo{
					{
						TotalSupply: totalsupply1,
						CoinType:    golang.OIP10,
					},
				},
			},
			{
				Symbol:  symbol2,
				Decimal: decimal,
				Infos: []*golang.TokenInfo{
					{
						TotalSupply: totalsupply2,
						CoinType:    golang.OIP20,
					},
				},
			},
		},
	}

	tokenListProto = protobuf.TokenList{
		Tokens: []*protobuf.Token{
			{
				Symbol:  symbol1,
				Decimal: decimal,
				Infos: []*protobuf.Token_TokenInfo{
					{
						TotalSupply: totalsupply1,
						CoinType:    protobuf.Token_OIP10,
					},
				},
			},
			{
				Symbol:  symbol2,
				Decimal: decimal,
				Infos: []*protobuf.Token_TokenInfo{
					{
						TotalSupply: totalsupply2,
						CoinType:    protobuf.Token_OIP20,
					},
				},
			},
		},
	}

	tokenListGogoProto = gogoprotobuf.TokenList{
		Tokens: []*gogoprotobuf.Token{
			{
				Symbol:  symbol1,
				Decimal: decimal,
				Infos: []*gogoprotobuf.Token_TokenInfo{
					{
						TotalSupply: totalsupply1,
						CoinType:    gogoprotobuf.OIP10,
					},
				},
			},
			{
				Symbol:  symbol2,
				Decimal: decimal,
				Infos: []*gogoprotobuf.Token_TokenInfo{
					{
						TotalSupply: totalsupply2,
						CoinType:    gogoprotobuf.OIP20,
					},
				},
			},
		},
	}
}

func TestLength(t *testing.T) {
	// 1. JSON
	rawBytes, err := json.Marshal(tokenListGo)
	if err != nil {
		log.Panic(err)
	}
	formatPrinter("json", rawBytes)

	// 2. amino
	aminoCdc := amino.NewCodec()
	rawBytes, err = aminoCdc.MarshalBinaryLengthPrefixed(tokenListGo)
	if err != nil {
		log.Panic(err)
	}
	formatPrinter("amino", rawBytes)

	// 3. goprotobuf
	rawBytes, err = proto.Marshal(&tokenListProto)
	if err != nil {
		log.Panic(err)
	}
	formatPrinter("goprotobuf", rawBytes)

	// 4. gogoprotobuf
	rawBytes, err = tokenListGogoProto.Marshal()
	if err != nil {
		log.Panic(err)
	}
	formatPrinter("gogoprotobuf", rawBytes)
}

func formatPrinter(kind string, rawBytes []byte) {
	fmt.Printf(`================================================================================================================================================

%s:
	length: 	%d
	encoded:	%s
`, kind, len(rawBytes), string(rawBytes))
}
