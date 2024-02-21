package game

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	SPADE    = iota // 黑桃
	HEART           // 红桃
	DIAMOND         // 方块
	CLUB            // 梅花
	JOKER           // 假牌
	MAX_SUIT        // 最大花色数量
)
const (
	HIDE = iota // 隐藏
	SHOW        // 显示
)

var back_img_black *ebiten.Image
var back_img_red *ebiten.Image

type Poker struct {
	id     int
	suit   byte
	value  byte
	img    *ebiten.Image
	x      float64
	y      float64
	rotate float64
	width  int
	height int
}

func InitBack(imgpath string, is_red bool) {
	img, _, err := ebitenutil.NewImageFromFile(imgpath, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	if is_red {
		back_img_red = img
		return
	}
	back_img_black = img
}
func NewPoker(id int, suit byte, value byte, imgpath string) *Poker {
	img, _, err := ebitenutil.NewImageFromFile(imgpath, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	p := &Poker{
		id:    id,
		suit:  suit,
		value: value,
	}
	p.img = img
	p.width, p.height = img.Size()
	return p
}

func (p *Poker) In(x, y float64) bool {
	if x < p.x || x > p.x+float64(p.width) {
		return false
	}

	if y < p.y || y > p.y+float64(p.height) {
		return false
	}

	return true
}

func (p *Poker) MoveX_offset(dx float64) {
	p.x += dx
}
func (p *Poker) MoveY_offset(dy float64) {
	p.y += dy
}
func (p *Poker) MoveXY(x, y float64) {
	p.x = x
	p.y = y
}
func (p *Poker) Rotate_offset(angle float64) {
	p.rotate += angle
}
func (p *Poker) Rotate(angle float64) {
	p.rotate = angle
}
func (p *Poker) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.x, p.y)
	op.GeoM.Rotate(p.rotate)
	screen.DrawImage(p.img, op)
}
