package c_elo_test

import (
	"testing"

	cielo "github.com/luciodesimone/elo-ds/pkg/elo/pkg"
)

var eloPairs []int = []int{2628, 2186, 1944, 2962, 1474, 1329, 347, 1916, 1801, 2313, 1936, 1456, 1242, 2215, 860, 1411, 570, 605, 2441, 2573, 919, 2835, 2460, 2757, 948, 2641, 134, 204, 2129, 2253, 1926, 786, 2950, 1674, 555, 2124, 101, 371, 2748, 1929, 982, 2130, 1313, 1777, 447, 1536, 2397, 2240, 1564, 1277, 1703, 872, 1085, 2478, 1015, 2731, 467, 1872}

func BenchmarkWinProbability(b *testing.B) {
	for i := 0; i < len(eloPairs); i += 2 {
		for n := 0; n < b.N; n++ {
			cielo.WinProbability(eloPairs[i], eloPairs[i+1])
		}
	}
}

func BenchmarkCWinProbability(b *testing.B) {
	for i := 0; i < len(eloPairs); i += 2 {
		for n := 0; n < b.N; n++ {
			cielo.CWinProbability(eloPairs[i], eloPairs[i+1])
		}
	}
}

func BenchmarkRatingDifferential(b *testing.B) {
	for i := 0; i < len(eloPairs); i += 2 {
		for n := 0; n < b.N; n++ {
			cielo.RatingDifferential(eloPairs[i], eloPairs[i+1], 0.5)
		}
	}
}

func BenchmarkCRatingDifferential(b *testing.B) {
	for i := 0; i < len(eloPairs); i += 2 {
		for n := 0; n < b.N; n++ {
			cielo.CRatingDifferential(eloPairs[i], eloPairs[i+1], 0.5)
		}
	}
}
