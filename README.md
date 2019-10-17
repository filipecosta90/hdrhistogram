

[![license](https://img.shields.io/github/license/filipecosta90/hdrhistogram.svg)](https://github.com/filipecosta90/hdrhistogram)
[![CircleCI](https://circleci.com/gh/filipecosta90/hdrhistogram/tree/master.svg?style=svg)](https://circleci.com/gh/filipecosta90/hdrhistogram/tree/master)
[![GitHub issues](https://img.shields.io/github/release/filipecosta90/hdrhistogram.svg)](https://github.com/filipecosta90/hdrhistogram/releases/latest)
[![Codecov](https://codecov.io/gh/filipecosta90/hdrhistogram/branch/master/graph/badge.svg)](https://codecov.io/gh/filipecosta90/hdrhistogram)
[![Go Report Card](https://goreportcard.com/badge/github.com/filipecosta90/hdrhistogram)](https://goreportcard.com/report/github.com/filipecosta90/hdrhistogram)
[![GoDoc](https://godoc.org/github.com/filipecosta90/hdrhistogram?status.svg)](https://godoc.org/github.com/filipecosta90/hdrhistogram)

hdrhistogram
============
A pure Go implementation of the [HDR Histogram](https://github.com/HdrHistogram/HdrHistogram).

> A Histogram that supports recording and analyzing sampled data value counts
> across a configurable integer value range with configurable value precision
> within the range. Value precision is expressed as the number of significant
> digits in the value recording, and provides control over value quantization
> behavior across the value range and the subsequent value resolution at any
> given level.

For documentation, check [godoc](http://godoc.org/github.com/filipecosta90/hdrhistogram).


## Installing 

```sh
go get github.com/filipecosta90/hdrhistogram
```


## Usage Example

```go
package main

import (
	"fmt"
    "github.com/filipecosta90/hdrhistogram"
)

func main() {
    // This latency Histogram could be used to track and analyze the counts of
    // observed integer values between 0 us and 30000000 us ( 30 secs )
    // while maintaining a value precision of 3 significant digits across that range,
    // translating to a value resolution of :
    //   - 1 microsecond up to 1 millisecond,
    //   - 1 millisecond (or better) up to one second,
    //   - 1 second (or better) up to it's maximum tracked value ( 30 seconds ).
    lH := hdrhistogram.New(1, 30000000, 3)
	input := []int64{
        459876, 669187, 711612, 816326, 931423, 1033197, 1131895, 2477317,
        3964974, 12718782,
    }

    for _, sample := range input {
        lH.RecordValue(sample)
    }

	fmt.Println( lH.PercentilesPrint( 1, 1.0 ) )
}
```