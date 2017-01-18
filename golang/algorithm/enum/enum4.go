package enum

import (
	"fmt"
)

// 全排列问题
// 123
// 123, 132, 213, 231, 312, 321
func Enum4() {
	var a, b, c, d int
	var total int = 0
	for a = 1; a <= 4; a++ {
		for b = 1; b <= 4; b++ {
			for c = 1; c <= 4; c++ {
				for d = 1; d <= 4; d++ {
					if a != b && a != c && a != d && b != c && b != d && c != d {
						total++
						fmt.Printf("%d%d%d%d\n", a, b, c, d)
					}
				}
			}
		}
	}
	fmt.Printf("total=%d", total)
}
