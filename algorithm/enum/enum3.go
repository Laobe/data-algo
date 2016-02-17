package enum

import (
	"fmt"
)

// 火柴棍问题
//  _     _  _       _   _  _   _   _
// | | |  _| _| |_| |_  |_   | |_| |_|
// |_| | |_  _|   |  _| |_|  | |_|  _|
//
func Enum3() {
	var a, b, c, m int
	sum := 0
	fmt.Printf("%s", "请输入火柴棍数:")
	fmt.Scanf("%d", &m)

	for a = 0; a <= 1111; a++ {
		for b = 0; b <= 1111; b++ {
			c = a + b
			if fun(a)+fun(b)+fun(c) == m-4 {
				fmt.Printf("%d+%d=%d\n", a, b, c)
				sum++
			}
		}
	}
	fmt.Printf("能拼出的总数:%d\n", sum)
}

func fun(x int) int {
	num := 0
	f := [10]int{6, 2, 5, 5, 4, 5, 6, 3, 7, 6}
	for x/10 != 0 {
		num += f[x%10]
		x = x / 10
	}
	num += f[x]
	return num
}
