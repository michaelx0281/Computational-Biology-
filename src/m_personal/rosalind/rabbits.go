package main

/*
Input: n months and k (the number of pairs born per month)
Output: The amount of total rabbits!!
*/

func WascallyWabbits(n, k int) {
	arr := make([]int, n)

	for i := 0; i < n+1; i++ {
		if i == 0 {
			arr[0] = 1
			continue
		} else if i == 1 {
			arr[1] = 1
			continue
		}

		if i%2 == 0 {
			arr[i] = arr[i-2] + 3
		}
		if i == 3 {

		}
		if i%2 == 1 {
			arr[i] = arr[i-2] + 3
		}
	}
}

func Mate(index, numRabbits int) int {

}
