package main

import "fmt"

type notifier interface {
	notify()

	test()
}

type user struct {
}

func (u *user) notify() {
	fmt.Println("notify")
}

func (u *user) test() {
	fmt.Println("test")
}

func main() {
	neo := user{}
	test(&neo)
	notify(&neo)
}

func notify(n notifier) {
	n.notify()
}

func test(n notifier) {
	n.test()
}
