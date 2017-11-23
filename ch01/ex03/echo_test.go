package main

import (
	"io/ioutil"
	"testing"
)

func getArgs(n int) []string {
	args := make([]string, 0, n)
	for i := 0; i < n; i++ {
		args = append(args, "hoge")
	}
	return args
}

func BenchmarkEchoA(b *testing.B) {
	args := getArgs(100)
	for i := 0; i < b.N; i++ {
		EchoA(ioutil.Discard, args)
	}
}

func BenchmarkEchoB(b *testing.B) {
	args := getArgs(100)
	for i := 0; i < b.N; i++ {
		EchoB(ioutil.Discard, args)
	}
}
