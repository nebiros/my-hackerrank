package quick_sort

// Quicksort, https://en.wikipedia.org/wiki/Quicksort.
// Based on https://www.bbminfo.com/golang/ds/sort/quick-sort.php.
func quickSort(arr []int, low int, high int) {
	if low < high {
		pivot := doPivot(arr, low, high)

		quickSort(arr, low, pivot)
		quickSort(arr, pivot+1, high)
	}
}

func doPivot(arr []int, low int, high int) int {
	pivot := arr[low]

	i := low
	j := high

	for i < j {
		for arr[i] <= pivot && i < high {
			i++
		}

		for arr[j] > pivot && j > low {
			j--
		}

		if i < j {
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}

	arr[low] = arr[j]
	arr[j] = pivot

	return j
}
