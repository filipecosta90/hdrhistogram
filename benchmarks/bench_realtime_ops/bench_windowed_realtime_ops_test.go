package bench_realtime_ops

import (
	"code.cloudfoundry.org/bytefmt"
	"github.com/filipecosta90/hdrhistogram"

	"testing"
)

// record 100K records per bench_histogram and rotate it
func Benchmark_WindowedHistogram_ms_60x_1min_precision_3_RecordAndRotate(b *testing.B) {
	var min int64 = 1
	var max int64 = 60000
	var precision = 3
	w := hdrhistogram.NewWindowed(60, min, max, precision)
	b.Logf("WindowedHistogram size: %s", bytefmt.ByteSize(uint64(w.ByteSize())))
	res := getValues(12345, b.N, min, max)
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := w.Current.RecordValue(res[i]); err != nil {
			b.Fatal(err)
		}

		if i%100000 == 1 {
			w.Rotate()
		}
	}
}
