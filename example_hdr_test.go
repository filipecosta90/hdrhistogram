package hdrhistogram_test

import (
	"fmt"
	"github.com/filipecosta90/hdrhistogram"
)

// This latency Histogram could be used to track and analyze the counts of
// observed integer values between 0 us and 30000000 us ( 30 secs )
// while maintaining a value precision of 4 significant digits across that range,
// translating to a value resolution of :
//   - 1 microsecond up to 10 milliseconds,
//   - 100 microsecond (or better) from 10 milliseconds up to 10 seconds,
//   - 300 microsecond (or better) from 10 seconds up to 30 seconds,
func ExampleNew() {
	lH := hdrhistogram.New(1, 30000000, 4)
	input := []int64{
		459876, 669187, 711612, 816326, 931423, 1033197, 1131895, 2477317,
		3964974, 12718782,
	}

	for _, sample := range input {
		lH.RecordValue(sample)
	}

	fmt.Printf("Percentile 50: %d\n", lH.ValueAtQuantile(50.0))

	// Output:
	// Percentile 50: 931423
}
