package main

import "fmt"

type Book struct {
	Title string
	Author string
}

func main(){
	
		var book1 Book
		book1.Title = "Pemograman Go"
		book1.Author = "John Doe"

		fmt.Println("Informasi tentang Book 1:")
		fmt.Println("Judul:", book1.Title)
		fmt.Println("Penulis:", book1.Author)

		var book2 Book
		book2.Title = "Algoritma dan Struktur Data"
		book2.Author = "Jane Smith"

		fmt.Print("\n")
		fmt.Println("Informasi tentang Book 2:")
		fmt.Println("Judul:", book2.Title)
		fmt.Println("Penulis:", book2.Author)


		var book3 Book
		book3.Title = "Machine Learning for Beginners"
		book3.Author = "David Johson"

		fmt.Print("\n")
		fmt.Println("Informasi tentang Book 3:")
		fmt.Println("Judul:", book3.Title)
		fmt.Println("Penulis:", book3.Author)

}