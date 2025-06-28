package main

type Pear struct {
	color    string
	tetoPear bool
	taste    string
}

func (p Pear) getColor() string {
	return p.color
}

func (p Pear) isTetoPear() bool {
	return p.tetoPear
}

func (p Pear) getTaste() string {
	return p.taste
}
