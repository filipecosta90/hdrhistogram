package microsecond_precision

import (
	"fmt"
	"github.com/filipecosta90/hdrhistogram"
)

func main() {
	// This latency Histogram could be used to track and analyze the counts of
	// observed integer values between 0 us and 30000000 us ( 30 secs )
	// while maintaining a value precision of 4 significant digits across that range,
	// translating to a value resolution of :
	//   - 1 microsecond up to 10 milliseconds,
	//   - 100 microsecond (or better) from 10 milliseconds up to 10 seconds,
	//   - 300 microsecond (or better) from 10 seconds up to 30 seconds,
	lH := hdrhistogram.New(1, 30000000, 4)
	input := []int64{
		459876, 669187, 711612, 816326, 931423, 1033197, 1131895, 2477317,
		3964974, 12718782,
	}

	for _, sample := range input {
		lH.RecordValue(sample)
	}

	fmt.Println(lH.PercentilesPrint(10, 1.0))
	//Value  Percentile      TotalCount      1/(1-Percentile)
	//
	//459887.000     0.000000            1         1.00
	//459887.000     0.050000            1         1.05
	//459887.000     0.100000            1         1.11
	//669215.000     0.150000            2         1.18
	//669215.000     0.200000            2         1.25
	//711615.000     0.250000            3         1.33
	//711615.000     0.300000            3         1.43
	//816351.000     0.350000            4         1.54
	//816351.000     0.400000            4         1.67
	//931423.000     0.450000            5         1.82
	//931423.000     0.500000            5         2.00
	//1033215.000     0.525000            6         2.11
	//1033215.000     0.550000            6         2.22
	//1033215.000     0.575000            6         2.35
	//1033215.000     0.600000            6         2.50
	//1131903.000     0.625000            7         2.67
	//1131903.000     0.650000            7         2.86
	//1131903.000     0.675000            7         3.08
	//1131903.000     0.700000            7         3.33
	//2477439.000     0.725000            8         3.64
	//2477439.000     0.750000            8         4.00
	//2477439.000     0.762500            8         4.21
	//2477439.000     0.775000            8         4.44
	//2477439.000     0.787500            8         4.71
	//2477439.000     0.800000            8         5.00
	//3965055.000     0.812500            9         5.33
	//3965055.000     0.825000            9         5.71
	//3965055.000     0.837500            9         6.15
	//3965055.000     0.850000            9         6.67
	//3965055.000     0.862500            9         7.27
	//3965055.000     0.875000            9         8.00
	//3965055.000     0.881250            9         8.42
	//3965055.000     0.887500            9         8.89
	//3965055.000     0.893750            9         9.41
	//3965055.000     0.900000            9        10.00
	//12719103.000     0.906250           10        10.67
	//12719103.000     1.000000           10         +Inf
	//#[Mean    =  2491471.200, StdDeviation   =  3558112.070]
	//#[Max     = 12719103.000, Total count    =           10]
	//#[Buckets =           11, SubBuckets     =        32768]
}
