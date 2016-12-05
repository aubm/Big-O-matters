package bigo

import "testing"

func benchmarkPowersOf2(v int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		PowersOf2(v)
	}
}

func BenchmarkPowersOf2_1(b *testing.B)        { benchmarkPowersOf2(1, b) }
func BenchmarkPowersOf2_2(b *testing.B)        { benchmarkPowersOf2(2, b) }
func BenchmarkPowersOf2_3(b *testing.B)        { benchmarkPowersOf2(3, b) }
func BenchmarkPowersOf2_4(b *testing.B)        { benchmarkPowersOf2(4, b) }
func BenchmarkPowersOf2_5(b *testing.B)        { benchmarkPowersOf2(5, b) }
func BenchmarkPowersOf2_6(b *testing.B)        { benchmarkPowersOf2(6, b) }
func BenchmarkPowersOf2_7(b *testing.B)        { benchmarkPowersOf2(7, b) }
func BenchmarkPowersOf2_8(b *testing.B)        { benchmarkPowersOf2(8, b) }
func BenchmarkPowersOf2_9(b *testing.B)        { benchmarkPowersOf2(9, b) }
func BenchmarkPowersOf2_10(b *testing.B)       { benchmarkPowersOf2(10, b) }
func BenchmarkPowersOf2_20(b *testing.B)       { benchmarkPowersOf2(20, b) }
func BenchmarkPowersOf2_30(b *testing.B)       { benchmarkPowersOf2(30, b) }
func BenchmarkPowersOf2_40(b *testing.B)       { benchmarkPowersOf2(40, b) }
func BenchmarkPowersOf2_50(b *testing.B)       { benchmarkPowersOf2(50, b) }
func BenchmarkPowersOf2_60(b *testing.B)       { benchmarkPowersOf2(60, b) }
func BenchmarkPowersOf2_70(b *testing.B)       { benchmarkPowersOf2(70, b) }
func BenchmarkPowersOf2_80(b *testing.B)       { benchmarkPowersOf2(80, b) }
func BenchmarkPowersOf2_90(b *testing.B)       { benchmarkPowersOf2(90, b) }
func BenchmarkPowersOf2_100(b *testing.B)      { benchmarkPowersOf2(100, b) }
func BenchmarkPowersOf2_200(b *testing.B)      { benchmarkPowersOf2(200, b) }
func BenchmarkPowersOf2_300(b *testing.B)      { benchmarkPowersOf2(300, b) }
func BenchmarkPowersOf2_400(b *testing.B)      { benchmarkPowersOf2(400, b) }
func BenchmarkPowersOf2_500(b *testing.B)      { benchmarkPowersOf2(500, b) }
func BenchmarkPowersOf2_600(b *testing.B)      { benchmarkPowersOf2(600, b) }
func BenchmarkPowersOf2_700(b *testing.B)      { benchmarkPowersOf2(700, b) }
func BenchmarkPowersOf2_800(b *testing.B)      { benchmarkPowersOf2(800, b) }
func BenchmarkPowersOf2_900(b *testing.B)      { benchmarkPowersOf2(900, b) }
func BenchmarkPowersOf2_1000(b *testing.B)     { benchmarkPowersOf2(1000, b) }
func BenchmarkPowersOf2_2000(b *testing.B)     { benchmarkPowersOf2(2000, b) }
func BenchmarkPowersOf2_3000(b *testing.B)     { benchmarkPowersOf2(3000, b) }
func BenchmarkPowersOf2_4000(b *testing.B)     { benchmarkPowersOf2(4000, b) }
func BenchmarkPowersOf2_5000(b *testing.B)     { benchmarkPowersOf2(5000, b) }
func BenchmarkPowersOf2_6000(b *testing.B)     { benchmarkPowersOf2(6000, b) }
func BenchmarkPowersOf2_7000(b *testing.B)     { benchmarkPowersOf2(7000, b) }
func BenchmarkPowersOf2_8000(b *testing.B)     { benchmarkPowersOf2(8000, b) }
func BenchmarkPowersOf2_9000(b *testing.B)     { benchmarkPowersOf2(9000, b) }
func BenchmarkPowersOf2_10000(b *testing.B)    { benchmarkPowersOf2(10000, b) }
func BenchmarkPowersOf2_20000(b *testing.B)    { benchmarkPowersOf2(20000, b) }
func BenchmarkPowersOf2_30000(b *testing.B)    { benchmarkPowersOf2(30000, b) }
func BenchmarkPowersOf2_40000(b *testing.B)    { benchmarkPowersOf2(40000, b) }
func BenchmarkPowersOf2_50000(b *testing.B)    { benchmarkPowersOf2(50000, b) }
func BenchmarkPowersOf2_60000(b *testing.B)    { benchmarkPowersOf2(60000, b) }
func BenchmarkPowersOf2_70000(b *testing.B)    { benchmarkPowersOf2(70000, b) }
func BenchmarkPowersOf2_80000(b *testing.B)    { benchmarkPowersOf2(80000, b) }
func BenchmarkPowersOf2_90000(b *testing.B)    { benchmarkPowersOf2(90000, b) }
func BenchmarkPowersOf2_100000(b *testing.B)   { benchmarkPowersOf2(100000, b) }
func BenchmarkPowersOf2_200000(b *testing.B)   { benchmarkPowersOf2(200000, b) }
func BenchmarkPowersOf2_300000(b *testing.B)   { benchmarkPowersOf2(300000, b) }
func BenchmarkPowersOf2_400000(b *testing.B)   { benchmarkPowersOf2(400000, b) }
func BenchmarkPowersOf2_500000(b *testing.B)   { benchmarkPowersOf2(500000, b) }
func BenchmarkPowersOf2_600000(b *testing.B)   { benchmarkPowersOf2(600000, b) }
func BenchmarkPowersOf2_700000(b *testing.B)   { benchmarkPowersOf2(700000, b) }
func BenchmarkPowersOf2_800000(b *testing.B)   { benchmarkPowersOf2(800000, b) }
func BenchmarkPowersOf2_900000(b *testing.B)   { benchmarkPowersOf2(900000, b) }
func BenchmarkPowersOf2_1000000(b *testing.B)  { benchmarkPowersOf2(1000000, b) }
func BenchmarkPowersOf2_2000000(b *testing.B)  { benchmarkPowersOf2(2000000, b) }
func BenchmarkPowersOf2_3000000(b *testing.B)  { benchmarkPowersOf2(3000000, b) }
func BenchmarkPowersOf2_4000000(b *testing.B)  { benchmarkPowersOf2(4000000, b) }
func BenchmarkPowersOf2_5000000(b *testing.B)  { benchmarkPowersOf2(5000000, b) }
func BenchmarkPowersOf2_6000000(b *testing.B)  { benchmarkPowersOf2(6000000, b) }
func BenchmarkPowersOf2_7000000(b *testing.B)  { benchmarkPowersOf2(7000000, b) }
func BenchmarkPowersOf2_8000000(b *testing.B)  { benchmarkPowersOf2(8000000, b) }
func BenchmarkPowersOf2_9000000(b *testing.B)  { benchmarkPowersOf2(9000000, b) }
func BenchmarkPowersOf2_10000000(b *testing.B) { benchmarkPowersOf2(10000000, b) }
