package pos

import (
	"testing"
)

func TestShouldReturnReceipt2(t *testing.T) {
	currentPos := []Position{
		{x: 1, y: 2, facing: 'W'},
		{x: 1, y: 2, facing: 'W'},
		{x: 1, y: 2, facing: 'W'},
		{x: 1, y: 2, facing: 'W'},
		{x: 1, y: 2, facing: 'W'},
	}
	newPos := []Position{
		{x: 0, y: 2, facing: 'W'},
		{x: 1, y: 2, facing: 'S'},
		{x: 1, y: 2, facing: 'N'},
		{x: 1, y: 1, facing: 'W'},
		{x: 1, y: 3, facing: 'W'},
	}
	command := []string{
		"M",
		"L",
		"R",
		"LMR",
		"RML",
	}
	for i := range currentPos {
		c := MarsRover(currentPos[i], command[i])
		if !(c.x == newPos[i].x && c.y == newPos[i].y && c.facing == newPos[i].facing) {
			t.Errorf("want %+v, but got %+v", newPos[i], c)
		}
		// c.x == newPos[i].x
		// c.y == newPos[i].y
		// c.byte == newPos[i].byte
		// marsEqual(c, newPos[i])
	}
}
