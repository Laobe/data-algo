package main

import (
	"fmt"
)

var input []int

func main() {
	// input = read_int()
	test()
	// fmt.Printf("%v\n", input)
}

func read_string() []string {
	var tmp string
	var input []string
	for {
		_, err := fmt.Scanf("%s", &tmp)
		if err != nil {
			break
		}
		input = append(input, tmp)
	}
	return input
}

func read_int() []int {
	var tmp int
	var input []int
	for {
		_, err := fmt.Scanf("%d", &tmp)
		if err != nil {
			break
		}
		input = append(input, tmp)
	}
	return input
}

func test() {
	// var tmp, tmp2 int
	// var i int
	// var input [3]int
	// input = append(input, 1)
	// input[0] = 3
	// for i := 0; i < 3; i++ {
	// fmt.Scanf("%d", &input[i])
	// }
	// fmt.Scanf("%d,%d", &input[0], &input[1])
	// fmt.Println(tmp)
	// fmt.Println(tmp2)
	// fmt.Printf("%v", input)
	a := make([][]int, 5)
	fmt.Printf("%v", a)
}
