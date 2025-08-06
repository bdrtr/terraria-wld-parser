package main

type MazeCell struct {
	IsCarved    bool
	Connections [4]bool // [right, left, down, up]
}

type MazeData struct {
	Width    int
	Height   int
	Cells    [][]MazeCell
	CellSize int
	HallSize int
	WallSize int
	Margin   int
}

func NewMazeData(worldWidth, worldHeight, margin int) *MazeData {
	cellSize := 8
	hallSize := 5
	wallSize := 3
	width := (worldWidth - 2*margin) / cellSize
	height := (worldHeight - 2*margin) / cellSize
	cells := make([][]MazeCell, width)
	for i := range cells {
		cells[i] = make([]MazeCell, height)
	}
	return &MazeData{
		Width:    width,
		Height:   height,
		Cells:    cells,
		CellSize: cellSize,
		HallSize: hallSize,
		WallSize: wallSize,
		Margin:   margin,
	}
}

// Maze generation and apply logic buraya eklenecek
