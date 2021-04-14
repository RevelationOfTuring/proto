package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"proto/protobuf/types"
)

func main() {
	// read file
	bytes, err := ioutil.ReadFile("encoded_json.txt")
	if err != nil {
		panic(err)
	}

	// TODO:
	// the pointer used for unmarshalling must be initialized with memory space
	// It will fail to unmarshal with JSON like:
	//			var tokenList *types.TokenListJSON
	tokenList := &types.TokenListJSON{}
	if err := json.Unmarshal(bytes, tokenList); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", *tokenList)
}
