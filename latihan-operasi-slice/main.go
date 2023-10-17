package main

import "fmt"

func main (){

	myslice := []int{10,20,30,40,50}
	fmt.Println("====[mySlice]====")
	fmt.Println("Data:",myslice)
	fmt.Println("Len:",len(myslice))
	fmt.Println("Cap:",cap(myslice))

	slice1 := myslice[0:3]
	fmt.Println("====[Slice1]====")
	fmt.Println("Data:",slice1)
	fmt.Println("Len:",len(slice1))
	fmt.Println("Cap:",cap(slice1))

	slice2 := myslice[3:5]
	fmt.Println("====[Slice2]====")
	fmt.Println("Data:",slice2)
	fmt.Println("Len:",len(slice2))
	fmt.Println("Cap:",cap(slice2))

	append_Slice1 := append(slice1,60)

	fmt.Println("====[Setelah Append ke Slice1]====")
	fmt.Println("Data Slice 1:",append_Slice1)
	fmt.Println("Len:",len(append_Slice1))
	fmt.Println("Cap:",cap(append_Slice1))

	append_Slice2 := append(slice2,70)

	fmt.Println("====[Setelah Append ke Slice2]====")
	fmt.Println("Data Slice 2:",append_Slice2)
	fmt.Println("Len:",len(append_Slice2))
	fmt.Println("Cap:",cap(append_Slice2))
	
}