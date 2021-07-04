package main

import (
	"fmt"
)

func main() {


	matrix := []int{
		1,1,1,1,1,1,1,1,1,1,1,
		1,0,0,0,0,0,0,0,0,0,1,
		1,0,0,0,1,0,0,0,0,0,1,
		1,0,0,0,1,0,0,0,0,0,1,
		1,0,0,0,1,0,0,0,0,0,1,
		1,0,0,0,1,1,1,0,0,0,1,
		1,0,0,0,0,0,1,0,0,0,1,
		1,0,0,0,0,0,0,0,0,0,1,
		1,1,1,1,1,1,1,1,1,1,1,
	}

	rowCount := 9
	colCount := 11

	startPointX := 1
	startPointY := 1

	finishPointX := 9
	finishPointY := 7


	var useDefault int
	fmt.Println("Use default? (1 - yes, 0 - no)")
	fmt.Scanf("%d", &useDefault)

	if useDefault == 0 {

		fmt.Println("Input row count: ")
		fmt.Scanf("%d", &rowCount)

		fmt.Println("Input column count: ")
		fmt.Scanf("%d", &colCount)


		matrix = make([]int,rowCount*colCount)
		for i := 0; i < rowCount; i++ {
			for j := 0; j < colCount; j++ {
				fmt.Printf("Input value for cell %d %d (0 - free, 1 - locked): ", j, i)
				fmt.Scanf("%d", &matrix[i*colCount+j])
				if matrix[i*colCount+j] < 0 || matrix[i*colCount+j] > 1 {
					panic("Wrong input")
				}
				println("")
			}
		}

		fmt.Println("Input start point x coord: ")
		fmt.Scanf("%d", &startPointX)
		fmt.Println("Input start point y coord: ")
		fmt.Scanf("%d", &startPointY)
		if startPointX<0 || startPointX>colCount-1 || startPointY<0 || startPointY>rowCount-1 || matrix[startPointY*colCount+startPointX]!=0 {
			panic("Wrong start point!")
		}


		fmt.Println("Input finish point x coord: ")
		fmt.Scanf("%d", &finishPointX)
		fmt.Println("Input finish point y coord: ")
		fmt.Scanf("%d", &finishPointY)
		if finishPointX<0 || finishPointX>colCount-1 || finishPointY<0 || finishPointY>rowCount-1 || matrix[finishPointY*colCount+finishPointX]!=0 {
			panic("Wrong finish point!")
		}
	}


	points := make([]Point,0)

	for i := 0; i < rowCount; i++ {
		for j := 0; j < colCount; j++ {
			point:=ConvertCoordsToPoint(j,i)
			switch matrix[i*colCount+j] {
				case 1:point.State = locked
				case 0:point.State = free
			}
			points = append(points, point)
		}

	}

	astar := AStar{}
	astar.matrix = points
	astar.rowCount = rowCount
	astar.colCount = colCount

	for _, point := range astar.findPath(points[startPointY*colCount+startPointX],points[finishPointY*colCount+finishPointX]) {
		println(point.X, point.Y)
	}
}


