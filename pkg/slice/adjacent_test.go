package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestElimnateAdjDuplicate(t *testing.T) {
	s := []string{"The", "fast", "fast", "and", " ", " ", "the", "furious", "furious"}
	ss := elimnateAdjDuplicate(s)
	assert.Equal(t, []string{"The", "fast", "and", " ", "the", "furious"}, ss)
}

func TestElimnateAdjDuplicate1(t *testing.T) {
	s := []string{"The", "fast", "fast", "and", " ", " ", "the", "furious", "furious"}
	ss := elimnateAdjDuplicate1(s)
	assert.Equal(t, []string{"The", "fast", "and", " ", "the", "furious"}, ss)
}

func bench(b *testing.B, f func([]string) []string) {
	for i := 0; i < b.N; i++ {
		s := []string{"The", "fast", "fast", "and", " ", " ", "the", "furious", "furious"}
		f(s)
	}
}

func BenchmarkElimnateAdjDuplicate(b *testing.B) {
	bench(b, elimnateAdjDuplicate)
}

func BenchmarkElimnateAdjDuplicate1(b *testing.B) {
	bench(b, elimnateAdjDuplicate1)
}
