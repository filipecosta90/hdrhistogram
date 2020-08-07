package bench_deferrable_ops

import (
	"github.com/filipecosta90/hdrhistogram"
	"math/rand"
	"testing"
)

var prepoluated_nvalues = 1000000

func getValues(seed int64, nElements int, min int64, max int64) (res []int64) {
	rand.Seed(seed)
	randN := max - min
	s := make([]int64, nElements, nElements)
	for i := 0; i < nElements; i++ {
		s[i] = min + rand.Int63n(randN)
	}
	return s
}

func makeHDR(min int64, max int64, sigfigs int) (res *hdrhistogram.Histogram) {
	res = hdrhistogram.New(min, max, sigfigs)
	return
}

func benchmarkHDR_RecordValue(res []int64, hist *hdrhistogram.Histogram, b *testing.B) {
	for n := 0; n < b.N; n++ {
		if err := hist.RecordValue(res[n]); err != nil {
			b.Fatal(err)
		}
	}
}

func prepopulatedHistogram(prepoluated_nvalues int, hist *hdrhistogram.Histogram, max int64, b *testing.B) (h *hdrhistogram.Histogram) {
	for i := 0; i < prepoluated_nvalues; i++ {
		if err := hist.RecordValue(int64(i) % max); err != nil {
			b.Fatal(err)
		}
	}
	return hist
}
