package main

import (
	"fmt"
	"lec1/animal"
	cat2 "lec1/cat"
	"lec1/duck"
)

func main() {
	cat := cat2.NewCat("meow")
	cat.SetSound("meow")
	dck := duck.Duck{Sound: "qrya"}
	sound1 := MakeASound(&cat)
	sound2 := MakeASound(&dck)

	fmt.Println("Sound for cat " + sound1)
	fmt.Println("Sound for duck " + sound2)
	dck.SecretFunction()
}

func MakeASound(a animal.Animal) string {
	return a.MakeASound()
}
