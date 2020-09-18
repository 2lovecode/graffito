package lib



func QuickSort(list []int, left int, right int) {
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

	QuickSort(list, start, left-1)
	QuickSort(list, right+1, end)
}
