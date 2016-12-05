package bigo

func PrintUnorderedPairs(array []int) {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			print(array[i], ",", array[j])
		}
	}
}
