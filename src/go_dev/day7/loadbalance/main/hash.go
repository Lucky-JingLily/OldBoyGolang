package main

import (
	"fmt"
	"go_dev/day7/loadbalance/balance"
	"hash/crc32"
	"math/rand"
)

type HashBalance struct {
}

func init() {
	balance.RegisterBalancer("hash", &HashBalance{})
}

func (p *HashBalance) DoBalance(insts []*balance.Instance, key ...string) (inst *balance.Instance, err error) {
	var defKey string = fmt.Sprintf("%d", rand.Int())
	if len(key) > 0 {
		defKey = key[0]
	}

	lens := len(insts)
	if lens == 0 {
		err = fmt.Errorf("No instance")
		return
	}

	hashValue := crc32.Checksum([]byte(defKey), crc32.MakeTable(crc32.IEEE))
	index := int(hashValue) % lens
	inst = insts[index]
	return
}
