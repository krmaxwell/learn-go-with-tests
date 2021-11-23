package treasurehunt

import (
	"testing"
)

func TestTreasureHunt(t *testing.T) {

	myTreasureMap := TreasureMap{
		{34, 21, 32, 41, 25},
		{14, 42, 43, 14, 31},
		{54, 45, 52, 42, 23},
		{33, 15, 51, 31, 35},
		{21, 52, 33, 13, 23},
	}

	t.Run("Find first location", func(t *testing.T) {
		firstLocation := MapLocation{X: 1, Y: 1}
		got := myTreasureMap.findNextTreasureLocation(firstLocation)
		want := MapLocation{X: 3, Y: 4}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
