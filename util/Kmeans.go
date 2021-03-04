package util

import (
	"../class"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"syscall"
)

type centroid struct {
	x float64
	y float64
}

func Kmeans(m, n int, edges []class.Edge) [][]int {
	if m > n {
		fmt.Print(errors.New("Clustering number is greater than edges number"))
		return nil
	}

	var centers []int
	for i := 0; i < n; i++ {
		centers = append(centers, i)
	}
	rand.Shuffle(len(centers), func(i, j int) {
		centers[i], centers[j] = centers[j], centers[i]
	})

	// get clusterCenter
	var clusterCenter []centroid
	for i := 0; i < m; i++ {
		clusterCenter = append(clusterCenter, centroid{float64(edges[centers[i]].Lx),
			float64(edges[centers[i]].Ly)})
	}

	var iterNumber = 0
	var clusterings [][]int
	var centerMap = make(map[int][]int)
	for {
		iterNumber++
		for i := 0; i < n; i++ {
			minDistance := float64(syscall.INFINITE)
			var kind int
			for j := 0; j < m; j++ {
				dis := Distance(edges[i], clusterCenter[j])
				if dis < minDistance {
					minDistance = dis
					kind = j
				}
			}
			if _, ok := centerMap[kind]; ok {
				centerMap[kind] = append(centerMap[kind], i)
			} else {
				centerMap[kind] = []int{i}
			}
		}

		newClusterCenter := PointAllNewCenters(centerMap, edges)

		var flag = true
		for i := 0; i < m; i++ {
			if !JudgeEquationOFCenter(clusterCenter[i], newClusterCenter[i]) {
				flag = false
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
		centerMap = make(map[int][]int)
		clusterCenter = newClusterCenter
	}
	return clusterings
}

func JudgeEquationOFCenter(a, b centroid) bool {
	if math.Abs(a.x-b.x) <= class.EPS && math.Abs(a.y-b.y) <= class.EPS {
		return true
	}
	return false
}

func Distance(a class.Edge, b centroid) float64 {
	t1 := float64(a.Lx) - b.x
	t2 := float64(a.Ly) - b.y
	return float64(t1*t1 + t2*t2)
}

func PointAllNewCenters(centerMap map[int][]int, edges []class.Edge) []centroid {
	var clusterCenter []centroid
	for i := 0; i < len(centerMap); i++ {
		clusterCenter = append(clusterCenter, FindNewCenter(centerMap[i], edges))
	}
	return clusterCenter
}

func FindNewCenter(c []int, edges []class.Edge) centroid {
	var sumx, sumy float64 = 0, 0
	for i := 0; i < len(c); i++ {
		sumx += float64(edges[c[i]].Lx)
		sumy += float64(edges[c[i]].Ly)
	}
	return centroid{sumx / float64(len(c)), sumy / float64(len(c))}
}
