/**
 * Programming Languages - Assignment #1
 *
 * Alan Herculano Diniz
 *
 * Solve grouping problem with the leader algorithm
 *
 * leader.go: implementing the actual leader algorithm
*/

package main

import "math"

/**
 * Calculating the euclidian distance sum and the point groups
 *
 * Inputs: the limit distance and the point slice
 *
 * Output: the euclidian distance sum and the point groups
 *
 * Observations:
 *
 * 1 - In the group slice, the first position of each group is the leader of that group
*/
func calculateResults(limit float64, points [][]float64) (float64, [][]int) {

	// Creating and initializing the group slice:
	var groups [][]int
	var group0 []int
	group0 = append(group0, 1)
	groups = append(groups, group0)

	// Iterating through the point slice:
	for i := 1; i < len(points); i++ {

		leader := true

		// Checking if the current point is a leader or not and treating it accordingly:
		for j := 0; j < len(groups); j++ {

			if pointDistance(points[groups[j][0] - 1], points[i]) <= limit {

				groups[j] = append(groups[j], i + 1)
				leader = false
				break
			}
		}

		// Adding a new group if the point is a leader:
		if leader {

			var group []int
			group = append(group, i + 1)
			groups = append(groups, group)
		}
	}

	sse := calculateSSE(points, groups)

	return sse, groups
}

/**
 * Calculating the distance between two points
*/
func pointDistance(a, b []float64) float64 {

	var distance float64 = 0
	for i := 0; i < len(a); i++ {

		distance += (a[i] - b[i]) * (a[i] - b[i])
	}
	distance = math.Sqrt(distance)

	return distance
}

/**
 * Calculating the group's mass center
 *
 * Input: the points slice and the positions that must be evaluated
 *
 * Output: the center point of the group
*/
func centerOfMass(points [][]float64, group []int) []float64 {

	dimension := len(points[0]) // The amount of coordinates in the point
	groupSize := len(group) // The amount of points in the group
	center := make([]float64, dimension) // Building the slice that represents the result point

	// Initializing the center of mass point:
	for i := 0; i < dimension; i++ {
		center[i] = 0
	}

	// Calculating the group's center of mass:
	for i := 0; i < groupSize; i++ {

		point := points[group[i] - 1] // Accessing the point in the group

		// Adding the coordinate to the center of mass:
		for j := 0; j < dimension; j++ {
			center[j] += point[j] / float64(groupSize)
		}
	}

	return center
}

/**
 * Calculating the sum of euclidian distances
 *
 * Inputs: the slices of points and groups
*/
func calculateSSE(points [][]float64, groups [][]int) float64 {

	var sse float64 = 0

	for i := 0; i < len(groups); i++ {

		center := centerOfMass(points, groups[i])
		for j := 0; j < len(groups[i]); j++ {

			sse += pointDistance(points[groups[i][j] - 1], center) * pointDistance(points[groups[i][j] - 1], center)
		}
	}

	return sse
}
