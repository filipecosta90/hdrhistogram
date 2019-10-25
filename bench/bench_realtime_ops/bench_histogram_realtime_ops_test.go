package bench_realtime_ops

import (
	"code.cloudfoundry.org/bytefmt"
	"testing"
)

func init() {
	/* load test data */

}

// HdrHistogram is designed for recording histograms of value measurements
// in latency and performance sensitive applications.
// Measurements show value recording times as low as 3-6 nanoseconds
// on modern (circa 2012) Intel CPUs.

//ns precision
func Benchmark_Histogram_RecordValue_us_3min_precision_4_emptystart(b *testing.B) {
	var min int64 = 1
	var max int64 = 180000000
	var precision = 4
	res := getValues(12345, b.N, min, max)
	hist := makeHDR(min, max, precision)
	b.Logf("hist size: %s", bytefmt.ByteSize(uint64(hist.ByteSize())))
	b.ResetTimer()
	b.ReportAllocs()
	benchmarkHDR_RecordValue(res, hist, b)
}

func Benchmark_Histogram_RecordValue_us_3min_precision_4_prepopulated1M(b *testing.B) {
	var min int64 = 1
	var max int64 = 180000000
	var precision = 4
	res := getValues(12345, b.N, min, max)
	hist := makeHDR(min, max, precision)
	b.Logf("hist size: %s", bytefmt.ByteSize(uint64(hist.ByteSize())))
	prepopulatedHistogram(prepoluated_nvalues, hist, max, b)
	b.ResetTimer()
	b.ReportAllocs()
	benchmarkHDR_RecordValue(res, hist, b)
}

func Benchmark_Histogram_RecordValue_us_3min_precision_3_emptystart(b *testing.B) {
	var min int64 = 1
	var max int64 = 180000000
	var precision = 3
	res := getValues(12345, b.N, min, max)
	hist := makeHDR(min, max, precision)
	b.Logf("hist size: %s", bytefmt.ByteSize(uint64(hist.ByteSize())))
	b.ResetTimer()
	b.ReportAllocs()
	benchmarkHDR_RecordValue(res, hist, b)
}

func Benchmark_Histogram_RecordValue_us_3min_precision_3_prepopulated1M(b *testing.B) {
	var min int64 = 1
	var max int64 = 180000000
	var precision = 3
	res := getValues(12345, b.N, min, max)
	hist := makeHDR(min, max, precision)
	b.Logf("hist size: %s", bytefmt.ByteSize(uint64(hist.ByteSize())))
	prepopulatedHistogram(prepoluated_nvalues, hist, max, b)
	b.ResetTimer()
	b.ReportAllocs()
	benchmarkHDR_RecordValue(res, hist, b)
}

func Benchmark_Histogram_RecordValue_us_1min_precision_4_emptystart(b *testing.B) {
	var min int64 = 1
	var max int64 = 60000000
	var precision = 4
	res := getValues(12345, b.N, min, max)
	hist := makeHDR(min, max, precision)
	b.Logf("hist size: %s", bytefmt.ByteSize(uint64(hist.ByteSize())))
	b.ResetTimer()
	b.ReportAllocs()
	benchmarkHDR_RecordValue(res, hist, b)
}

func Benchmark_Histogram_RecordValue_us_1min_precision_4_prepopulated1M(b *testing.B) {
	var min int64 = 1
	var max int64 = 60000000
	var precision = 4
	res := getValues(12345, b.N, min, max)
	hist := makeHDR(min, max, precision)
	b.Logf("hist size: %s", bytefmt.ByteSize(uint64(hist.ByteSize())))
	prepopulatedHistogram(prepoluated_nvalues, hist, max, b)
	b.ResetTimer()
	b.ReportAllocs()
	benchmarkHDR_RecordValue(res, hist, b)
}

func Benchmark_Histogram_RecordValue_us_1min_precision_3_emptystart(b *testing.B) {
	var min int64 = 1
	var max int64 = 60000000
	var precision = 3
	res := getValues(12345, b.N, min, max)
	hist := makeHDR(min, max, precision)
	b.Logf("hist size: %s", bytefmt.ByteSize(uint64(hist.ByteSize())))
	b.ResetTimer()
	b.ReportAllocs()
	benchmarkHDR_RecordValue(res, hist, b)
}

