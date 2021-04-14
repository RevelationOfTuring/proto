package main

import (
	"fmt"
	"io/ioutil"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"proto/protobuf/types"
)

func main() {
	// define
	curTime := timestamppb.Now()
	token1 := &types.Token{
		Symbol:  "okt",
		Decimal: 18,
		Infos: []*types.Token_TokenInfo{
			{
				TotalSupply: 1024,
				CoinType:    types.Token_OIP10,
			},
			{
				TotalSupply: 2048,
				CoinType:    types.Token_OIP20,
			},
		},
		IcoTime: curTime,
	}

	token2 := &types.Token{
		Symbol:  "okb",
		Decimal: 18,
		Infos: []*types.Token_TokenInfo{
			{
				TotalSupply: 4096,
				CoinType:    types.Token_OIP20,
			},
		},
		IcoTime: curTime,
	}

	tokenList := &types.TokenList{
		Tokens: []*types.Token{
			token1, token2,
		},
	}

	bytes, err := proto.Marshal(tokenList)
	if err != nil {
		panic(err)
	}

	// length
	fmt.Println("protobuf3 length:", len(bytes))

	// write file
	if err = ioutil.WriteFile("encoded_proto3.txt", bytes, 0644); err != nil {
		panic(err)
	}
}
