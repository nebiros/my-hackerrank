package quick_sort_rand_pivot

import "math/rand"

func quickSort(arr []int) []int {
	length := len(arr)

	if length <= 1 {
		cp := make([]int, length)
		copy(cp, arr)
		return cp
	}

	pivot := arr[rand.Intn(length)]

	less := make([]int, 0, length)
	middle := make([]int, 0, length)
	more := make([]int, 0, length)

	for _, item := range arr {
		switch {
		case item < pivot:
			less = append(less, item)

		case item == pivot:
			middle = append(middle, item)

		case item > pivot:
			more = append(more, item)
		}
	}

	less, more = quickSort(less), quickSort(more)

	less = append(less, middle...)
	less = append(less, more...)

	return less
}
