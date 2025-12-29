package mathutil

// Sum returns the sum of all provided integers.
func Sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Average returns the average of the provided integers.
// Returns 0 if no numbers are provided.
func Average(nums ...int) float64 {
	if len(nums) == 0 {
		return 0
	}
	return float64(Sum(nums...)) / float64(len(nums))
}

// Max returns the maximum value from the provided integers.
// Panics if no numbers are provided.
func Max(nums ...int) int {
	if len(nums) == 0 {
		panic("Max requires at least one argument")
	}
	max := nums[0]
	for _, n := range nums[1:] {
		if n > max {
			max = n
		}
	}
	return max
}
