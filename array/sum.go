package array

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}

func Reverse(numbers []int) []int {
	for i := 0; i < len(numbers)/2; i++ {
		numbers[i], numbers[len(numbers)-i-1] = numbers[len(numbers)-i-1], numbers[i]
	}
	return numbers
}
