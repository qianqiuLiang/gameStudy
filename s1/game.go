package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Game struct {
	i *Input
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.i.Update()
	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x00, 0x44, 0x44, 0xff}) // 设置背景颜色
	ebitenutil.DebugPrint(screen, g.i.msg)
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 240 // 设置游戏窗口的尺寸
}
func NewGame() *Game {
	return &Game{i: &Input{msg: "hello"}}
}
