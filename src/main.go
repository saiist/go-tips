package src

import (
	"fmt"
)

func Main() {
	fmt.Println("Hello, World!")

	i := 0
	if true {
		i := 1

		fmt.Println(i)
	}

	fmt.Println(i)
}
