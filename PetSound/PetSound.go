package PetSound

import (
	Cat "github.com/amanvatsdev/EasyAccess/PetSound/CatSound"
	Dog "github.com/amanvatsdev/EasyAccess/PetSound/DogBark"
)

func CheckPet(X string) {
	switch X {
	case "Cat":
		Cat.CatSound()
	case "Dog":
		Dog.DogSound()
	}
}
