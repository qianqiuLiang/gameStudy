package game

const (
	STACK_DEALED = iota //已经发出
	STACK_HAND
	STACK_SUCCESS
)

type PokerStack struct {
	pokers   []*Poker
	id2Index map[int]int
	status   uint8

	x, y          float64
	width, height float64
}

func NewPokerStack(kind uint8) *PokerStack {
	return &PokerStack{status: kind}
}
func (this *PokerStack) AddPoker(poker *Poker) {
	this.pokers = append(this.pokers, poker)
	this.id2Index[poker.id] = len(this.pokers) - 1
}
func (this *PokerStack) GetPoker() *Poker {
	return this.pokers[len(this.pokers)-1]
}
func (this *PokerStack) ClearPokers() {
	this.pokers = this.pokers[:0]
}
func (this *PokerStack) InStack(x, y float64) (bool, *Poker) {
	if x >= this.x && x <= this.x+this.width &&
		y >= this.y && y <= this.y+this.height {
		return this.inPoker(x, y)
	}
	return false, nil
}
func (this *PokerStack) inPoker(x, y float64) (bool, *Poker) {
	for _, poker := range this.pokers {
		if poker.In(x, y) {
			return true, poker
		}
	}
	return false, nil
}
func (this *PokerStack) MoveOutLastPoker() (bool, *Poker) {
	if len(this.pokers) == 0 {
		return false, nil
	}
	poker := this.pokers[len(this.pokers)-1]
	return true, poker
}

// 连续才可移动
func (this *PokerStack) CanMove(poker *Poker) bool {
	if this.status != STACK_DEALED {
		return true
	}
	index, ok := this.id2Index[poker.id]
	if !ok {
		return false
	}
	tmpValue := poker.value + 1
	tmpSuit := poker.suit
	for i := index + 1; i < len(this.pokers); i++ {
		if this.pokers[i].value != tmpValue {
			return false
		}
		if this.pokers[i].suit != tmpSuit {
			return false
		}
		tmpValue++
	}
	return true
}
func (this *PokerStack) MoveInPokers(poker *Poker) bool {
	if this.status != STACK_DEALED {
		this.AddPoker(poker)
		return true
	}

}
