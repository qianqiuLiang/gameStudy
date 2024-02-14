package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Input struct {
	msg string
}

func (i *Input) Update() {
	// 处理输入
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		i.msg = "pressed left key"
		fmt.Println("pressed left key")
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		i.msg = "pressed right key"
		fmt.Println("pressed right key")
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
		i.msg = "pressed up key"
		fmt.Println("pressed up key")
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		i.msg = "pressed down key"
		fmt.Println("pressed down key")
	}
}

type Game struct {
	i *Input
}

func (g *Game) Update(screen *ebiten.Image) error {
	//fmt.Println("游戏开始")
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

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello World")
	g := NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
