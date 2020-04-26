package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDraw(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		name string
		fun  func() string
	}{
		{"draw1", draw1},
		{"draw2", draw2},
		{"draw3", draw3},
	}

	for _, test := range tests {
		r1 := test.fun()
		r2, err2 := ioutil.ReadFile("output.txt")
		assert.Nil(err2)
		assert.True(bytes.Equal([]byte(r1), r2), test.name)
	}

}

func BenchmarkDraw(b *testing.B) {

	benchmarks := []struct {
		name string
		fun  func() string
	}{
		{"draw1", draw1},
		{"draw2", draw2},
		{"draw3", draw3},
	}
	for _, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("%s", benchmark.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				benchmark.fun()
			}
		})
	}
}
