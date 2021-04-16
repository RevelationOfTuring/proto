package main

import (
	"fmt"
	"github.com/tendermint/go-amino"
)

func main() {
	// amino cdc
	cdc := amino.NewCodec()
	bytes, err := cdc.MarshalBinaryLengthPrefixed(11)
	if err != nil {
		panic(err)
	}
	fmt.Println(bytes)
	// JSON

	// amino

	// goprotobuf

	// gogoprotobuf
}
