package main

import (
	"fmt"
)

func main() {
	//comb.Set(10,[]string{"a", "b"})

	a := GetFirst()

	for i := 0; i < 100; i++ {
		s, err := Next(a)

		if err != nil {
			fmt.Println(err)
		}

		a = s

		fmt.Println(s)
	}
}
