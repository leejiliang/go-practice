package ch14

import (
	"bytes"
	"testing"
)

func concat(arr []string) string {
	var ret string
	for _, v := range arr {
		ret += v
	}
	return ret
}

func concatByBuffer(arr []string) string {
	var ret string
	var buffer bytes.Buffer
	for _, v := range arr {
		buffer.WriteString(v)
	}
	return ret
}

func BenchmarkConcat(b *testing.B) {
	strArr := [...]string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		concat(strArr[:])
	}
	b.StopTimer()
}

func BenchmarkConcatByBuffer(b *testing.B) {
	strArr := [...]string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		concatByBuffer(strArr[:])
	}
	b.StopTimer()
}
