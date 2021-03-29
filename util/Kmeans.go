package util

import (
	"../class"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"syscall"
)

type Centroid struct {
	X float64
	Y float64
}

func Kmeans(m, n int, edges []class.Edge) [][]class.Edge {
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
	var clusterCenter []Centroid
	for i := 0; i < m; i++ {
		clusterCenter = append(clusterCenter, Centroid{float64(edges[centers[i]].Lx),
			float64(edges[centers[i]].Ly)})
	}

	var iterNumber = 0
	var clusterings [][]class.Edge
	for ; ; iterNumber++ {
		var centerMap = make(map[int][]class.Edge)
		for i := 0; i < n; i++ {
			minDistance := float64(syscall.INFINITE)
			var kind int
			for j := 0; j < len(clusterCenter); j++ {
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

		newClusterCenter := PointAllNewCenters(centerMap)

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
		clusterCenter = newClusterCenter
	}
	for i:=0; i<len(clusterings); i++{
		for j:=0; j<len(clusterings[i]); j++ {
			clusterings[i][j].Cluster = i
		}
	}
	return clusterings
}

func JudgeEquationOFCenter(a, b Centroid) bool {
	if math.Abs(a.X-b.X) <= class.EPS && math.Abs(a.Y-b.Y) <= class.EPS {
		return true
	}
	return false
}

func Distance(a class.Edge, b Centroid) float64 {
	t1 := float64(a.Lx) - b.X
	t2 := float64(a.Ly) - b.Y
	return t1*t1 + t2*t2
}

func PointAllNewCenters(centerMap map[int][]class.Edge) []Centroid {
	var clusterCenter []Centroid
	for i := 0; i < len(centerMap); i++ {
		clusterCenter = append(clusterCenter, FindNewCenter(centerMap[i]))
	}
	return clusterCenter
}

func FindNewCenter(c []class.Edge) Centroid {
	sumx, sumy := 0.0, 0.0
	for i := 0; i < len(c); i++ {
		sumx += float64(c[i].Lx)
		sumy += float64(c[i].Ly)
	}
	return Centroid{sumx / float64(len(c)), sumy / float64(len(c))}
}
