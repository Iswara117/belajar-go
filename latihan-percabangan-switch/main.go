package main

import "fmt"

func main ()  {
	var huruf string
	fmt.Print("masukkan huruf :")
	fmt.Scan(&huruf)
	var result string

	for _, char := range huruf {
        switch char {
        case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
            result = "vokal "
        default:
                result = "konsonan "
            
        }
    }

	println(result)


}