package main

import "fmt"

func main() {
	var i int
	fmt.Scan(&i)
	var result string

	if i % 2 != 0 { result = "ganjil" } else { result = "genap" }

	println(result)
  }