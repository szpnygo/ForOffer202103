package main

import "fmt"

type user struct {
	name  string
	email string
}

func (c user) print() {
	fmt.Println(c.name, "=", c.email)
}

func (c *user) changeName() {
	c.name = "*changeName"
}

func (c user) changeEmail() {
	c.email = "changeEmail"
}

func (c *user) changeEmail2() {
	c.email = "*changeEmail2"
}

func main() {
	neo := user{
		name:  "neo",
		email: "neo@neobaran.com",
	}
	neo.print()
	neo.changeName()
	neo.print()
	neo.changeEmail()
	neo.print()
	neo.changeEmail2()
	neo.print()

	su := &user{
		name:  "su",
		email: "su@neobaran.com",
	}
	su.print()
	su.changeName()
	su.print()
	(*su).changeEmail()
	su.print()
	su.changeEmail()
	su.print()
}
