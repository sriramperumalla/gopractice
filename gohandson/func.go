package gohandson

import (
	"fmt"
	//"os"
	//"io"
)

// EvenNumGenerator function calls an anonymous function  func() which returns a func() and an unsigned int
func EvenNumGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) { // closure anonymous function accessing variable i defined outside its body and returns uint
		ret = i
		i += 2
		return
	}
}
func generateEven() {
	nextEven := EvenNumGenerator() // declared nextEven variable and initialized with the function as its value
	fmt.Println(nextEven())        // 0
	fmt.Println(nextEven())        // 2
	fmt.Println(nextEven())        // 4
}
