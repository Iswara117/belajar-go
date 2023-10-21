package main

import (
	"strconv"
	"fmt"
)

type Player struct {
	Name  string
	Score int
}

func (p Player) AddScore(newScore int) Player {
		p.Score += newScore
		return p
}

func (p Player) DisplayInfo() (string) {
	Score := strconv.Itoa(p.Score)
	
	return "Nama Pemain:" + " " + p.Name + ", "  + " Skor: " + Score
}

func main()  {
	player := Player{
		Name: "John",
		Score: 0,
	}

	player = player.AddScore(10)
	player = player.AddScore(5)
	fmt.Println(player.DisplayInfo())


}