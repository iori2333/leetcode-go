package utils

func intRange(start, end, step int) []int {
	if step == 0 {
		panic("step cannot be zero")
	}

	if (step > 0 && start >= end) || (step < 0 && start <= end) {
		return nil
	}

	n := (end - start) / step
	if n < 0 {
		n = -n
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = start + i*step
	}

	return res
}

func Range(indices ...int) []int {
	if len(indices) == 1 {
		return intRange(0, indices[0], 1)
	}

	if len(indices) == 2 {
		return intRange(indices[0], indices[1], 1)
	}

	if len(indices) == 3 {
		return intRange(indices[0], indices[1], indices[2])
	}

	panic("invalid indices")
}
