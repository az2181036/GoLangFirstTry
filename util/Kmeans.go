package util

import (
	"errors"
	"fmt"
	"github.com/user/class"
	"math"
	"math/rand"
	"syscall"
)

func Kmeans(m, n int, clusterCenter []class.Edge, edges []class.Edge,
	clusterings [][]class.Edge) ([][]class.Edge){
	var flag = false
	var centers []int
	var centerMap = make(map[int][]class.Edge)
	var iterNumber = 0
	if m > n {
		fmt.Print(errors.New("Clustering number is greater than edges number"))
		return nil
	}
	for i := 0; i < n; i++ {
		centers = append(centers, i)
	}
	rand.Shuffle(len(centers), func(i, j int) {
		centers[i], centers[j] = centers[j], centers[i]
	})

	// get clusterCenter
	for i := 0; i < m; i++ {
		clusterCenter = append(clusterCenter, edges[centers[i]])
	}
	for {
		iterNumber++
		for i := 0; i < n; i++ {
			minDistance := syscall.INFINITE
			var kind int
			for j := 0; j < m; j++ {
				dis := Distance(edges[i], clusterCenter[j])
				if dis < minDistance {
					minDistance = dis
					kind = j
				}
			}
			if _, ok := centerMap[kind]; ok {
				centerMap[kind] = append(centerMap[kind], edges[i])
			} else {
				centerMap[kind] = []class.Edge{edges[i]}
			}
		}
		newClusterCenter := PointAllNewCenters(m, centerMap)
		for i := 0; i < m; i++ {
			if JudgeEquationOFCenter(clusterCenter[i], newClusterCenter[i]) {
				flag = true
				break
			}
		}
		if flag {
			fmt.Println("The iteration number of K-means is:", iterNumber)
			for i := 0; i < m; i++ {
				clusterings = append(clusterings, centerMap[i])
			}
			break
		}
		centerMap = make(map[int][]class.Edge)
		clusterCenter = newClusterCenter
	}
	return clusterings
}

func JudgeEquationOFCenter(a, b class.Edge) bool {
	if math.Abs(float64(a.Lx-b.Lx)) <= class.EPS && math.Abs(float64(a.Ly-b.Ly)) <= class.EPS {
		return true
	}
	return false
}

func Distance(a, b class.Edge) int {
	t1 := a.Lx - b.Lx
	t2 := a.Ly - b.Ly
	return t1*t1 + t2*t2
}

func PointAllNewCenters(m int, centerMap map[int][]class.Edge) []class.Edge {
	var clusterCenter []class.Edge
	for i := 0; i < m; i++ {
		clusterCenter = append(clusterCenter, FindNewCenter(centerMap[i]))
	}
	return clusterCenter
}

func FindNewCenter(c []class.Edge) class.Edge {
	sumx, sumy := 0, 0
	for i := 0; i < len(c); i++ {
		sumx += c[i].Lx
		sumy += c[i].Ly
	}
	return class.Edge{sumx/len(c),sumy/len(c),[]class.Task{}, []class.Task{},
		[]class.Task{}, 0,0}
}
