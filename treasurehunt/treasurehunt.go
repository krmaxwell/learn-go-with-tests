package treasurehunt

import "fmt"

type TreasureMap [5][5]int
type MapLocation struct {
	X int
	Y int
}

func (t TreasureMap) findNextTreasureLocation(location MapLocation) MapLocation {
	locX := location.X
	locY := location.Y
	nextX := t[locX-1][locY-1] / 10
	nextY := t[locX-1][locY-1] % 10

	loc := MapLocation{nextX, nextY}
	return loc
}

func (t TreasureMap) FindTreasure() MapLocation {
	currentLocation := MapLocation{1, 1}
	for {
		fmt.Printf("Currently at X %d Y %d\n", currentLocation.X, currentLocation.Y)
		currentLocation = t.findNextTreasureLocation(currentLocation)
		if currentLocation.X == currentLocation.Y {
			break
		}
	}
	fmt.Printf("Treasure is at X %d Y %d\n", currentLocation.X, currentLocation.Y)
	return currentLocation
}
