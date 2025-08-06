package src

type TileEntity struct {
	TypeID uint8
	ID     int32
	X      int32
	Y      int32
	// Ekstra alanlar: ör. ItemFrame, Mannequin, Pylon, vb. için interface veya struct eklenebilir
}
