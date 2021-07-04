package main

type AStar struct {
	matrix []Point
	rowCount int
	colCount int
}

type PointState int
const (
	free = iota
	checking
	locked
)

type Point struct {
	X int
	Y int
	State PointState
	Parent *Point

	priceToStart int
	priceToFinish int
	priceTotal int

}

func (this *AStar) findPath(start, finish Point) []Point {
	mapChecking := make([]Point,0)
	start.State = checking
	mapChecking = append(mapChecking, start)

	for ok := true; ok; ok = len(mapChecking)>0 {

		for i := 0; i < len(mapChecking); i++ {
			mapChecking = this.calcPrices(mapChecking, i, finish)
			if this.pathIsFound(mapChecking, i) {
				path := this.buildPath(mapChecking, i, finish)
				return reverse(path)
			}
		}

		totalMin := this.GetMinTotal(mapChecking)
		mapChecking = this.GenerateNewMapChecking(mapChecking, totalMin)
	}
	return make([]Point,0)
}

func reverse(arr []Point) []Point {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func (this *AStar) buildPath(mapChecking []Point, index int, finish Point) []Point {
	currentCheckingPoint := mapChecking[index]
	path := make([]Point, 0)
	if this.matrix[currentCheckingPoint.Y*this.colCount+currentCheckingPoint.X].priceToStart > 0 {
		path = append(path, currentCheckingPoint)
		for currentCheckingPoint.Parent!=nil {
			cp := currentCheckingPoint.Parent
			path = append(path, *cp)
			currentCheckingPoint = *cp
		}
		return path
 	} else {
		path = append(path, finish)
	}
	return path
}

func (this *AStar) GenerateNewMapChecking(mapChecking []Point, totalMin int) []Point {

	for i := len(mapChecking)-1; i >=0; i-- {
		currentCheckingPoint := mapChecking[i]

		if this.matrix[currentCheckingPoint.Y*this.colCount+currentCheckingPoint.X].priceTotal == totalMin {
			if this.matrix[currentCheckingPoint.Y*this.colCount+(currentCheckingPoint.X-1)].State == free {
				point:=ConvertCoordsToPoint(currentCheckingPoint.X-1, currentCheckingPoint.Y)
				point.Parent = &currentCheckingPoint
				this.matrix[(currentCheckingPoint.Y)*this.colCount+(currentCheckingPoint.X-1)].Parent = &currentCheckingPoint
				mapChecking = append(mapChecking, point)
			}
			if this.matrix[currentCheckingPoint.Y*this.colCount+(currentCheckingPoint.X+1)].State == free {
				point:=ConvertCoordsToPoint(currentCheckingPoint.X+1, currentCheckingPoint.Y)
				point.Parent = &currentCheckingPoint
				this.matrix[(currentCheckingPoint.Y)*this.colCount+(currentCheckingPoint.X+1)].Parent = &currentCheckingPoint
				mapChecking = append(mapChecking, point)
			}
			if this.matrix[(currentCheckingPoint.Y-1)*this.colCount+(currentCheckingPoint.X)].State == free {
				point:=ConvertCoordsToPoint(currentCheckingPoint.X, currentCheckingPoint.Y-1)
				point.Parent = &currentCheckingPoint
				this.matrix[(currentCheckingPoint.Y-1)*this.colCount+(currentCheckingPoint.X)].Parent = &currentCheckingPoint
				mapChecking = append(mapChecking, point)
			}
			if this.matrix[(currentCheckingPoint.Y+1)*this.colCount+(currentCheckingPoint.X)].State == free {
				point:=ConvertCoordsToPoint(currentCheckingPoint.X, currentCheckingPoint.Y+1)
				point.Parent = &currentCheckingPoint
				this.matrix[(currentCheckingPoint.Y+1)*this.colCount+(currentCheckingPoint.X)].Parent = &currentCheckingPoint
				mapChecking = append(mapChecking, point)
			}
			this.matrix[currentCheckingPoint.Y*this.colCount+currentCheckingPoint.X].State = locked
			mapChecking = append(mapChecking[:i], mapChecking[i+1:]...)
		}
	}

	return mapChecking

}

func (this *AStar) GetMinTotal(mapChecking []Point) int {
	startX:=mapChecking[0].X
	startY:=mapChecking[0].Y
	totalMin:=this.matrix[startY*this.colCount+startX].priceTotal

	for i := 1; i < len(mapChecking); i++ {
		currentX := mapChecking[i].X
		currentY := mapChecking[i].Y
		currentPriceTotal := this.matrix[currentY*this.colCount+currentX].priceTotal
		if totalMin>currentPriceTotal {
			totalMin=currentPriceTotal
		}
	}
	return totalMin
}

func (this *AStar) pathIsFound(mapChecking []Point, index int) bool {
	currentCheckingPoint := mapChecking[index]
	return this.matrix[currentCheckingPoint.Y*this.colCount+currentCheckingPoint.X].priceToFinish==0
}

func (this *AStar) calcPrices(mapChecking []Point, index int, finish Point) []Point {
	currentCheckingPoint := mapChecking[index]

	if parent:=this.matrix[currentCheckingPoint.Y*this.colCount+currentCheckingPoint.X].Parent; parent != nil {
		this.matrix[currentCheckingPoint.Y*this.colCount+currentCheckingPoint.X].priceToStart = parent.priceToStart+1
	} else {
		this.matrix[currentCheckingPoint.Y*this.colCount+currentCheckingPoint.X].priceToStart = 0
	}
	this.matrix[currentCheckingPoint.Y*this.colCount+currentCheckingPoint.X].priceToFinish = Abs(currentCheckingPoint.X-finish.X)+Abs(currentCheckingPoint.Y-finish.Y)
	this.matrix[currentCheckingPoint.Y*this.colCount+currentCheckingPoint.X].priceTotal = this.matrix[currentCheckingPoint.Y*this.colCount+currentCheckingPoint.X].priceToStart + this.matrix[currentCheckingPoint.Y*this.colCount+currentCheckingPoint.X].priceToFinish
	mapChecking[index] = this.matrix[currentCheckingPoint.Y*this.colCount+currentCheckingPoint.X]
	return mapChecking
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ConvertCoordsToPoint(x,y int) Point {
	point := Point{}
	point.X = x
	point.Y = y
	point.priceTotal = -1
	point.priceToStart= -1
	point.priceToFinish = -1
	point.State = checking
	return point
}
