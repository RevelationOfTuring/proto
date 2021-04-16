package main

import (
	"fmt"

	"proto/gogoprotobuf/types/gogoprotobuf"
	"proto/gogoprotobuf/types/protobuf"
)

func main() {
	// protobuf text format
	var testProto protobuf.Test5
	testProto.Name = "michael.w"
	fmt.Println(testProto.String())

	// golang text format
	var testGogoProto gogoprotobuf.Test6
	testGogoProto.Name = "michael.w"
	fmt.Printf("%#v\n", testGogoProto)
}
