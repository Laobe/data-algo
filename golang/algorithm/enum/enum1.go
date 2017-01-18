package enum

import (
	"fmt"
)

// 计算奥数问题
// ***+***=***
// ********* in [1,2,3,4,5,6,7,8,9]
func Enum1() {
	a := make([]int, 10)
	books := make([]int, 10)
	var total int = 0
	var i, sum int
	for a[1] = 1; a[1] <= 9; a[1]++ {
		for a[2] = 1; a[2] <= 9; a[2]++ {
			for a[3] = 1; a[3] <= 9; a[3]++ {
				for a[4] = 1; a[4] <= 9; a[4]++ {
					for a[5] = 1; a[5] <= 9; a[5]++ {
						for a[6] = 1; a[6] <= 9; a[6]++ {
							for a[7] = 1; a[7] <= 9; a[7]++ {
								for a[8] = 1; a[8] <= 9; a[8]++ {
									for a[9] = 1; a[9] <= 9; a[9]++ {
										for i = 0; i <= 9; i++ {
											books[i] = 0
										}
										for i = 1; i <= 9; i++ {
											books[a[i]] = 1
										}
										// 统计工出现了多少个不同的数
										sum = 0
										for _, book := range books {
											sum += book
										}
										if sum == 9 && a[1]*100+a[2]*10+a[3]+a[4]*100+a[5]*10+a[6] == a[7]*100+a[8]*10+a[9] {
											total++
											fmt.Printf("%d%d%d+%d%d%d=%d%d%d\n", a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9])
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Printf("total=%d", total/2)
}
