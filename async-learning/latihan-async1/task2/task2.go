package main

import (
	"sort"
	"fmt"
	// "strconv"
	)
type Examination struct{
    Name string
    Score int
}


func median(v []Examination) interface{} {
	scoreComparison := func(i, j int) bool {
		return v[i].Score < v[j].Score
	}

	sort.Slice(v, scoreComparison)

	middleIndex := len(v) / 2

	if middleIndex == 0 {
		mid1 := v[len(v)/2-1].Score
		mid2 := v[len(v)/2].Score
		return (mid1 + mid2) / 2
	}

	return v[middleIndex].Score
}

func min(v []Examination) interface{} {
	scoreComparison := func(i, j int) bool {
		return v[i].Score < v[j].Score
	}

	sort.Slice(v, scoreComparison)

	return v[0].Score
}

func max(v []Examination) interface{} {
	scoreComparison := func(i, j int) bool {
		return v[i].Score < v[j].Score
	}

	sort.Slice(v, scoreComparison)

	return v[len(v)-1].Score
}

func average(v []Examination) float64 {
	var countValue int
	fmt.Println(countValue)
	for  _,n := range v {
		countValue += n.Score
    }

	fmt.Println(countValue)
	return float64(countValue) / float64(len(v))
}

func binarySearchGreaterThan80(v []Examination, target int) int {
	left, right := 0, len(v)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2

		if v[mid].Score > target {
			result = mid
			right = mid - 1 
		} else {
			left = mid + 1
		}
	}

	return result
}

func main(){
	var examResult = []Examination{
		{
			Name: "Budi",
			Score: 90,
		},
		{
			Name: "Andi",
			Score: 85,
		},
		{
			Name: "Nayla",
			Score: 87,
		},
		{
			Name: "Danu",
			Score: 80,
		},
		{
			Name: "Rahmat",
			Score: 75,
		},
		{
			Name: "Erika",
			Score: 83,
		},
		{
			Name: "Siska",
			Score: 87,
		},
		{
			Name: "Mita",
			Score: 94,
		},
		{
			Name: "Shinta",
			Score: 82,
		},
		{
			Name: "Jojo",
			Score: 83,
		},
	}

	fmt.Println(median(examResult))
	fmt.Println(max(examResult))
	fmt.Println(min(examResult))
	fmt.Println(average(examResult))
	fmt.Println(binarySearchGreaterThan80(examResult, 80))
	
}