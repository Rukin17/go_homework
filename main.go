package main

import (
	"fmt"
	"strings"
)

func main() {
	nums := GetNums()
	BubbleSortedReverse(nums)
	res := fmt.Sprint(nums)
	res = strings.Trim(res, "[]")
	fmt.Printf("Отсортированные элементы: %s\n", res)
	fmt.Printf("Самое большое число: %d\n", nums[0])
	fmt.Printf("Самое маленькое число: %d\n", nums[4])
	fmt.Printf("Среднее-арифметическое: %.1f\n", GetAverage(nums))
}

func BubbleSortedReverse(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] < nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func GetAverage(nums []int) float64 {
	var avg float64
	sum := GetSum(nums)
	avg = float64(sum) / float64(len(nums))
	return avg

}

func GetSum(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

func GetNums() []int {
	nums := []int{}
	for i := 0; i < 5; i++ {
		var num int
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	return nums
}
