package main

import (
	"fmt"

	"github.com/amanvatsdev/EasyAccess/PetSound"
	Coffeecorner "github.com/amanvatsdev/Practice/CoffeeCorner"
	Teacorner "github.com/amanvatsdev/Practice/TeaCorner"
)

type PetInfo struct {
	Category string
	Age      int
}

func main() {
	PetInfo1 := PetInfo{
		Category: "Dog",
		Age:      2,
	}
	PetInfo2 := PetInfo{
		Category: "Cat",
		Age:      3,
	}

	fmt.Println(PetInfo1)
	fmt.Println(PetInfo2)
	PetSound.CheckPet(PetInfo1.Category)
	PetSound.CheckPet(PetInfo2.Category)


	fmt.Println("-------------This is the corner for Pet Animals-------")
	Teacorner.PrintTea()
	Coffeecorner.PrintCoffee()

}
