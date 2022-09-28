package main

import (
	"fmt"
	"math"
	"sort"
)

type Point struct {
	x float64
	y float64
}

// Take distance between two points
func getDistance(pointA Point, pointB Point) float64 {
	firstDifference := math.Pow(pointA.x-pointB.x, 2)
	secondDifference := math.Pow(pointA.y-pointB.y, 2)
	return math.Sqrt(firstDifference + secondDifference)
}

func findClosestPointsInsideSide(points []Point, leftSideIndex int, rightSideIndex int) (float64, []Point) {
	var minimalDistance float64 = -1
	closestPoints := make([]Point, 2)

	for i := leftSideIndex; i < rightSideIndex+1; i++ {
		for j := i + 1; j < rightSideIndex+1; j++ {
			currentDistance := getDistance(points[i], points[j])
			if currentDistance < minimalDistance {
				minimalDistance = currentDistance
				closestPoints = append(closestPoints, points[i])
				closestPoints = append(closestPoints, points[j])
			}
		}
	}

	return minimalDistance, closestPoints
}

func findClosestPoints(pointsSortedByX []Point, pointsSortedByY []Point, leftSideIndex int, rightSideIndex int) (float64, []Point) {
	leftSidePointsForY := []Point{}
	rightSidePointsForY := []Point{}
	pointsMemory := []Point{}

	var currentDistance float64 = -1
	var minimalDistance float64 = -1
	closestPoints := make([]Point, 2)

	if (rightSideIndex - leftSideIndex + 1) <= 3 {
		currentDistance, closestPoints = findClosestPointsInsideSide(pointsSortedByX, leftSideIndex, rightSideIndex)
		return currentDistance, closestPoints
	}

	middleIndex := (leftSideIndex + rightSideIndex) / 2

	for i := 0; i < len(pointsSortedByY); i++ {
		if pointsSortedByY[i].x < pointsSortedByX[middleIndex].x {
			leftSidePointsForY = append(leftSidePointsForY, pointsSortedByY[i])
		} else {
			rightSidePointsForY = append(rightSidePointsForY, pointsSortedByY[i])
		}
	}

	distance1, closestPoints1 := findClosestPoints(pointsSortedByX, leftSidePointsForY, leftSideIndex, middleIndex)
	distance2, closestPoints2 := findClosestPoints(pointsSortedByX, rightSidePointsForY, middleIndex+1, rightSideIndex)

	if distance1 < distance2 {
		minimalDistance = distance1
		//closestPoints = closestPoints1
		copy(closestPoints, closestPoints1)
	} else {
		minimalDistance = distance2
		//closestPoints = closestPoints2
		copy(closestPoints, closestPoints2)
	}

	for i := 0; i < len(pointsSortedByY); i++ {
		if math.Abs(pointsSortedByY[i].x-pointsSortedByX[middleIndex].x) <= minimalDistance {
			pointsMemory = append(pointsMemory, pointsSortedByY[i])
		}
	}

	distance3 := math.Inf(+1)
	closestPoints3 := []Point{}

	for i := 0; i < len(pointsMemory); i++ {
		for j := i + 1; j < len(pointsMemory); j++ {
			if (pointsMemory[j].y - pointsMemory[i].y) >= minimalDistance {
				break
			}

			currentDistance = getDistance(pointsMemory[i], pointsMemory[j])

			if currentDistance < distance3 {
				distance3 = currentDistance
				closestPoints = append(closestPoints, pointsMemory[i])
				closestPoints = append(closestPoints, pointsMemory[j])
			}
		}
	}

	if minimalDistance > distance3 {
		return distance3, closestPoints3
	}

	return minimalDistance, closestPoints
}

func sortBothArr(points []Point) ([]Point, []Point) {
	fmt.Println("До сортировки: ", points)
	sort.Slice(points, func(i, j int) bool {
		return points[i].x < points[j].x
	})

	pointsSortedByX := make([]Point, len(points))
	copy(pointsSortedByX, points)

	fmt.Println("После сортировки по х: ", pointsSortedByX)

	sort.Slice(points, func(i, j int) bool {
		return points[i].y < points[j].y
	})

	pointsSortedByY := make([]Point, len(points))
	copy(pointsSortedByY, points)

	fmt.Println("После сортировки по y: ", pointsSortedByY)

	return pointsSortedByX, pointsSortedByY
}

func main() {
	input := []Point{{1, 1}, {1, 2}, {2, 1}, {2, 2}, {10, 10}}

	inputSortedByX, inputSortedByY := sortBothArr(input)

	fmt.Println(findClosestPoints(inputSortedByX, inputSortedByY, 0, len(input)-1))

}
