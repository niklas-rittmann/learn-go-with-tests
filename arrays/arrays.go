package arrays

// Return the sum over a slice of integers
func Sum(arr []int) (sum int) {
	for _, num := range arr {
		sum += num
	}
	return
}

// Return the sum over slices of integers
func SumAll(arr ...[]int) (sumOfArrs []int) {
	for _, arr := range arr {
		sumOfArrs = append(sumOfArrs, Sum(arr))
	}
	return
}

// Return the sum over all slices excluding the first element
func SumAllTails(arr ...[]int) (sumOfArrs []int) {
	for _, arr := range arr {
		if len(arr) == 0 {
			sumOfArrs = append(sumOfArrs, 0)

		} else {
			sumOfArrs = append(sumOfArrs, Sum(arr[1:]))
		}
	}
	return
}
