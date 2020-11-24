package combination

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	//Set(10,[]string{"a", "b", "b", "v", "a"})
	Set(10, "asdfghj")

	a := GetFirst()

	for i := 0; i < 10; i++ {
		s, err := Next(a)

		if err != nil {
			fmt.Println(err)
		}

		a = s

		fmt.Println(s)
	}
}
