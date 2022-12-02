package main

type Podium struct {
	first  int
	second int
	third  int
}

func (p *Podium) sum() int {
	return p.first + p.second + p.third
}

func (p *Podium) handle(v int) {
	if v > p.first {
		old := p.first
		p.first = v
		p.handle(old)
	} else if v > p.second {
		old := p.second
		p.second = v
		p.handle(old)
	} else if v > p.third {
		old := p.third
		p.third = v
		p.handle(old)
	}
}
