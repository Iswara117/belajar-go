package main

import "fmt"

func main (){

	var str1 string;

	for i := 0; i <= 5; i++ {
		for j := 0; j < 5-i; j++ {
			str1 += ""
		}

		for j := 5; j > i; j-- {
			str1 += "*"
		}
		str1 += "\n"
	}

	fmt.Print(str1)


	fmt.Print("\n")

	 var str2 string;
	for i := 0; i <= 5; i++ {
		for j := 0; j < 5-i; j++ {
			str2 += ""
		}

		for j := 0; j < i; j++ {
			str2 += "*"
		}
		str2 += "\n"
	}

	fmt.Print(str2)

	for i := 1; i <= 50; i++ {
        output := ""

        if i%3 == 0 {
            output += "Noo"
        }

        if i%5 == 0 {
            output += "Bee"
        }

        if output == "" {
            output = fmt.Sprint(i)
        }

        if i == 50 {
            fmt.Print(output)
        } else {
            fmt.Print(output, ", ")
        }
    }
	
}