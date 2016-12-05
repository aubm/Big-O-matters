package bigo

func PrintPairs(array []int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			print(array[i], ",", array[j])
		}
	}
}
