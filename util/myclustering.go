package util

import (
	"../class"
	"fmt"
	"math/rand"
	"syscall"
)

func Clustering(m int, edges []class.Edge) ([][]class.Edge, []Centroid) {
	if m > len(edges) {
		m = len(edges)
	}

	var centers []int
	for i := 0; i < len(edges); i++ {
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
	var cluster_idx, edge_idx, map_idx int

	for ; ; iterNumber++ {
		var maxDistance = float64(-1)
		var centerMap = make(map[int][]class.Edge)
		for i := 0; i < len(edges); i++ {
			minDistance := float64(syscall.INFINITE)
			var kind int
			for j := 0; j < len(clusterCenter); j++ {
				dis := Distance(edges[i], clusterCenter[j])
				if dis < minDistance {
					minDistance = dis
					kind = j
				}
			}
			if minDistance > maxDistance {
				maxDistance = minDistance
				cluster_idx = kind
				edge_idx = i
				map_idx = len(centerMap[kind])
			}
			if _, ok := centerMap[kind]; ok {
				centerMap[kind] = append(centerMap[kind], edges[i])
			} else {
				centerMap[kind] = []class.Edge{edges[i]}
			}
		}
		if maxDistance > class.SIGNAL_RANGE*class.SIGNAL_RANGE/4.0 {
			centerMap[cluster_idx] = append(centerMap[cluster_idx][:map_idx], centerMap[cluster_idx][map_idx+1:]...)
			centerMap[len(centerMap)] = []class.Edge{edges[edge_idx]}
		}

		for key := range centerMap {
			if len(centerMap[key]) == 0 {
				delete(centerMap, key)
			}
		}

		newClusterCenter := PointAllNewCenters(centerMap)

		var flag = true
		if len(newClusterCenter) == len(clusterCenter) {
			for i := 0; i < len(clusterCenter); i++ {
				if !JudgeEquationOFCenter(clusterCenter[i], newClusterCenter[i]) {
					flag = false
					break
				}
			}
		} else {
			flag = false
		}
		if flag {
			fmt.Println("The iteration number of My Clustering Algorithm is:", iterNumber)
			fmt.Println("The cluster number is:", len(centerMap))
			for i := 0; i < len(clusterCenter); i++ {
				clusterings = append(clusterings, centerMap[i])
			}
			break
		}
		clusterCenter = newClusterCenter
	}
	return clusterings, clusterCenter
}
