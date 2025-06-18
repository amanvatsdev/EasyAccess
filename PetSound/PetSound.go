package PetSound

import (
	Cat "Compiler/PetSound/CatSound"
	Dog "Compiler/PetSound/DogBark"
)

func CheckPet(X string) {
	switch X {
	case "Cat":
		Cat.CatSound()
	case "Dog":
		Dog.DogSound()
	}
}
