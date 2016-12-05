package bigo

import "testing"

func benchmarkPrintPairs(length int, b *testing.B) {
	array := newIntArray(length)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		PrintPairs(array)
	}
}

func BenchmarkPrintPairs10(b *testing.B)   { benchmarkPrintPairs(10, b) }
func BenchmarkPrintPairs20(b *testing.B)   { benchmarkPrintPairs(20, b) }
func BenchmarkPrintPairs30(b *testing.B)   { benchmarkPrintPairs(30, b) }
func BenchmarkPrintPairs40(b *testing.B)   { benchmarkPrintPairs(40, b) }
func BenchmarkPrintPairs50(b *testing.B)   { benchmarkPrintPairs(50, b) }
func BenchmarkPrintPairs60(b *testing.B)   { benchmarkPrintPairs(60, b) }
func BenchmarkPrintPairs70(b *testing.B)   { benchmarkPrintPairs(70, b) }
func BenchmarkPrintPairs80(b *testing.B)   { benchmarkPrintPairs(80, b) }
func BenchmarkPrintPairs90(b *testing.B)   { benchmarkPrintPairs(90, b) }
func BenchmarkPrintPairs100(b *testing.B)  { benchmarkPrintPairs(100, b) }
func BenchmarkPrintPairs200(b *testing.B)  { benchmarkPrintPairs(200, b) }
func BenchmarkPrintPairs300(b *testing.B)  { benchmarkPrintPairs(300, b) }
func BenchmarkPrintPairs400(b *testing.B)  { benchmarkPrintPairs(400, b) }
func BenchmarkPrintPairs500(b *testing.B)  { benchmarkPrintPairs(500, b) }
func BenchmarkPrintPairs600(b *testing.B)  { benchmarkPrintPairs(600, b) }
func BenchmarkPrintPairs700(b *testing.B)  { benchmarkPrintPairs(700, b) }
func BenchmarkPrintPairs800(b *testing.B)  { benchmarkPrintPairs(800, b) }
func BenchmarkPrintPairs900(b *testing.B)  { benchmarkPrintPairs(900, b) }
func BenchmarkPrintPairs1000(b *testing.B) { benchmarkPrintPairs(1000, b) }
func BenchmarkPrintPairs2000(b *testing.B) { benchmarkPrintPairs(2000, b) }
func BenchmarkPrintPairs3000(b *testing.B) { benchmarkPrintPairs(3000, b) }
func BenchmarkPrintPairs4000(b *testing.B) { benchmarkPrintPairs(4000, b) }
func BenchmarkPrintPairs5000(b *testing.B) { benchmarkPrintPairs(5000, b) }
