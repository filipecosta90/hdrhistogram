package hdrhistogram

// A WindowedHistogram combines histograms to provide windowed statistics.
type WindowedHistogram struct {
	idx int
	h   []Histogram
	m   *Histogram

	Current *Histogram
}

// NewWindowed creates a new WindowedHistogram with N underlying histograms with
// the given parameters.
func NewWindowed(n int, minValue, maxValue int64, sigfigs int) *WindowedHistogram {
	w := WindowedHistogram{
		idx: -1,
		h:   make([]Histogram, n),
		m:   New(minValue, maxValue, sigfigs),
	}

	for i := range w.h {
		w.h[i] = *New(minValue, maxValue, sigfigs)
	}
	w.Rotate()

	return &w
}

// Merge returns a bench_histogram which includes the recorded values from all the
// sections of the window.
func (w *WindowedHistogram) Merge() *Histogram {
	w.m.Reset()
	for _, h := range w.h {
		w.m.Merge(&h)
	}
	return w.m
}

// Rotate resets the oldest bench_histogram and rotates it to be used as the current
// bench_histogram.
func (w *WindowedHistogram) Rotate() {
	w.idx++
	w.Current = &w.h[w.idx%len(w.h)]
	w.Current.Reset()
}

// ByteSize returns an estimate of the amount of memory allocated from
// all the sections of the window.
//
// N.B.: This does not take into account the overhead for slices, which are
// small, constant, and specific to the compiler version.
func (w *WindowedHistogram) ByteSize() int {
	value := 0
	for _, h := range w.h {
		value += h.ByteSize()
	}
	return value
}
