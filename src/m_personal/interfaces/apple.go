package main

type Apple struct {
	color    string
	tetoPear bool
	taste    string
}

func (a Apple) getColor() string {
	return a.color
}

func (a Apple) isTetoPear() bool {
	return a.tetoPear
}

func (a Apple) getTaste() string {
	return a.taste
}
