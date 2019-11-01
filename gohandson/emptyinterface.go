package gohandson

import (
	"fmt"
)

// DisplayInterface - Empty interface can hold value of any type. The type implements atleast zero methods.Empty interfaces are used by code that handles values of unknown type
func DisplayInterface() {
	var i interface{}
	show(i)

	i = 10
	show(i)

	i = "Hello Interface"
	show(i)

}

func show(i interface{}) {
	fmt.Println("--- Empty Interface Implementation ---")
	fmt.Printf("Type of i is %T and value is %v\n", i, i)
	fmt.Println("--- Empty Interface Implementation ---")
}
