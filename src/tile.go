package src

type Tile struct {
	BlockID     uint16
	BlockActive bool
	// Diğer tile özellikleri eklenebilir (ör. wall, liquid, vb.)
}

type TileMatrix struct {
	Tiles [][]Tile
}

func NewTileMatrix(width, height int) *TileMatrix {
	tiles := make([][]Tile, width)
	for i := range tiles {
		tiles[i] = make([]Tile, height)
	} 
	return &TileMatrix{Tiles: tiles}
}
