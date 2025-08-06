package src

type Chest struct {
	X        int32
	Y        int32
	Name     string
	Contents []ItemStack // Her slot için bir ItemStack veya nil
}

type ItemStack struct {
	Quantity int16
	TypeID   int32
	Prefix   uint8
}
