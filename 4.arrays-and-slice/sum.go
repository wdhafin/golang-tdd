package arraysandslice

func Sum(number [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += number[i]
	}
	return sum
}
