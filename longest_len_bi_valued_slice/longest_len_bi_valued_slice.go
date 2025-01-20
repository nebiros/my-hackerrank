package longest_len_bi_valued_slice

import "math"

// longestLen returns the length of the longest bi-valued slice
//
// We call an array bi-valued if it contains at most two different numbers. For example, [4, 5, 5, 4, 5] is bi-valued because it contains two different numbers: 4 and 5. However, [1, 5, 4, 5] is not bi-valued because it contains three different numbers.
// What is the length of the longest bi-valued slice (continuous fragment) in a given array A?
// Write a function:
// def solution(A): # Your code here
// that, given an array A consisting of N integers, returns the length of the longest bi-valued slice in A.
//
// Examples:
//
//  1. Given A = [4, 2, 2, 4, 2], the function should return 5, because the whole array is bi-valued.
//  2. Given A = [1, 2, 3, 2], the function should return 3. The longest bi-valued slice is [2, 3, 2].
//  3. Given A = [0, 5, 4, 4, 5, 1], the function should return 4. The longest bi-valued slice is [5, 4, 4, 5].
//  4. Given A = [4, 4], the function should return 2.
//
// Assume that:
//   - N is an integer within the range [1..100].
//   - Each element of array A is an integer within the range [-1,000,000,000..1,000,000,000].In your solution, focus on correctness. The performance of your solution will not be the focus of the assessment.
func longestLen(A []int) int {
	var (
		lbs  int
		temp int
		lsr  int

		ls  = -1
		sls = -1
	)

	for _, v := range A {
		if v == ls || v == sls {
			temp++
		} else {
			// if the current number is not in our read list it means new series has started, tempCounter value in this case will be
			// how many times lastSeen number repeated before this new number encountered + 1 for current number
			temp = lsr + 1
		}

		if v == ls {
			lsr++
		} else {
			lsr = 1

			sls = ls
			ls = v
		}

		lbs = int(math.Max(float64(temp), float64(lbs)))
	}

	return lbs
}
