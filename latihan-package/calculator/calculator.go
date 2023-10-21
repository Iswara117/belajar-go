package calculator


type Calculate interface {
	multiply() int
	Add() int
}

type Num1 struct {
	Num     int
}


func NewNum1(num int) Calculate {
	return &Num1{Num: num}
}

func (n Num1) Add() int {
	return n.Num + n.Num
}

func (n Num1) multiply() int {
	return n.Num + n.Num
}