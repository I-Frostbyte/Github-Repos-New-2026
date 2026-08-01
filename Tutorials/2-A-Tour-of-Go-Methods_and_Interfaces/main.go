package main

import (
	"fmt"

	methods "github.com/I-Frostbyte/Github-Repos-New-2026/Tutorials/2-A-Tour-of-Go-Methods_and_Interfaces/1-Methods"
)

func main() {
	v := methods.Vertex{X: 3, Y: 4}
	fmt.Println(v.Abs())
	newUser := methods.NewUser(
		"Alvin",
		"al@gmail.com",
		"+237695456923",
		"Bokwango",
		23,
	)
	fmt.Println(newUser.GetName(newUser))
}