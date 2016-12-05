package bigo

import "testing"

func benchmarkFib(v int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(v)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(1, b) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(2, b) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(3, b) }
func BenchmarkFib4(b *testing.B)  { benchmarkFib(4, b) }
func BenchmarkFib5(b *testing.B)  { benchmarkFib(5, b) }
func BenchmarkFib6(b *testing.B)  { benchmarkFib(6, b) }
func BenchmarkFib7(b *testing.B)  { benchmarkFib(7, b) }
func BenchmarkFib8(b *testing.B)  { benchmarkFib(8, b) }
func BenchmarkFib9(b *testing.B)  { benchmarkFib(9, b) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
func BenchmarkFib11(b *testing.B) { benchmarkFib(11, b) }
func BenchmarkFib12(b *testing.B) { benchmarkFib(12, b) }
func BenchmarkFib13(b *testing.B) { benchmarkFib(13, b) }
func BenchmarkFib14(b *testing.B) { benchmarkFib(14, b) }
func BenchmarkFib15(b *testing.B) { benchmarkFib(15, b) }
func BenchmarkFib16(b *testing.B) { benchmarkFib(16, b) }
func BenchmarkFib17(b *testing.B) { benchmarkFib(17, b) }
func BenchmarkFib18(b *testing.B) { benchmarkFib(18, b) }
func BenchmarkFib19(b *testing.B) { benchmarkFib(19, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
func BenchmarkFib21(b *testing.B) { benchmarkFib(21, b) }
func BenchmarkFib22(b *testing.B) { benchmarkFib(22, b) }
func BenchmarkFib23(b *testing.B) { benchmarkFib(23, b) }
