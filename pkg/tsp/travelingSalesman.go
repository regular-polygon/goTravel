package tsp

import (
	"math"
)

func BruteForceTSP(points []DDCoordinates, distanceFunc func(p1, p2 DDCoordinates) float64) ([]DDCoordinates, float64) {
	n := len(points)
	bestPath := make([]DDCoordinates, n)
	bestDist := math.Inf(1)

	permutation := make([]int, n)
	for i := 0; i < n; i++ {
		permutation[i] = i
	}

	for {
		var path []DDCoordinates
		var dist float64

		for i := 0; i < n-1; i++ {
			path = append(path, points[permutation[i]])
			dist += distanceFunc(points[permutation[i]], points[permutation[i+1]])
		}
		path = append(path, points[permutation[n-1]])
		dist += distanceFunc(points[permutation[n-1]], points[permutation[0]])

		if dist < bestDist {
			copy(bestPath, path)
			bestDist = dist
		}

		if !nextPermutation(permutation) {
			break
		}
	}

	return bestPath, bestDist
}

func nextPermutation(permutation []int) bool {
	n := len(permutation)
	i := n - 2
	for i >= 0 && permutation[i] >= permutation[i+1] {
		i--
	}
	if i < 0 {
		return false
	}
	j := n - 1
	for permutation[j] <= permutation[i] {
		j--
	}
	permutation[i], permutation[j] = permutation[j], permutation[i]
	reverse(permutation[i+1:])
	return true
}

func reverse(slice []int) {
	n := len(slice)
	for i := 0; i < n/2; i++ {
		slice[i], slice[n-1-i] = slice[n-1-i], slice[i]
	}
}
