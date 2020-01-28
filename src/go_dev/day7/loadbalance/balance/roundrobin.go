package balance

import (
	"errors"
)

type Roundrobin struct {
	curIndex int
}

func init() {
	RegisterBalancer("roundrobin", &Roundrobin{})
}

func (p *Roundrobin) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("No instance")
		return
	}

	lens := len(insts)
	if p.curIndex >= lens {
		p.curIndex = 0
	}

	inst = insts[p.curIndex]
	p.curIndex = (p.curIndex + 1) % lens
	return
}
