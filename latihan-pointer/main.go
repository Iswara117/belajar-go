package main

import "fmt"

type Student struct {
	Name  string
	Class string
}

func (s *Student) SetMyName(name *string){
	s.Name = *name
}

func (s *Student) CallMyName(){
	fmt.Println("Hello, My Name is", s.Name)
}

func main()  {
	
	student := Student{
		Name : "Noobee",
		Class : "Basic Go",
	} 

	newName := "Anak Agung Gede Iswara Wijaya"

	student.SetMyName(&newName)
	student.CallMyName()
}