package main

import "fmt"


func binary(find int, nums []int) (interface{}){
	// left := nums[:middleIndex]
	// right := nums[middleIndex:len(nums)]
	
	for{
		middleIndex := len(nums) / 2
		if nums[middleIndex] == find {
		return true
	}else if find < nums[middleIndex] {
		nums = nums[:middleIndex]
	} else if find > nums[middleIndex] {
		nums = nums[middleIndex:]
	}

	if middleIndex == 0 {
		break
	}}
	
	return false

}


func main (){
	sortedArray := []int{1,2,3,4,5,6,7,8}
    find := 8

	fmt.Println(binary(find, sortedArray))
}