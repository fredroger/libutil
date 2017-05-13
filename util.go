package libutil

import (
	"errors"
	"fmt"
)

// MyFirstFunction is my first function
func MyFirstFunction(string) error {

	fmt.Println("MyFirstFunction was called")

	return errors.New("My first error")
}

// MySecondFunction is my second function
func MySecondFunction(condition int) int {

	if condition == 0 {
		return 100
	}

	return -100
}
