## Running the benchmarks

HdrHistogram is designed for recoding histograms of value measurements in latency and performance sensitive applications. 
With that in mind we've splited the benchmarks into two major groups:
- real-time ops : the ops with real-time requirements like recording a value, rotating widowed histograms, etc...
- deferrable ops : the ones without real-time requirements or events that are non-deterministic regarding it's requirements like histogram creation, or windowed histogram merging that occurs for example once an hour

Measurements show value recording times as low as 11-15 nanoseconds even on local laptop machines. That is, 1,000,000,000 (1 billion) recordings can be made at a total cost of around 15 seconds on modern hardware. A Histogram's memory footprint is constant, with no allocation operations involved in recording data values or in iterating through them. The memory footprint is fixed regardless of the number of data value samples recorded, and depends solely on the dynamic range and precision chosen. 

### Common use-cases expected memory footprint
To illustrate the expected memory footprint of common setups, we've logged on each benchmark the 
estimate of the amount of memory allocated to the histogram ( This does not take into account the overhead for slices, which are small, constant, and specific to the compiler version. )

##### Microsecond precision
- **152.1 KB** to track and analyze the counts of observed integer values up to 180 seconds ( 3min ) with **microsecond precision**:
while maintaining a value precision of 3 significant digits across that range, meaning that value quantization within the range will thus be no larger than 1/1,000th (or 0.1%) of any value

- **1.9 MB** to track and analyze the counts of observed integer values up to 180 seconds ( 3min ) with **microsecond precision**:
while maintaining a value precision of 4 significant digits across that range, meaning that value quantization within the range will thus be no larger than 1/10,000th (or 0.01%) of any value
    
- **2.4 MB** to track and analyze the counts of observed integer values up to 3600 seconds ( 60min ) with **microsecond precision**:
while maintaining a value precision of 4 significant digits across that range, meaning that value quantization within the range will thus be no larger than 1/10,000th (or 0.01%) of any value
     
        
##### Millisecond precision
- **72.1 KB** to track and analyze the counts of observed integer values up to 180 seconds ( 3min ) with **millisecond precision**:
while maintaining a value precision of 3 significant digits across that range, meaning that value quantization within the range will thus be no larger than 1/1,000th (or 0.1%) of any value


- **640.1 KB** to track and analyze the counts of observed integer values up to 180 seconds ( 3min ) with **millisecond precision**:
while maintaining a value precision of 4 significant digits across that range, meaning that value quantization within the range will thus be no larger than 1/10,000th (or 0.01%) of any value
    
- **1.1 MB** to track and analyze the counts of observed integer values up to 3600 seconds ( 60min ) with **millisecond precision**:
while maintaining a value precision of 4 significant digits across that range, meaning that value quantization within the range will thus be no larger than 1/10,000th (or 0.01%) of any value
     
          
### Constant access time          
The amount of work involved in recording a sample is constant, and directly computes storage index locations such that no iteration or searching is ever involved in recording data values. From the real-time results bellow you can see that there is no difference in storing data in a histogram prepopulated with 1 million records vs storing data in a fresh new histogram. 

## Sample benchmark results 

