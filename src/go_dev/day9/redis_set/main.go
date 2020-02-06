package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	var p *int
	var a int
	p = &a
	*p = 0

	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("connect redis failed,", err)
		return
	}
	defer c.Close()

	_, err = c.Do("MSet", "abc", 123, "efg", 321)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Ints(c.Do("MGet", "abc", "efg"))
	if err != nil {
		fmt.Println("get adb failed,", err)
		return
	}
	fmt.Println(r)

	_, err = c.Do("expire", "abc", 10)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = c.Do("lpush", "book_list", "abc", "ceq", 300)
	if err != nil {
		fmt.Println(err)
		return
	}

	s, err := redis.String(c.Do("lpop", "book_list"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}

	fmt.Println(s)

}
