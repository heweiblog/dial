package main

import (
	"fmt"
	"time"
)

type User struct {
	name string
	age  int
}

func run(xxx *User, set map[*User]bool) {
	for {
		if set[xxx] {
			fmt.Println("set[xxx]=", set[xxx])
		} else {
			//break
			fmt.Println("set[xxx]=", set[xxx])
		}
		time.Sleep(time.Second)
	}
	fmt.Println("set[xxx]=", set[xxx])
}

func main() {
	set := make(map[*User]bool)
	fmt.Println(set)
	hww := &User{"hww", 18}
	set[hww] = true
	mnn := &User{"mnn", 16}
	set[mnn] = true
	xxx := &User{"xxx", 22}
	set[xxx] = true
	fmt.Println(set)
	go run(xxx, set)
	time.Sleep(time.Second)
	delete(set, xxx)
	time.Sleep(time.Second * 5)

	ma := make(map[string]interface{})
	au := ma["dsf"]
	fmt.Println("au=", au)
	if aa, ok := au.(User); ok {
		fmt.Println("ok", aa)
	} else {
		fmt.Println("no", aa)
	}
}
