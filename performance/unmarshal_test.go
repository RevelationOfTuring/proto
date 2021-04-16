package main

import (
	"encoding/json"
	"testing"

	"proto/performance/types/gogoprotobuf"
	"proto/performance/types/golang"
	"proto/performance/types/protobuf"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/go-amino"
	"google.golang.org/protobuf/proto"
)

func BenchmarkUnmarshalJSON(b *testing.B) {
	rawBytes, err := json.Marshal(tokenListGo)
	require.NoError(b, err)

	var tmp golang.TokenList
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(rawBytes, &tmp)
	}
}

func BenchmarkUnmarshalAmino(b *testing.B) {
	aminoCdc := amino.NewCodec()
	rawBytes, err := aminoCdc.MarshalBinaryLengthPrefixed(tokenListGo)
	require.NoError(b, err)

	var tmp golang.TokenList
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = amino.UnmarshalBinaryLengthPrefixed(rawBytes, &tmp)
	}
}

func BenchmarkUnmarshalGoprotobuf(b *testing.B) {
	rawBytes, err := proto.Marshal(&tokenListProto)
	require.NoError(b, err)

	tmp := &protobuf.TokenList{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = proto.Unmarshal(rawBytes, tmp)
	}
}

func BenchmarkUnmarshalGogoprotobuf(b *testing.B) {
	rawBytes, err := tokenListGogoProto.Marshal()
	require.NoError(b, err)

	tmp := &gogoprotobuf.TokenList{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = tmp.Unmarshal(rawBytes)
	}
}
