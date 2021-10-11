package main

func countNeighbour(p golParams, world [][]byte, x int, y int) int {
	neighbours := 0
	for i:=-1; i<2; i++{
		for j:=-1; j<2; j++{
			col := (x + i + p.imageWidth)%p.imageWidth
			row := (y + j + p.imageHeight)%p.imageHeight

			if (i!= 0 || j!= 0) && world[col][row] == 255 {neighbours++}
		}
	}
	return neighbours
}

func nextWorld(p golParams, world [][]byte) [][]byte{
	var nextWorld [][]byte
	for y:=0; y<p.imageWidth; y++{
		nextWorld = append(nextWorld, []byte{})
		for x:=0; x<p.imageHeight; x++{
			nextWorld[y] = append(nextWorld[y], world[y][x])
		}
	}
	return nextWorld
}

func calculateNextState(p golParams, world [][]byte) [][]byte {
	cols := p.imageWidth
	rows := p.imageHeight
	next := nextWorld(p, world)

	for i:=0; i<cols; i++{
		for j:=0; j<rows; j++{
			state := world[i][j]
			livingNeighbours := countNeighbour(p,world, i, j)
			//logic is not working
			if state == 0 && livingNeighbours == 3{
				next[i][j] = 255
			}else if state == 255{
				if livingNeighbours < 2 || livingNeighbours > 3{ next[i][j] = 0}
			}else {
				next[i][j] = state
			}
		}
	}
	return next
}



func calculateAliveCells(p golParams, world [][]byte) []cell {
	var ALiveCells []cell
	for col:=0; col < p.imageWidth; col++ {
		for row:=0; row < p.imageHeight; row++{
			if world[row][col] == 255 {
				ALiveCells = append(ALiveCells, cell{col,row})
			}
		}
	}
	return ALiveCells
}

