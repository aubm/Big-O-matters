package bigo

func PowersOf2(n int) int {
	if n < 1 {
		return 0
	} else if n == 1 {
		print(1)
		return 1
	} else {
		prev := PowersOf2(n / 2)
		curr := prev * 2
		print(curr)
		return curr
	}
}
