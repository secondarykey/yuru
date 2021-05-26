package window

import (
	"fmt"
	"time"

	"github.com/secondarykey/yuru/logic"
)

type Play struct {
	start time.Time

	end   string
	route logic.Route
	drop  *Drop
}

func NewPlay(start *Drop) *Play {
	var p Play
	p.route = make([]int, 0)
	p.drop = start
	p.end = ""
	return &p
}

func (p *Play) isStarted() bool {
	if p == nil {
		return false
	}
	if p.start.IsZero() {
		return false
	}
	return true
}

func (p *Play) isStoped() bool {
	if p == nil {
		return false
	}
	if p.end == "" {
		return false
	}
	return true
}

func (p *Play) startPlay() {
	p.start = time.Now()
	p.end = ""
	p.route = make([]int, 0)
}

func (p *Play) stopPlay() {
	if !p.isStarted() {
		return
	}

	p.end = p.duration()
	p.start = time.Time{}
}

func (p *Play) move(d *Drop) {
	if !p.isStarted() {
		p.startPlay()
	}
	i := p.getDirection(d)
	p.route = append(p.route, i)
}

func (p *Play) getDirection(d *Drop) int {

	sp1 := p.drop.Sprite
	sp2 := d.Sprite

	x := sp2.x - sp1.x
	y := sp2.y - sp1.y

	dx := 0
	if x > 0 {
		dx = 1
	} else if x < 0 {
		dx = -1
	}

	dy := 0
	if y > 0 {
		dy = 1
	} else if y < 0 {
		dy = -1
	}

	for i := 0; i < logic.N; i++ {
		if logic.DC[i] == dx && logic.DR[i] == dy {
			return i
		}
	}
	return -1
}

func (p *Play) duration() string {
	d := p.end
	if d == "" {
		now := time.Now()
		sub := now.Sub(p.start)
		d = fmt.Sprintf("%0.2f", sub.Seconds())
	}
	return d
}

func (p *Play) routeString() string {
	return p.route.DirectionString()
}
