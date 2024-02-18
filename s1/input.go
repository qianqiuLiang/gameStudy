package main

import "github.com/hajimehoshi/ebiten"

type Input struct {
	msg string
}

func (i *Input) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		i.msg = "pressed left"
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		i.msg = "pressed right"
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
		i.msg = "pressed up"
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		i.msg = "pressed down"
	} else {
		i.msg = ""
	}
}
