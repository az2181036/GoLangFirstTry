package main

import (
	"fmt"
)

type Position struct {
	x,y int
}

func main(){
	var s = make([]Position,0,2)
	var a = []Position{
		{1, 2},
		{2, 3},
	}
	s = append(s, a[0])
	fmt.Println(s)
	a[0].x = 3
	fmt.Println(a, s)
}
