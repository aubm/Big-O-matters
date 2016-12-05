package bigo

func SumAndProduct(array []int) {
	sum := 0
	product := 1
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	for i := 0; i < len(array); i++ {
		product *= array[i]
	}
}
