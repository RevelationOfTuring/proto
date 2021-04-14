package main

import (
	"fmt"
	"io/ioutil"

	"proto/protobuf/types"

	"google.golang.org/protobuf/proto"
)

func main() {
	// read file
	bytes, err := ioutil.ReadFile("encoded_proto3.txt")
	if err != nil {
		panic(err)
	}

	// TODO:
	// the pointer used for unmarshalling must be initialized with memory space
	// It will fail to unmarshal with proto3 like:
	//			var tokenList *types.TokenList
	tokenList := &types.TokenList{}
	if err := proto.Unmarshal(bytes, tokenList); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", tokenList)
}
