package gohandson

import (
	"fmt"
)

func myFunc(fn func()) {
	fmt.Println("Hello from myFunc")
	fn()
}

// Anonymous function func() is called from myFunc
func Anonymous() {
	fmt.Println("--- Anonymous Function ---")
	myFunc(func() {
		fmt.Println("Hello from the anonymous function")
	})
}
