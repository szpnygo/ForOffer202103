package main

import "fmt"

type user struct {
	name string
}

func main() {
	slice := []int{0, 1, 2, 3, 4}
	fmt.Println(slice)
	changeSlice(slice)
	fmt.Println(slice)
	newSlice := slice[1:3]
	fmt.Println(newSlice)
	newSlice[1] = 5
	fmt.Println(newSlice)
	fmt.Println(slice)

	newSlice = append(newSlice, 7)
	fmt.Println(newSlice)
	fmt.Println(slice)

	newSlice2 := slice[3:4:4]
	fmt.Println(newSlice2)
	newSlice2 = append(newSlice2, 8)
	fmt.Println(newSlice2)
	fmt.Println(slice)

	neo := user{name: "neo"}
	fmt.Println(neo)
	newNeo := changeUsername(neo)
	fmt.Println(neo)
	fmt.Println(newNeo)
	newNeo.name = "test2"
	fmt.Println(neo)
	fmt.Println(newNeo)
}

func changeSlice(slice []int) {
	slice[1] = 0
}

func changeUsername(u user) user {
	u.name = "test"
	return u
}
