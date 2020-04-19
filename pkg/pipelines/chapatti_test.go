package pipelines

import (
	"testing"
	"time"
)

var defaults = Chappati{
	Verbose:     false,
	Count:       20,
	NumBakers:   1,
	BakeTime:    10 * time.Millisecond,
	MakeTime:    10 * time.Millisecond,
	PackageTime: 10 * time.Millisecond,
}

func Benchmark(b *testing.B) {
	chapattiAriana := defaults
	chapattiAriana.Run(b.N) // ~2.25 s
}

func BenchmarkBakers(b *testing.B) {
	chapattiAriana := defaults
	chapattiAriana.NumBakers++ // add one baker to the team
	chapattiAriana.Run(b.N)    // ~2.25 s
}

// what if we add many bakers?
func BenchmarkBuffers(b *testing.B) {
	// Adding buffers has no effect.
	chapattiAriana := defaults
	chapattiAriana.NumBakers = 100 // add one baker to the team
	chapattiAriana.Run(b.N)        // this will not be faster since most of the bakers are jobless through out the business day
}
