package game

type Table struct { 
	name   string
	pokers []*Poker
}

func NewTable(name string) *Table {
	return &Table{name, []*Poker{}}
}

func (t *Table) AddPoker(p *Poker) {
    
}
