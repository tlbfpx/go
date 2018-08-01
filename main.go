package main

import (
	"fmt"
	"gotest/user"
)

func main() {

	users := []User{
		{"zhang san", 12},
		{"li si", 30},
		{"wang wu", 52},
		{"zhao liu", 26},
	}

	name := "jack"
	fmt.Printf("hello,%s\n", name)

	user.Sort(users)

}
