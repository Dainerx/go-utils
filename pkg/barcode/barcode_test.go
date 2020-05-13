package barcode

import "testing"

func BenchmarkEncodeEan(b *testing.B) {
	b.StopTimer()
	eans, _ := InitEans("ean-seed.txt")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, ean := range eans {
			EncodeEan(ean) // 600
		}
	}
}
