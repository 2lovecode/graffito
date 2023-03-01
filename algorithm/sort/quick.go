package sort

func Quick(list []int, left int, right int) {
	start := left
	end := right

	if left >= right {
		return
	}
	flagNum := list[left]

	for left < right {
		for list[right] >= flagNum && right > left {
			right--
		}
		list[left] = list[right]
		for list[left] <= flagNum && right > left {
			left++
		}
		list[right] = list[left]
	}
	list[left] = flagNum

	Quick(list, start, left-1)
	Quick(list, right+1, end)
}
