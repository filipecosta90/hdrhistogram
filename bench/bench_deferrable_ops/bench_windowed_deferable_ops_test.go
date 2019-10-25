package bench_deferrable_ops

import (
	"code.cloudfoundry.org/bytefmt"
	"github.com/filipecosta90/hdrhistogram"
	"testing"
)

func Benchmark_WindowedHistogram_ms_60x_1min_precision_3_Merge(b *testing.B) {
	var min int64 = 1
	var max int64 = 60000
	w := hdrhistogram.NewWindowed(60, min, max, 3)
	for i := 0; i < 60; i++ {
		prepopulatedHistogram(prepoluated_nvalues, w.Current, max, b)
		w.Rotate()
	}
	b.Logf("WindowedHistogram size: %s", bytefmt.ByteSize(uint64(w.ByteSize())))
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		w.Merge()
	}
}
