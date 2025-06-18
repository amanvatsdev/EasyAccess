package Baby

import "fmt"

func CheckBaby(age int) {
	if age < 5 {
		fmt.Println("Yes its baby ")
	}else if age<=5 {
		fmt.Println("no its not baby its child")
	}else {
		fmt.Println("its neither baby nor child")
	}

}