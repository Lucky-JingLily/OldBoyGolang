package balance

import (
	"errors"
	"math/rand"
)

type RandomBalance struct {
}

func init() {
	RegisterBalancer("random", &RandomBalance{})
}

func (p *RandomBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("No instance")
		return
	}

	lens := len(insts)
	index := rand.Intn(lens)
	inst = insts[index]

	return
}
