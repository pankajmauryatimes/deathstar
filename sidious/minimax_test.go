package main

import (
	"github.com/danmane/deathstar/implgame"
	"math"
	"testing"
)

var s implgame.State = implgame.Standard

func Benchmark_Minimax_1(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		Minimax(&s, 1, true, defaultWeights)
	}
}

func Benchmark_Minimax_2(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		Minimax(&s, 2, true, defaultWeights)
	}
}

func Benchmark_AlphaBeta_1(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		AlphaBeta(&s, 1, math.MinInt64, math.MaxInt64, true, defaultWeights)
	}
}

func Benchmark_AlphaBeta_2(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		AlphaBeta(&s, 2, math.MinInt64, math.MaxInt64, true, defaultWeights)
	}
}

func Benchmark_AlphaBeta_3(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		AlphaBeta(&s, 3, math.MinInt64, math.MaxInt64, true, defaultWeights)
	}
}

func Benchmark_TimedAlphaBeta_3(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		TimedAlphaBeta(&s, 3, true, 3*1000*1000*1000, defaultWeights)
	}
}

// func Benchmark_bustaMove_3(b *testing.B) {
// 	for n := 0; n <= b.N; n++ {
// 		Minimax(&s, 3)
// 	}
// }
