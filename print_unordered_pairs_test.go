package bigo

import "testing"

func benchmarkPrintUnorderedPairs(length int, b *testing.B) {
	array := newIntArray(length)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		PrintUnorderedPairs(array)
	}
}

func BenchmarkPrintUnorderedPairs10(b *testing.B)   { benchmarkPrintUnorderedPairs(10, b) }
func BenchmarkPrintUnorderedPairs20(b *testing.B)   { benchmarkPrintUnorderedPairs(20, b) }
func BenchmarkPrintUnorderedPairs30(b *testing.B)   { benchmarkPrintUnorderedPairs(30, b) }
func BenchmarkPrintUnorderedPairs40(b *testing.B)   { benchmarkPrintUnorderedPairs(40, b) }
func BenchmarkPrintUnorderedPairs50(b *testing.B)   { benchmarkPrintUnorderedPairs(50, b) }
func BenchmarkPrintUnorderedPairs60(b *testing.B)   { benchmarkPrintUnorderedPairs(60, b) }
func BenchmarkPrintUnorderedPairs70(b *testing.B)   { benchmarkPrintUnorderedPairs(70, b) }
func BenchmarkPrintUnorderedPairs80(b *testing.B)   { benchmarkPrintUnorderedPairs(80, b) }
func BenchmarkPrintUnorderedPairs90(b *testing.B)   { benchmarkPrintUnorderedPairs(90, b) }
func BenchmarkPrintUnorderedPairs100(b *testing.B)  { benchmarkPrintUnorderedPairs(100, b) }
func BenchmarkPrintUnorderedPairs200(b *testing.B)  { benchmarkPrintUnorderedPairs(200, b) }
func BenchmarkPrintUnorderedPairs300(b *testing.B)  { benchmarkPrintUnorderedPairs(300, b) }
func BenchmarkPrintUnorderedPairs400(b *testing.B)  { benchmarkPrintUnorderedPairs(400, b) }
func BenchmarkPrintUnorderedPairs500(b *testing.B)  { benchmarkPrintUnorderedPairs(500, b) }
func BenchmarkPrintUnorderedPairs600(b *testing.B)  { benchmarkPrintUnorderedPairs(600, b) }
func BenchmarkPrintUnorderedPairs700(b *testing.B)  { benchmarkPrintUnorderedPairs(700, b) }
func BenchmarkPrintUnorderedPairs800(b *testing.B)  { benchmarkPrintUnorderedPairs(800, b) }
func BenchmarkPrintUnorderedPairs900(b *testing.B)  { benchmarkPrintUnorderedPairs(900, b) }
func BenchmarkPrintUnorderedPairs1000(b *testing.B) { benchmarkPrintUnorderedPairs(1000, b) }
func BenchmarkPrintUnorderedPairs2000(b *testing.B) { benchmarkPrintUnorderedPairs(2000, b) }
func BenchmarkPrintUnorderedPairs3000(b *testing.B) { benchmarkPrintUnorderedPairs(3000, b) }
func BenchmarkPrintUnorderedPairs4000(b *testing.B) { benchmarkPrintUnorderedPairs(4000, b) }
func BenchmarkPrintUnorderedPairs5000(b *testing.B) { benchmarkPrintUnorderedPairs(5000, b) }
