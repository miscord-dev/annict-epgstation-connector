package epgstation

// this is pretty ugly but i dont have any other better options.

func NewTruePointer() *bool {
	v := true
	return &v
}

func NewFalsePointer() *bool {
	v := false
	return &v
}
