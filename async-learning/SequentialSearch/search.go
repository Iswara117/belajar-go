package main

import "fmt"


func Sequential(find int, nums ...int) (found bool, index int){
	for i, v := range nums {
        if v == find {
			found = true
			index = i
			return // jika sudah mendapatkan hasil kita perlu return untuk menghentikan for dan mengeluarkan
		} else {
			index = -1
		}
		
    }


	return
}

func main(){
	find := 2

    // setup list angka
	nums := []int{4, 2, 1, 2, 4, 8, 5, 10, 4, 9, 4}

    // panggil function Sequential
    found,index := Sequential(find, nums...)
    fmt.Println(found, index)
}