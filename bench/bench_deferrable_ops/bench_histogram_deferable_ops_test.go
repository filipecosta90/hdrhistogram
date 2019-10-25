package bench_deferrable_ops

import (
	"github.com/filipecosta90/hdrhistogram"
	"testing"
)

func init() {
	/* load test data */

}

func Benchmark_Histogram_New_us_60sec_precision_4(b *testing.B) {
	var min int64 = 1
	var max int64 = 60000000
	var precision = 4
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		hdrhistogram.New(min, max, precision) // this could track 1ms-1min
	}
}

func Benchmark_Histogram_New_us_60sec_precision_3(b *testing.B) {
	var min int64 = 1
	var max int64 = 60000000
	var precision = 3
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		hdrhistogram.New(min, max, precision) // this could track 1ms-1min
	}
}

func Benchmark_Histogram_New_ms_60sec_precision_4(b *testing.B) {
	var min int64 = 1
	var max int64 = 60000
	var precision = 4
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		hdrhistogram.New(min, max, precision) // this could track 1ms-1min
	}
}

func Benchmark_Histogram_New_ms_60sec_precision_3(b *testing.B) {
	var min int64 = 1
	var max int64 = 60000
	var precision = 3
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		hdrhistogram.New(min, max, precision) // this could track 1ms-1min
	}
}
