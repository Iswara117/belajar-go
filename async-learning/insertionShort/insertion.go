package main

import "fmt"


func InsertionSort(nums ...int) (sorted []int){
	for i:=0 ; i < len(nums); i++{
		for j:= i ; j >0; j--{  // for untuk mendapatkan index 1 langkah lebih depan
			if nums[j] < nums[j-1]{ 
				temp := nums[j-1]  
				nums[j-1] = nums[j]
				nums[j] = temp
			}

		}
	}
	return nums
}

func main(){
	sorted := InsertionSort(6,5,3,1,8,7,2,4)
    fmt.Println(sorted)
}