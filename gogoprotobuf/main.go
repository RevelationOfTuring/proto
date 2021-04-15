package main

import (
	"fmt"

	"proto/gogoprotobuf/types/gogoprotobuf"
	"proto/gogoprotobuf/types/protobuf"
)

func main() {
	// protobuf text format
	var testProto protobuf.Test4
	testProto.Name = "michael.w"
	fmt.Println(testProto.String())

	// go text format
	var testGogoProto gogoprotobuf.Test5
	testGogoProto.Name = "michael.w"
	fmt.Printf("%#v\n", testGogoProto)
}
