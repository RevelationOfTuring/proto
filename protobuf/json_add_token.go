package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"proto/protobuf/types"
)

func main() {
	// define
	curTime := time.Now()
	token1 := types.TokenJSON{
		Symbol:  "okt",
		Decimal: 18,
		Infos: []types.TokenInfoJSON{
			{
				TotalSupply: 1024,
				Type:        types.OIP10,
			},
			{
				TotalSupply: 2048,
				Type:        types.OIP20,
			},
		},
		IcoTime: curTime,
	}

	token2 := types.TokenJSON{
		Symbol:  "okb",
		Decimal: 18,
		Infos: []types.TokenInfoJSON{
			{
				TotalSupply: 4096,
				Type:        types.OIP20,
			},
		},
		IcoTime: curTime,
	}

	tokenList := types.TokenListJSON{
		Tokens: []types.TokenJSON{
			token1, token2,
		},
	}

	bytes, err := json.Marshal(tokenList)
	if err != nil {
		panic(err)
	}

	// length
	fmt.Println("JSON length:",len(bytes))

	// write file
	if err = ioutil.WriteFile("encoded_json.txt", bytes, 0644); err != nil {
		panic(err)
	}
}
