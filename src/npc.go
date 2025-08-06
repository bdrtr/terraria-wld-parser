package src

type NPC struct {
	TypeID     int32
	Name       string
	X          float32
	Y          float32
	IsHomeless bool
	HomeX      int32
	HomeY      int32
	Variation  int32
}
