package main

import (
	"encoding/json"
	"testing"

	"github.com/tendermint/go-amino"
	"google.golang.org/protobuf/proto"
)

func BenchmarkMarshalJSON(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(tokenListGo)
	}
}

func BenchmarkMarshalAmino(b *testing.B) {
	aminoCdc := amino.NewCodec()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = aminoCdc.MarshalBinaryLengthPrefixed(tokenListGo)
	}
}

func BenchmarkMarshalGoprotobuf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = proto.Marshal(&tokenListProto)
	}
}

func BenchmarkMarshalGogoprotobuf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = tokenListGogoProto.Marshal()
	}
}
