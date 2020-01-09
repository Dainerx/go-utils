package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommaRec(t *testing.T) {
	assert.Equal(t, "1,236,749", CommaRec("1236749"))
}

func TestCommaBytesBuffer(t *testing.T) {
	assert.Equal(t, "1,233", CommaBytesBuffer("1233"))
	assert.Equal(t, "12", CommaBytesBuffer("12"))
	assert.Equal(t, "1,236,749", CommaBytesBuffer("1236749"))
}

func TestCommaFloatEnhanced(t *testing.T) {
	assert.Equal(t, CommaFloatEnhancedSafe("-1235.2991237"), CommaFloatEnhanced("-1235.2991237"))
}

func bench(b *testing.B, f func(string) string) {
	for i := 0; i < b.N; i++ {
		s := "1236749"
		f(s)
	}
}
func BenchmarkCommaRec(b *testing.B) {
	bench(b, CommaRec)
}

func BenchmarkCommaBytesBuffer(b *testing.B) {
	bench(b, CommaBytesBuffer)
}

func BenchmarkCommaFloatEnhanced(b *testing.B) {
	bench(b, CommaFloatEnhanced)
}

func BenchmarkCommaFloatEnhancedsafe(b *testing.B) {
	bench(b, CommaFloatEnhancedSafe)
}
