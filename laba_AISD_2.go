package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
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

	var wg sync.WaitGroup                  // this is wait-group for goroutine(thread in basic).it helps to keep track of all the threads
	distanceChannel := make(chan Distance) // this is a channel, thanks to which we transfer data from gorutin to main

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

	wg.Add(1)                                      // adding one goroutine(thread) to track it
	go func(wg *sync.WaitGroup, c chan Distance) { // start thread
		defer wg.Done()                                                                                                               // when func will finish, we will track it
		distanceLeftSide, closestPointsLeftSide := findClosestPoints(pointsSortedByX, leftSidePointsForY, leftSideIndex, middleIndex) // get closest of left side
		c <- Distance{distance: distanceLeftSide, coordinatesOfPoints: closestPointsLeftSide}                                         // write it to channel for contain this data
	}(&wg, distanceChannel)

	wg.Add(1)
	go func(wg *sync.WaitGroup, c chan Distance) {
		defer wg.Done()                                                                                                                     // when func will finish, we will track it
		distanceRightSide, closestPointsRightSide := findClosestPoints(pointsSortedByX, rightSidePointsForY, middleIndex+1, rightSideIndex) // get closest of right side
		c <- Distance{distance: distanceRightSide, coordinatesOfPoints: closestPointsRightSide}                                             // write it to channel for contain this data
	}(&wg, distanceChannel)

	distanceArray := []Distance{} // array for contain all info about getting from goroutines(threads)

	for len(distanceArray) < 2 {
		distanceArray = append(distanceArray, <-distanceChannel) // read from channel data and put it into array
	}

	wg.Wait()              // we are waiting for closing all threads
	close(distanceChannel) // we close a channel

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
	fmt.Println("До сортировки: ", points)
	sort.Slice(points, func(i, j int) bool { // sorting by x
		return points[i].x < points[j].x
	})

	pointsSortedByX := make([]Point, len(points))
	copy(pointsSortedByX, points)

	fmt.Println("После сортировки по х: ", pointsSortedByX)

	sort.Slice(points, func(i, j int) bool { // sorting by y
		return points[i].y < points[j].y
	})

	pointsSortedByY := make([]Point, len(points))
	copy(pointsSortedByY, points)

	fmt.Println("После сортировки по y: ", pointsSortedByY)

	return pointsSortedByX, pointsSortedByY
}

func main() {
	//input := []Point{{1, 1}, {1, 2}, {2, 1}, {2, 2}, {10, 10}}
	input := []Point{{4, 10}, {3, 7}, {9, 7}, {3, 4}, {5, 6}, {5, 4}, {6, 3}, {8, 1}, {3, 0}, {1, 6}}
	inputSortedByX, inputSortedByY := sortBothArr(input)

	fmt.Println(findClosestPoints(inputSortedByX, inputSortedByY, 0, len(input)-1))

}
