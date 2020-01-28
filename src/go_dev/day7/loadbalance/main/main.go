package main

import (
	"fmt"
	"go_dev/day7/loadbalance/balance"
	"math/rand"
	"os"
	"time"
)

func main() {
	var insts []*balance.Instance
	for i := 0; i < 10; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		port := 10000 + rand.Intn(10000)
		one := balance.NewInstance(host, port)
		insts = append(insts, one)
	}

	//var balancer balance.Balancer
	var balancerName = "random"
	if len(os.Args) > 1 {
		balancerName = os.Args[1]
	}
	//if balancerName == "random" {
	//	fmt.Println("use random")
	//	balancer = &balance.RandomBalance{}
	//} else if balancerName == "roundrobin" {
	//	fmt.Println("use roundrobin")
	//	balancer = &balance.Roundrobin{}
	//} else {
	//	fmt.Println("use random")
	//	balancer = &balance.RandomBalance{}
	//}

	// 随机
	//var balancer balance.RandomBalance
	//for {
	//	inst, err := balancer.DoBalance(insts)
	//	if err != nil {
	//		fmt.Println("do balance err: ", err)
	//		continue
	//	}
	//	fmt.Println(inst)
	//	time.Sleep(time.Second)
	//}

	// 轮训
	//var balancer balance.Roundrobin
	//for {
	//	inst, err := balancer.DoBalance(insts)
	//	if err != nil {
	//		fmt.Println("do balance err: ", err)
	//		continue
	//	}
	//	fmt.Println(inst)
	//	time.Sleep(time.Second)
	//}

	for {
		inst, err := balance.DoBalance(balancerName, insts)
		if err != nil {
			fmt.Println("do balance err: ", err)
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}
}