### Real-time operations
100,000,000 (100 million) random deterministic insertions 
```
# make sure you're at the correct folder 
$ cd $GOPATH/src/github.com/filipecosta90/hdrhistogram/bench/bench_realtime_ops
$  go test  -run=XXX -bench=.  -benchtime=100000000x
  goos: darwin
  goarch: amd64
  pkg: github.com/filipecosta90/hdrhistogram/bench/bench_realtime_ops
  Benchmark_Histogram_RecordValue_us_60min_precision_4_emptystart-12              100000000               16.4 ns/op             0 B/op          0 allocs/op
  --- BENCH: Benchmark_Histogram_RecordValue_us_60min_precision_4_emptystart-12
      bench_histogram_realtime_ops_test.go:25: hist size: 2.4M
      bench_histogram_realtime_ops_test.go:25: hist size: 2.4M
  Benchmark_Histogram_RecordValue_us_60min_precision_4_prepopulated1M-12          100000000               16.3 ns/op             0 B/op          0 allocs/op
  --- BENCH: Benchmark_Histogram_RecordValue_us_60min_precision_4_prepopulated1M-12
      bench_histogram_realtime_ops_test.go:37: hist size: 2.4M
      bench_histogram_realtime_ops_test.go:37: hist size: 2.4M
  Benchmark_Histogram_RecordValue_us_60min_precision_3_emptystart-12              100000000               14.8 ns/op             0 B/op          0 allocs/op
  --- BENCH: Benchmark_Histogram_RecordValue_us_60min_precision_3_emptystart-12
      bench_histogram_realtime_ops_test.go:50: hist size: 184.1K
      bench_histogram_realtime_ops_test.go:50: hist size: 184.1K
  Benchmark_Histogram_RecordValue_us_60min_precision_3_prepopulated1M-12          100000000               14.7 ns/op             0 B/op          0 allocs/op
  --- BENCH: Benchmark_Histogram_RecordValue_us_60min_precision_3_prepopulated1M-12
      bench_histogram_realtime_ops_test.go:62: hist size: 184.1K
      bench_histogram_realtime_ops_test.go:62: hist size: 184.1K
  Benchmark_Histogram_RecordValue_us_3min_precision_4_emptystart-12               100000000               12.9 ns/op             0 B/op          0 allocs/op
  --- BENCH: Benchmark_Histogram_RecordValue_us_3min_precision_4_emptystart-12
      bench_histogram_realtime_ops_test.go:75: hist size: 1.9M
      bench_histogram_realtime_ops_test.go:75: hist size: 1.9M
  Benchmark_Histogram_RecordValue_us_3min_precision_4_prepopulated1M-12           100000000               12.9 ns/op             0 B/op          0 allocs/op
  --- BENCH: Benchmark_Histogram_RecordValue_us_3min_precision_4_prepopulated1M-12
      bench_histogram_realtime_ops_test.go:87: hist size: 1.9M
      bench_histogram_realtime_ops_test.go:87: hist size: 1.9M
  Benchmark_Histogram_RecordValue_us_3min_precision_3_emptystart-12               100000000               11.1 ns/op             0 B/op          0 allocs/op
  --- BENCH: Benchmark_Histogram_RecordValue_us_3min_precision_3_emptystart-12
      bench_histogram_realtime_ops_test.go:100: hist size: 152.1K
      bench_histogram_realtime_ops_test.go:100: hist size: 152.1K
  Benchmark_Histogram_RecordValue_us_3min_precision_3_prepopulated1M-12           100000000               11.2 ns/op             0 B/op          0 allocs/op
  --- BENCH: Benchmark_Histogram_RecordValue_us_3min_precision_3_prepopulated1M-12
      bench_histogram_realtime_ops_test.go:112: hist size: 152.1K
      bench_histogram_realtime_ops_test.go:112: hist size: 152.1K
  Benchmark_Histogram_RecordValue_us_1min_precision_4_emptystart-12               100000000               12.9 ns/op             0 B/op          0 allocs/op
  --- BENCH: Benchmark_Histogram_RecordValue_us_1min_precision_4_emptystart-12
      bench_histogram_realtime_ops_test.go:125: hist size: 1.6M
      bench_histogram_realtime_ops_test.go:125: hist size: 1.6M
  Benchmark_Histogram_RecordValue_us_1min_precision_4_prepopulated1M-12           100000000               12.9 ns/op             0 B/op          0 allocs/op
  --- BENCH: Benchmark_Histogram_RecordValue_us_1min_precision_4_prepopulated1M-12
      bench_histogram_realtime_ops_test.go:137: hist size: 1.6M
      bench_histogram_realtime_ops_test.go:137: hist size: 1.6M
  Benchmark_Histogram_RecordValue_us_1min_precision_3_emptystart-12               100000000               11.2 ns/op             0 B/op          0 allocs/op
  --- BENCH: Benchmark_Histogram_RecordValue_us_1min_precision_3_emptystart-12
      bench_histogram_realtime_ops_test.go:150: hist size: 136.1K
      bench_histogram_realtime_ops_test.go:150: hist size: 136.1K
  Benchmark_Histogram_RecordValue_us_1min_precision_3_prepopulated1M-12           100000000               11.2 ns/op             0 B/op          0 allocs/op
  --- BENCH: Benchmark_Histogram_RecordValue_us_1min_precision_3_prepopulated1M-12
      bench_histogram_realtime_ops_test.go:162: hist size: 136.1K
      bench_histogram_realtime_ops_test.go:162: hist size: 136.1K
  Benchmark_Histogram_RecordValue_us_1sec_precision_4_emptystart-12               100000000               13.0 ns/op             0 B/op          0 allocs/op
  --- BENCH: Benchmark_Histogram_RecordValue_us_1sec_precision_4_emptystart-12
      bench_histogram_realtime_ops_test.go:175: hist size: 896.1K
      bench_histogram_realtime_ops_test.go:175: hist size: 896.1K
  Benchmark_Histogram_RecordValue_us_1sec_precision_4_prepopulated1M-12           100000000               13.0 ns/op             0 B/op          0 allocs/op
  --- BENCH: Benchmark_Histogram_RecordValue_us_1sec_precision_4_prepopulated1M-12
      bench_histogram_realtime_ops_test.go:187: hist size: 896.1K
      bench_histogram_realtime_ops_test.go:187: hist size: 896.1K
  Benchmark_Histogram_RecordValue_us_1sec_precision_3_emptystart-12               100000000               11.4 ns/op             0 B/op          0 allocs/op
  --- BENCH: Benchmark_Histogram_RecordValue_us_1sec_precision_3_emptystart-12
      bench_histogram_realtime_ops_test.go:200: hist size: 88.1K
      bench_histogram_realtime_ops_test.go:200: hist size: 88.1K
  Benchmark_Histogram_RecordValue_us_1sec_precision_3_prepopulated1M-12           100000000               11.3 ns/op             0 B/op          0 allocs/op
  --- BENCH: Benchmark_Histogram_RecordValue_us_1sec_precision_3_prepopulated1M-12
      bench_histogram_realtime_ops_test.go:212: hist size: 88.1K
      bench_histogram_realtime_ops_test.go:212: hist size: 88.1K
  Benchmark_Histogram_RecordValue_ms_60min_precision_4_emptystart-12              100000000               13.0 ns/op             0 B/op          0 allocs/op
  --- BENCH: Benchmark_Histogram_RecordValue_ms_60min_precision_4_emptystart-12
      bench_histogram_realtime_ops_test.go:226: hist size: 1.1M
      bench_histogram_realtime_ops_test.go:226: hist size: 1.1M
  Benchmark_Histogram_RecordValue_ms_60min_precision_4_prepopulated1M-12          100000000               12.8 ns/op             0 B/op          0 allocs/op
  --- BENCH: Benchmark_Histogram_RecordValue_ms_60min_precision_4_prepopulated1M-12
      bench_histogram_realtime_ops_test.go:238: hist size: 1.1M
      bench_histogram_realtime_ops_test.go:238: hist size: 1.1M
  Benchmark_Histogram_RecordValue_ms_60min_precision_3_emptystart-12              100000000               11.1 ns/op             0 B/op          0 allocs/op
  --- BENCH: Benchmark_Histogram_RecordValue_ms_60min_precision_3_emptystart-12
      bench_histogram_realtime_ops_test.go:251: hist size: 104.1K
      bench_histogram_realtime_ops_test.go:251: hist size: 104.1K
  Benchmark_Histogram_RecordValue_ms_60min_precision_3_prepopulated1M-12          100000000               11.1 ns/op             0 B/op          0 allocs/op
  --- BENCH: Benchmark_Histogram_RecordValue_ms_60min_precision_3_prepopulated1M-12
      bench_histogram_realtime_ops_test.go:263: hist size: 104.1K
      bench_histogram_realtime_ops_test.go:263: hist size: 104.1K
  Benchmark_Histogram_RecordValue_ms_3min_precision_4_emptystart-12               100000000               13.6 ns/op             0 B/op          0 allocs/op
  --- BENCH: Benchmark_Histogram_RecordValue_ms_3min_precision_4_emptystart-12
      bench_histogram_realtime_ops_test.go:276: hist size: 640.1K
      bench_histogram_realtime_ops_test.go:276: hist size: 640.1K
  Benchmark_Histogram_RecordValue_ms_3min_precision_3_emptystart-12               100000000               12.5 ns/op             0 B/op          0 allocs/op
  --- BENCH: Benchmark_Histogram_RecordValue_ms_3min_precision_3_emptystart-12
      bench_histogram_realtime_ops_test.go:288: hist size: 72.1K
      bench_histogram_realtime_ops_test.go:288: hist size: 72.1K
  Benchmark_Histogram_RecordValue_ms_1min_precision_4_emptystart-12               100000000               15.1 ns/op             0 B/op          0 allocs/op
  --- BENCH: Benchmark_Histogram_RecordValue_ms_1min_precision_4_emptystart-12
      bench_histogram_realtime_ops_test.go:300: hist size: 384.1K
      bench_histogram_realtime_ops_test.go:300: hist size: 384.1K
  Benchmark_Histogram_RecordValue_ms_1min_precision_4_prepopulated1M-12           100000000               15.1 ns/op             0 B/op          0 allocs/op
  --- BENCH: Benchmark_Histogram_RecordValue_ms_1min_precision_4_prepopulated1M-12
      bench_histogram_realtime_ops_test.go:312: hist size: 384.1K
      bench_histogram_realtime_ops_test.go:312: hist size: 384.1K
  Benchmark_Histogram_RecordValue_ms_1min_precision_3_emptystart-12               100000000               14.6 ns/op             0 B/op          0 allocs/op
  --- BENCH: Benchmark_Histogram_RecordValue_ms_1min_precision_3_emptystart-12
      bench_histogram_realtime_ops_test.go:325: hist size: 56.1K
      bench_histogram_realtime_ops_test.go:325: hist size: 56.1K
  Benchmark_Histogram_RecordValue_ms_1min_precision_3_prepopulated1M-12           100000000               14.6 ns/op             0 B/op          0 allocs/op
  --- BENCH: Benchmark_Histogram_RecordValue_ms_1min_precision_3_prepopulated1M-12
      bench_histogram_realtime_ops_test.go:337: hist size: 56.1K
      bench_histogram_realtime_ops_test.go:337: hist size: 56.1K
  Benchmark_Histogram_RecordValue_ms_1sec_precision_4_emptystart-12               100000000               10.9 ns/op             0 B/op          0 allocs/op
  --- BENCH: Benchmark_Histogram_RecordValue_ms_1sec_precision_4_emptystart-12
      bench_histogram_realtime_ops_test.go:350: hist size: 256.1K
      bench_histogram_realtime_ops_test.go:350: hist size: 256.1K
  Benchmark_Histogram_RecordValue_ms_1sec_precision_4_prepopulated1M-12           100000000               10.9 ns/op             0 B/op          0 allocs/op
  --- BENCH: Benchmark_Histogram_RecordValue_ms_1sec_precision_4_prepopulated1M-12
      bench_histogram_realtime_ops_test.go:362: hist size: 256.1K
      bench_histogram_realtime_ops_test.go:362: hist size: 256.1K
  Benchmark_Histogram_RecordValue_ms_1sec_precision_3_emptystart-12               100000000               11.0 ns/op             0 B/op          0 allocs/op
  --- BENCH: Benchmark_Histogram_RecordValue_ms_1sec_precision_3_emptystart-12
      bench_histogram_realtime_ops_test.go:375: hist size: 16.1K
      bench_histogram_realtime_ops_test.go:375: hist size: 16.1K
  Benchmark_Histogram_RecordValue_ms_1sec_precision_3_prepopulated1M-12           100000000               11.0 ns/op             0 B/op          0 allocs/op
  --- BENCH: Benchmark_Histogram_RecordValue_ms_1sec_precision_3_prepopulated1M-12
      bench_histogram_realtime_ops_test.go:387: hist size: 16.1K, 
      bench_histogram_realtime_ops_test.go:387: hist size: 16.1K, 
  Benchmark_WindowedHistogram_ms_60x_1min_precision_3_RecordAndRotate-12          100000000               15.4 ns/op             0 B/op          0 allocs/op
  --- BENCH: Benchmark_WindowedHistogram_ms_60x_1min_precision_3_RecordAndRotate-12
      bench_windowed_realtime_ops_test.go:16: WindowedHistogram size: 3.3M
      bench_windowed_realtime_ops_test.go:16: WindowedHistogram size: 3.3M
  PASS
  ok      github.com/filipecosta90/hdrhistogram/bench/bench_realtime_ops  123.438s

```


