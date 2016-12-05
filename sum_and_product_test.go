package bigo

import "testing"

func benchmarkSumAndProduct(length int, b *testing.B) {
	array := newIntArray(length)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		SumAndProduct(array)
	}
}

func BenchmarkSumAndProduct10(b *testing.B)   { benchmarkSumAndProduct(10, b) }
func BenchmarkSumAndProduct20(b *testing.B)   { benchmarkSumAndProduct(20, b) }
func BenchmarkSumAndProduct30(b *testing.B)   { benchmarkSumAndProduct(30, b) }
func BenchmarkSumAndProduct40(b *testing.B)   { benchmarkSumAndProduct(40, b) }
func BenchmarkSumAndProduct50(b *testing.B)   { benchmarkSumAndProduct(50, b) }
func BenchmarkSumAndProduct60(b *testing.B)   { benchmarkSumAndProduct(60, b) }
func BenchmarkSumAndProduct70(b *testing.B)   { benchmarkSumAndProduct(70, b) }
func BenchmarkSumAndProduct80(b *testing.B)   { benchmarkSumAndProduct(80, b) }
func BenchmarkSumAndProduct90(b *testing.B)   { benchmarkSumAndProduct(90, b) }
func BenchmarkSumAndProduct100(b *testing.B)  { benchmarkSumAndProduct(100, b) }
func BenchmarkSumAndProduct200(b *testing.B)  { benchmarkSumAndProduct(200, b) }
func BenchmarkSumAndProduct300(b *testing.B)  { benchmarkSumAndProduct(300, b) }
func BenchmarkSumAndProduct400(b *testing.B)  { benchmarkSumAndProduct(400, b) }
func BenchmarkSumAndProduct500(b *testing.B)  { benchmarkSumAndProduct(500, b) }
func BenchmarkSumAndProduct600(b *testing.B)  { benchmarkSumAndProduct(600, b) }
func BenchmarkSumAndProduct700(b *testing.B)  { benchmarkSumAndProduct(700, b) }
func BenchmarkSumAndProduct800(b *testing.B)  { benchmarkSumAndProduct(800, b) }
func BenchmarkSumAndProduct900(b *testing.B)  { benchmarkSumAndProduct(900, b) }
func BenchmarkSumAndProduct1000(b *testing.B) { benchmarkSumAndProduct(1000, b) }
func BenchmarkSumAndProduct2000(b *testing.B) { benchmarkSumAndProduct(2000, b) }
func BenchmarkSumAndProduct3000(b *testing.B) { benchmarkSumAndProduct(3000, b) }
func BenchmarkSumAndProduct4000(b *testing.B) { benchmarkSumAndProduct(4000, b) }
func BenchmarkSumAndProduct5000(b *testing.B) { benchmarkSumAndProduct(5000, b) }
