package util

import (
	"../class"
	"fmt"
	"math/rand"
	"syscall"
)

func Clustering(m, n int, edges []class.Edge) [][]class.Edge {
	var flag = true
	var centers []int
	var centerMap = make(map[int][]class.Edge)
	var iterNumber = 0
	if m > n {
		m = n
	}
	for i := 0; i < n; i++ {
		centers = append(centers, i)
	}
	rand.Shuffle(len(centers), func(i, j int) {
		centers[i], centers[j] = centers[j], centers[i]
	})

	// get clusterCenter
	var clusterCenter []class.Edge
	for i := 0; i < m; i++ {
		clusterCenter = append(clusterCenter, edges[centers[i]])
	}

	var clusterings [][]class.Edge
	var cluster_idx, edge_idx, map_idx int
	for ; ; iterNumber++ {

		var maxDistance = float64(-1)
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
			if minDistance > maxDistance {
				maxDistance = minDistance
				cluster_idx = kind
				edge_idx = i
				map_idx = len(centerMap[kind]) + 1
			}
			if _, ok := centerMap[kind]; ok {
				centerMap[kind] = append(centerMap[kind], edges[i])
			} else {
				centerMap[kind] = []class.Edge{edges[i]}
			}
		}
		if maxDistance > class.SIGNAL_RANGE/2.0 {
			centerMap[cluster_idx] = append(centerMap[cluster_idx][:map_idx], centerMap[cluster_idx][map_idx+1:]...)
			centerMap[len(centerMap)+1] = []class.Edge{edges[edge_idx]}
		}
		newClusterCenter := PointAllNewCenters(centerMap)
		if len(newClusterCenter) == len(clusterCenter) {
			for i := 0; i < m; i++ {
				if !JudgeEquationOFCenter(clusterCenter[i], newClusterCenter[i]) {
					flag = false
					break
				}
			}
		}
		if flag {
			fmt.Println("The iteration number of K-means is:", iterNumber)
			fmt.Println("The cluster number is:", len(centerMap))
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
