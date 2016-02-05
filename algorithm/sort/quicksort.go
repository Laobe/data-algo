package sort

var a [101]int
var n int

func quicksort(left int, right int) {
	var i, j, t, temp int
	if left > right {
		return
	}
	temp = a[left]
	i = left
	j = right
	for i != j {
		for a[j] >= temp && i < j {
			j--
		}
		for a[i] <= temp && i < j {
			i++
		}

		// 交换两个数在数组中的位置
		if i < j {
			a[i], a[j] = a[j], a[i]
		}
	}

	// 最后将基准数归位
	a[left] = a[i]
	a[i] = temp
}
