package treasurehunt

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
	return currentLocation
}