### Deferrable/non-real-time operations 
Each of the Histograms with a WindowedHistogram is prepopulated with 1,000,000 (1 million) values 
```
# make sure you're at the correct folder 
$ cd $GOPATH/src/github.com/filipecosta90/hdrhistogram/bench/bench_deferable_ops
$ go test  -run=XXX -bench=.  -benchtime=10000x  
  goos: darwin
  goarch: amd64
  pkg: github.com/filipecosta90/hdrhistogram/bench/bench_deferable_ops
  Benchmark_Histogram_New_ns_60sec_precision_4-12                            10000            116638 ns/op         1704041 B/op          2 allocs/op
  Benchmark_Histogram_New_ns_60sec_precision_3-12                            10000              7708 ns/op          139360 B/op          2 allocs/op
  Benchmark_Histogram_New_ms_60sec_precision_4-12                            10000             20564 ns/op          393314 B/op          2 allocs/op
  Benchmark_Histogram_New_ms_60sec_precision_3-12                            10000              3735 ns/op           57440 B/op          2 allocs/op
  Benchmark_WindowedHistogram_ms_60x_1min_precision_3_Merge-12               10000          13538546 ns/op               0 B/op          0 allocs/op
  --- BENCH: Benchmark_WindowedHistogram_ms_60x_1min_precision_3_Merge-12
      bench_windowed_deferable_ops_test.go:17: WindowedHistogram size: 3.3M
      bench_windowed_deferable_ops_test.go:17: WindowedHistogram size: 3.3M
  PASS
  ok      github.com/filipecosta90/hdrhistogram/bench/bench_deferable_ops 139.117s
```