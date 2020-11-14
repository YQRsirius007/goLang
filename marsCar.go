package pos

import "fmt"

type Position struct {
	x, y   int
	facing byte
}

func (pos *Position) move() {
	if pos.facing == 'E' {
		pos.x++
	} else if pos.facing == 'W' {
		pos.x--
	} else if pos.facing == 'N' {
		pos.y++
	} else if pos.facing == 'S' {
		pos.y--
	}
}
func (pos *Position) turnRight() {
	if pos.facing == 'E' {
		pos.facing = 'S'
	} else if pos.facing == 'W' {
		pos.facing = 'N'
	} else if pos.facing == 'N' {
		pos.facing = 'E'
	} else if pos.facing == 'S' {
		pos.facing = 'W'
	}
}
func (pos *Position) turnLeft() {
	if pos.facing == 'E' {
		pos.facing = 'N'
	} else if pos.facing == 'W' {
		pos.facing = 'S'
	} else if pos.facing == 'N' {
		pos.facing = 'W'
	} else if pos.facing == 'S' {
		pos.facing = 'E'
	}
}
func MarsRover(currentPos Position, command string) Position {
	newPos := currentPos
	if command == "" {
		return newPos
	}
	comArr := []byte(command)
	for _, v := range comArr {
		if v == 'M' {
			newPos.move()
		} else if v == 'L' {
			newPos.turnLeft()
		} else if v == 'R' {
			newPos.turnRight()
		}
		fmt.Println(v, newPos)
	}

	return newPos
}
