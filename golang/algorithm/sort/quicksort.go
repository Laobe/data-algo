package sort

import (
	"fmt"
)

// var a [101]int
var n int

func QuickSort(left int, right int, a []int) {
	var i, j int
	fmt.Println(left)
	fmt.Println(right)
	if left >= right {
		return
	}
	// temp = a[left]
	i = left
	j = right
	for i != j {
		for a[j] >= a[left] && i < j {
			j--
		}
		for a[i] <= a[left] && i < j {
			i++
		}

		// 交换两个数在数组中的位置
		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}

	// 最后将基准数归位
	a[left], a[i] = a[i], a[left]
	fmt.Println(left)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Printf("%v\n", a)

	QuickSort(left, i-1, a[left:i])
	QuickSort(i+1, right, a[i:right+1])
	return
}
