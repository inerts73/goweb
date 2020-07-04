package main

import "fmt"

func removeDuplicates(nums []int) int {
	i := 0
	j := 1
	for i < len(nums) - 1 {
		for j < len(nums) {
			if nums[i] == nums[j] {
				nums = append(nums[:j], nums[j+1:]...)		
			} else {
				i++
				j++
			}	
		}
	}
	
	// fmt.Println(nums)
	return len(nums)
}

func main(){
	input := []int{0,0,1,1,1,2,2,3,3,4}
	// input = []int{1,1,2}
	// input = append(input[:2], input[3:]...)
	// fmt.Println(len(input))
	fmt.Println(removeDuplicates(input))
}