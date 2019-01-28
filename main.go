package main

import (
	"fmt"

	"github.com/ellenkorbes/parent"
	"github.com/ellenkorbes/partner"
)

func main() {

	fmt.Println("Parent:")
	fmt.Println(
		parent.Water(),
	// parent.Coffee(),
	// parent.Bread()
	)

	fmt.Println("\nPartner:")
	fmt.Println(
		partner.Water(),
	// partner.Coffee(),
	// partner.Bread()
	)

}
