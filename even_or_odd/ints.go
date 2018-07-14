package main

type ints []int

func intSlice(num int) ints {
	integers := ints{}

	for i := 1; i <= num; i++ {
		integers = append(integers, i)
	}

	return integers
}
