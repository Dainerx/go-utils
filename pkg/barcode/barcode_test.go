// package main

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func BenchmarkEncodeEan(b *testing.B) {
// 	b.StopTimer()
// 	eans, _ := InitEans("ean-seed.txt")
// 	b.StartTimer()
// 	for i := 0; i < b.N; i++ {
// 		for _, ean := range eans {
// 			EncodeEan(ean) // 600
// 		}
// 	}
// }

// func BenchmarkWriteImage(b *testing.B) {
// 	assert := assert.New(b)
// 	b.StopTimer()
// 	eans, _ := InitEans("seed-cleaned.txt")
// 	b.StartTimer()
// 	for i := 0; i < b.N; i++ {
// 		for _, ean := range eans {
// 			b.StopTimer()
// 			image, err := EncodeEan(ean) // 600
// 			if err != nil {
// 				continue
// 			}
// 			b.StartTimer()
// 			err = WriteImage(image, ean)
// 			assert.Nil(err, "Write image failed")
// 		}
// 	}
// }
