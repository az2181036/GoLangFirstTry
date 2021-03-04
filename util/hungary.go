package util

import (
	"../class"
)

var p = make([]int, class.N_UPPER+1)
var vis = make([]bool, class.N_UPPER+1)

func Hungary(adjacencyList [][]int) []int {
	p = make([]int, class.N_UPPER+1)
	for i := 0; i < len(p); i++ {
		p[i] = -1
	}
	for i := 0; i < len(adjacencyList); i++ {
		vis = make([]bool, class.N_UPPER+1)
		match(i, adjacencyList)
	}
	return p
}

func match(i int, adjacencyList [][]int) bool {
	for j := 0; j < len(adjacencyList[i]); j++ {
		idx := adjacencyList[i][j]
		if !vis[idx] {
			vis[idx] = true
			if p[idx] == -1 || match(p[idx], adjacencyList) {
				p[idx] = i
				return true
			}
		}
	}
	return false
}
