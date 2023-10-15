package main

import "fmt"

func countPairs(nums []int, k int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]%k == 0 || nums[j]%k == 0 {
				res++
			}
		}
	}
	return res
}

func countPairs2(nums []int, k int) int64 {
	dp := make([]int64, len(nums))
	single := make([]int64, len(nums))
	dp[0] = 0
	if nums[0]%k == 0 {
		single[0] = 1
	} else {
		single[0] = 0
	}
	if nums[1]%k == 0 {
		single[1] = single[0] + 1
	} else {
		single[1] = single[0]
	}
	if single[1] == 2 {
		dp[1] = 1
	}
	for i := 2; i < len(nums); i++ {
		if k < nums[i] {

			if nums[i]%k == 0 {
				dp[i] = dp[i-1] + int64(i) + 1
				single[i] = single[i-1] + 1
			} else {
				dp[i] = dp[i-1] + single[i-1]
				single[i] = single[i-1]
			}
			// fmt.Printf("index: %d, dp: %d, s: %d \n", i, dp[i], single[i])
		}
	}
	return dp[len(nums)-1]
}

func countPairs3(nums []int, k int) int64 {
	m := make(map[int]int64, 0)
	var ans int64
	for i := 0; i < len(nums); i++ {
		num1 := findGCD(nums[i], k)
		for num2, fre := range m {
			if (num1*num2)%k == 0 {
				ans = ans + fre
			}
		}
		m[num1] = m[num1] + 1
	}
	return ans
}

func findGCD(x, y int) int {
	if y > x {
		return findGCD(y, x)
	}
	if y == 0 {
		return x
	} else {
		return findGCD(y, x%y)
	}
}

func main() {
	fmt.Println(countPairs([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(countPairs([]int{1, 2, 3, 4}, 5))
	fmt.Println(countPairs2([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(countPairs2([]int{1, 2, 3, 4}, 5))
	fmt.Println(countPairs3([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(countPairs3([]int{1, 2, 3, 4}, 5))
	fmt.Println(findGCD(3, 8))
}
