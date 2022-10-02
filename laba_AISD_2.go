package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

type Point struct {
	x float64
	y float64
}

type Distance struct {
	distance            float64
	coordinatesOfPoints []Point
}

// Take distance between two points
func getDistance(pointA Point, pointB Point) float64 {
	firstDifference := math.Pow(pointA.x-pointB.x, 2)
	secondDifference := math.Pow(pointA.y-pointB.y, 2)
	return math.Sqrt(firstDifference + secondDifference) //this is a classic formula for finding the length of a segment
}

// Func which count all segments in one side
func findClosestPointsInsideSide(points []Point, leftSideIndex int, rightSideIndex int) (float64, []Point) {
	var minimalDistance = math.Inf(+1) // create max value to compare with
	closestPoints := make([]Point, 2)  // array of two points, between which we are looking for a distance

	for i := leftSideIndex; i < rightSideIndex+1; i++ { //iterating through all points
		for j := i + 1; j < rightSideIndex+1; j++ {
			currentDistance := getDistance(points[i], points[j]) // get current distance
			if currentDistance < minimalDistance {               // finding minimal distance in current side
				closestPoints = make([]Point, 0)                            // create empty arr
				minimalDistance = currentDistance                           // memorize minimal value
				closestPoints = append(closestPoints, points[i], points[j]) // add to array of points our points
			}
		}
	}
	return minimalDistance, closestPoints
}

// Func which count all segments by algorithm
func findClosestPoints(pointsSortedByX []Point, pointsSortedByY []Point, leftSideIndex int, rightSideIndex int) (float64, []Point) {
	var leftSidePointsForY []Point  //memorize all points which is in left side
	var rightSidePointsForY []Point //memorize all points which is in right side
	var pointsMemory []Point        // arr for memorize points

	var currentDistance float64 = -1
	var minimalDistance float64 = -1
	closestPoints := make([]Point, 2) // array for containing two closest Points

	if (rightSideIndex - leftSideIndex + 1) <= 3 { // If the points are less than 4, just count all of segments
		currentDistance, closestPoints = findClosestPointsInsideSide(pointsSortedByX, leftSideIndex, rightSideIndex)
		return currentDistance, closestPoints
	}

	middleIndex := (leftSideIndex + rightSideIndex) / 2 // middle index of points in current side
	for i := 0; i < len(pointsSortedByY); i++ {         // dividing points into two side
		if pointsSortedByY[i].x < pointsSortedByX[middleIndex].x {
			leftSidePointsForY = append(leftSidePointsForY, pointsSortedByY[i])
		} else {
			rightSidePointsForY = append(rightSidePointsForY, pointsSortedByY[i])
		}
	}
	distanceArray := []Distance{}

	distanceStruct := Distance{}

	distanceStruct.distance, distanceStruct.coordinatesOfPoints = findClosestPoints(pointsSortedByX, leftSidePointsForY, leftSideIndex, middleIndex)
	distanceArray = append(distanceArray, distanceStruct)

	distanceStruct.distance, distanceStruct.coordinatesOfPoints = findClosestPoints(pointsSortedByX, rightSidePointsForY, middleIndex+1, rightSideIndex)
	distanceArray = append(distanceArray, distanceStruct)

	if distanceArray[0].distance < distanceArray[1].distance { //finding minimal distance between two side
		minimalDistance = distanceArray[0].distance
		copy(closestPoints, distanceArray[0].coordinatesOfPoints)
	} else {
		minimalDistance = distanceArray[1].distance
		copy(closestPoints, distanceArray[1].coordinatesOfPoints)
	}

	for i := 0; i < len(pointsSortedByY); i++ { // adding points which are near center points and which are located at a distance less than the minimum distance
		if math.Abs(pointsSortedByY[i].x-pointsSortedByX[middleIndex].x) <= minimalDistance {
			pointsMemory = append(pointsMemory, pointsSortedByY[i])
		}
	}

	distanceFromMiddle := math.Inf(+1) // value to compare with
	closestPointsFromMiddle := []Point{}

	for i := 0; i < len(pointsMemory); i++ { // we sort through all the points that we sorted relative to the center by the y and look for the minimum of them
		for j := i + 1; j < len(pointsMemory); j++ {
			if (pointsMemory[j].y - pointsMemory[i].y) >= minimalDistance {
				break
			}

			currentDistance = getDistance(pointsMemory[i], pointsMemory[j])

			if currentDistance < distanceFromMiddle {
				distanceFromMiddle = currentDistance
				closestPointsFromMiddle = append(closestPoints, pointsMemory[i])
				closestPointsFromMiddle = append(closestPoints, pointsMemory[j])
			}
		}
	}

	if minimalDistance > distanceFromMiddle { // if we find a point not on the left and right side, but from middle point to this point
		return distanceFromMiddle, closestPointsFromMiddle
	}

	return minimalDistance, closestPoints
}

func sortBothArr(points []Point) ([]Point, []Point) {
	fmt.Println("Entered snail coordinates: ", points)
	sort.Slice(points, func(i, j int) bool { // sorting by x
		return points[i].x < points[j].x
	})

	pointsSortedByX := make([]Point, len(points))
	copy(pointsSortedByX, points)

	fmt.Println("Coordinates after sorting by х: ", pointsSortedByX)

	sort.Slice(points, func(i, j int) bool { // sorting by y
		return points[i].y < points[j].y
	})

	pointsSortedByY := make([]Point, len(points))
	copy(pointsSortedByY, points)

	fmt.Println("Coordinates after sorting by y: ", pointsSortedByY)

	return pointsSortedByX, pointsSortedByY
}

func main() {
	numberOfSnails := 0
	fmt.Print("Enter the number of snails(2 or more): ")
	fmt.Fscan(os.Stdin, &numberOfSnails)

	if numberOfSnails < 2 {
		fmt.Println("Please rerun program and enter more than 1 snail")
		return
	}
	var inputCoordinates []Point

	for i := 0; i < numberOfSnails; i++ {
		snailCoordinates := Point{}
		fmt.Print("Enter the coordinates of snail separated by a space(in meters):  ")
		fmt.Scanln(&snailCoordinates.x, &snailCoordinates.y)
		inputCoordinates = append(inputCoordinates, snailCoordinates)
	}

	inputSortedByX, inputSortedByY := sortBothArr(inputCoordinates)

	minimalDistance, coordinates := findClosestPoints(inputSortedByX, inputSortedByY, 0, len(inputCoordinates)-1)
	minimalTime := minimalDistance / 2 //divide by two because the snails are crawling towards and their speed is 1 sm /s ^ 2

	fmt.Printf("\nResults. \nThe first pair of snails will reach each other during %.2f seconds.\nThey will go the same way %.2f meters. \nCoordinates.\nFirst snail: X: %.1f ,Y: %.1f.\nSecond snail: X: %.1f ,Y: %.1f", minimalTime, minimalDistance, coordinates[0].x, coordinates[0].y, coordinates[1].x, coordinates[1].y)
	// in previous commit i tried to do another way by using goroutine, but different of ellapsed time win
	//For this example ellapsed time less than with using goroutine
}
