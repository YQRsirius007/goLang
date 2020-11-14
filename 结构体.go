package pos

import "fmt"

type test struct {
	rate     uint
	interval uint
}

func tt() {
	t := &test{
		rate: 10,
	}
	fmt.Println(t)
}
