package main

import (
	"log"

	"example.com/qianqiuLiang/gameStudy/game"
	"github.com/hajimehoshi/ebiten"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello World")
	g := game.NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
