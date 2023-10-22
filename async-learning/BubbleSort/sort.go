package main

import "fmt"


//ASC SHORT
func BubbleSort(nums ...int) (sorted []int) {
	// var temp int;

	for i:=0 ; i < len(nums)-1; i++{
		for j:= i+1 ; j < len(nums); j++{
			if nums[i] > nums[j] {
				temp := nums[i]  //temp menyimpan nilai nums dengan index yang di dapat dari i
				nums[i] = nums[j] // nums[i] di re-assign dengan hasil dari nums ber index j 
				nums[j] = temp // num[j] di re-assign dengan value dari temp

				// fmt.Println(nums)
			}
		}
	}

	return nums
}

// // DESC SHORT
// func BubbleSort(nums ...int) (sorted []int) {
// 	// var temp int;

// 	for i:=0 ; i < len(nums)-1; i++{
// 		for j:= i+1 ; j < len(nums); j++{
// 			if nums[i] < nums[j] {
// 				temp := nums[i]  //temp menyimpan nilai nums dengan index yang di dapat dari i
// 				nums[i] = nums[j] // nums[i] di re-assign dengan hasil dari nums ber index j 
// 				nums[j] = temp // num[j] di re-assign dengan value dari temp

// 				// fmt.Println(nums)
// 			}
// 		}
// 	}

// 	return nums
// }

func main(){
	sorted := BubbleSort(3,1,4,2)
    fmt.Println("Sorted array",sorted)
}