func Benchmark_Histogram_RecordValue_us_1min_precision_3_prepopulated1M(b *testing.B) {
	var min int64 = 1
	var max int64 = 60000000
	var precision = 3
	res := getValues(12345, b.N, min, max)
	hist := makeHDR(min, max, precision)
	b.Logf("hist size: %s", bytefmt.ByteSize(uint64(hist.ByteSize())))
	prepopulatedHistogram(prepoluated_nvalues, hist, max, b)
	b.ResetTimer()
	b.ReportAllocs()
	benchmarkHDR_RecordValue(res, hist, b)
}

func Benchmark_Histogram_RecordValue_us_1sec_precision_4_emptystart(b *testing.B) {
	var min int64 = 1
	var max int64 = 1000000
	var precision = 4
	res := getValues(12345, b.N, min, max)
	hist := makeHDR(min, max, precision)
	b.Logf("hist size: %s", bytefmt.ByteSize(uint64(hist.ByteSize())))
	b.ResetTimer()
	b.ReportAllocs()
	benchmarkHDR_RecordValue(res, hist, b)
}

func Benchmark_Histogram_RecordValue_us_1sec_precision_4_prepopulated1M(b *testing.B) {
	var min int64 = 1
	var max int64 = 1000000
	var precision = 4
	res := getValues(12345, b.N, min, max)
	hist := makeHDR(min, max, precision)
	b.Logf("hist size: %s", bytefmt.ByteSize(uint64(hist.ByteSize())))
	prepopulatedHistogram(prepoluated_nvalues, hist, max, b)
	b.ResetTimer()
	b.ReportAllocs()
	benchmarkHDR_RecordValue(res, hist, b)
}

func Benchmark_Histogram_RecordValue_us_1sec_precision_3_emptystart(b *testing.B) {
	var min int64 = 1
	var max int64 = 1000000
	var precision = 3
	res := getValues(12345, b.N, min, max)
	hist := makeHDR(min, max, precision)
	b.Logf("hist size: %s", bytefmt.ByteSize(uint64(hist.ByteSize())))
	b.ResetTimer()
	b.ReportAllocs()
	benchmarkHDR_RecordValue(res, hist, b)
}

func Benchmark_Histogram_RecordValue_us_1sec_precision_3_prepopulated1M(b *testing.B) {
	var min int64 = 1
	var max int64 = 1000000
	var precision = 3
	res := getValues(12345, b.N, min, max)
	hist := makeHDR(min, max, precision)
	b.Logf("hist size: %s", bytefmt.ByteSize(uint64(hist.ByteSize())))
	prepopulatedHistogram(prepoluated_nvalues, hist, max, b)
	b.ResetTimer()
	b.ReportAllocs()
	benchmarkHDR_RecordValue(res, hist, b)
}

//ms precision
func Benchmark_Histogram_RecordValue_ms_3min_precision_4_emptystart(b *testing.B) {
	var min int64 = 1
	var max int64 = 180000
	var precision = 4
	res := getValues(12345, b.N, min, max)
	hist := makeHDR(min, max, precision)
	b.Logf("hist size: %s", bytefmt.ByteSize(uint64(hist.ByteSize())))
	b.ResetTimer()
	b.ReportAllocs()
	benchmarkHDR_RecordValue(res, hist, b)
}

func Benchmark_Histogram_RecordValue_ms_3min_precision_3_emptystart(b *testing.B) {
	var min int64 = 1
	var max int64 = 180000
	var precision = 3
	res := getValues(12345, b.N, min, max)
	hist := makeHDR(min, max, precision)
	b.Logf("hist size: %s", bytefmt.ByteSize(uint64(hist.ByteSize())))
	b.ResetTimer()
	b.ReportAllocs()
	benchmarkHDR_RecordValue(res, hist, b)
}

