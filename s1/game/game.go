package game

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	GAME_LEVEL_EASY = iota
	GAME_LEVEL_MID
	GAME_LEVEL_HARD
)

type Game struct {
	pokers      []*Poker
	score       int
	level       int
	dealedStack PokerStack
	handleStack PokerStack
	sucessStack PokerStack

	i *Input
}

func (g *Game) Shuffer() {
	g.score = 0
	//g.
	id := 0
	if len(g.pokers) == 0 {
		for i := 1; i < 13; i++ {
			img_path := "imgs/club-" + strconv.Itoa(i) + ".png"
			g.pokers = append(g.pokers, NewPoker(id, CLUB, byte(i), img_path))
			id++
		}
		for i := 1; i < 13; i++ {
			img_path := "imgs/diamond-" + strconv.Itoa(i) + ".png"
			g.pokers = append(g.pokers, NewPoker(id, DIAMOND, byte(i), img_path))
			id++
		}
		for i := 1; i < 13; i++ {
			img_path := "imgs/heart-" + strconv.Itoa(i) + ".png"
			g.pokers = append(g.pokers, NewPoker(id, HEART, byte(i), img_path))
			id++
		}
		for i := 1; i < 13; i++ {
			img_path := "imgs/spade-" + strconv.Itoa(i) + ".png"
			g.pokers = append(g.pokers, NewPoker(id, SPADE, byte(i), img_path))
			id++
		}
		InitBack("imgs/back1.png", false)
		InitBack("imgs/back2.png", true)

		g.dealedStack = NewPokerStack()
		
	}
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.i.Update()
	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x00, 0x44, 0x44, 0xff}) // 设置背景颜色
	ebitenutil.DebugPrint(screen, g.i.msg)

	for _, p := range g.pokers {
		p.Draw(screen)
	}
}
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 240 // 设置游戏窗口的尺寸
}
func NewGame() *Game {
	return &Game{}
}
