package sum

func Sum(values []int) int {
	var sum int
	for _, number := range values {
		sum += number
	}
	return sum
}
func SumAll(numbersToSum ...[]int) []int {

	// legthOfNumbers := len(numbersToSum)
	// sums := make([]int, legthOfNumbers)
	// for i, numbers := range numbersToSum {
	// 	sums[i] = Sum(numbers)
	// }

	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
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