func Benchmark_Histogram_RecordValue_ms_1min_precision_4_emptystart(b *testing.B) {
	var min int64 = 1
	var max int64 = 60000
	var precision = 4
	res := getValues(12345, b.N, min, max)
	hist := makeHDR(min, max, precision)
	b.Logf("hist size: %s", bytefmt.ByteSize(uint64(hist.ByteSize())))
	b.ResetTimer()
	b.ReportAllocs()
	benchmarkHDR_RecordValue(res, hist, b)
}

func Benchmark_Histogram_RecordValue_ms_1min_precision_4_prepopulated1M(b *testing.B) {
	var min int64 = 1
	var max int64 = 60000
	var precision = 4
	res := getValues(12345, b.N, min, max)
	hist := makeHDR(min, max, precision)
	b.Logf("hist size: %s", bytefmt.ByteSize(uint64(hist.ByteSize())))
	prepopulatedHistogram(prepoluated_nvalues, hist, max, b)
	b.ResetTimer()
	b.ReportAllocs()
	benchmarkHDR_RecordValue(res, hist, b)
}

func Benchmark_Histogram_RecordValue_ms_1min_precision_3_emptystart(b *testing.B) {
	var min int64 = 1
	var max int64 = 60000
	var precision = 3
	res := getValues(12345, b.N, min, max)
	hist := makeHDR(min, max, precision)
	b.Logf("hist size: %s", bytefmt.ByteSize(uint64(hist.ByteSize())))
	b.ResetTimer()
	b.ReportAllocs()
	benchmarkHDR_RecordValue(res, hist, b)
}

func Benchmark_Histogram_RecordValue_ms_1min_precision_3_prepopulated1M(b *testing.B) {
	var min int64 = 1
	var max int64 = 60000
	var precision = 3
	res := getValues(12345, b.N, min, max)
	hist := makeHDR(min, max, precision)
	b.Logf("hist size: %s", bytefmt.ByteSize(uint64(hist.ByteSize())))
	prepopulatedHistogram(prepoluated_nvalues, hist, max, b)
	b.ResetTimer()
	b.ReportAllocs()
	benchmarkHDR_RecordValue(res, hist, b)
}

func Benchmark_Histogram_RecordValue_ms_1sec_precision_4_emptystart(b *testing.B) {
	var min int64 = 1
	var max int64 = 1000
	var precision = 4
	res := getValues(12345, b.N, min, max)
	hist := makeHDR(min, max, precision)
	b.Logf("hist size: %s", bytefmt.ByteSize(uint64(hist.ByteSize())))
	b.ResetTimer()
	b.ReportAllocs()
	benchmarkHDR_RecordValue(res, hist, b)
}

func Benchmark_Histogram_RecordValue_ms_1sec_precision_4_prepopulated1M(b *testing.B) {
	var min int64 = 1
	var max int64 = 1000
	var precision = 4
	res := getValues(12345, b.N, min, max)
	hist := makeHDR(min, max, precision)
	b.Logf("hist size: %s", bytefmt.ByteSize(uint64(hist.ByteSize())))
	prepopulatedHistogram(prepoluated_nvalues, hist, max, b)
	b.ResetTimer()
	b.ReportAllocs()
	benchmarkHDR_RecordValue(res, hist, b)
}

func Benchmark_Histogram_RecordValue_ms_1sec_precision_3_emptystart(b *testing.B) {
	var min int64 = 1
	var max int64 = 1000
	var precision = 3
	res := getValues(12345, b.N, min, max)
	hist := makeHDR(min, max, precision)
	b.Logf("hist size: %s", bytefmt.ByteSize(uint64(hist.ByteSize())))
	b.ResetTimer()
	b.ReportAllocs()
	benchmarkHDR_RecordValue(res, hist, b)
}

func Benchmark_Histogram_RecordValue_ms_1sec_precision_3_prepopulated1M(b *testing.B) {
	var min int64 = 1
	var max int64 = 1000
	var precision = 3
	res := getValues(12345, b.N, min, max)
	hist := makeHDR(min, max, precision)
	b.Logf("hist size: %s, ", bytefmt.ByteSize(uint64(hist.ByteSize())))
	prepopulatedHistogram(prepoluated_nvalues, hist, max, b)
	b.ResetTimer()
	b.ReportAllocs()
	benchmarkHDR_RecordValue(res, hist, b)
}
