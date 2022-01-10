package sort

//快排

func QuickSort(arr []int) {
	sort(arr, 0, len(arr)-1)
}

func sort(arr []int, left, right int) {
	if left < right {
		idx := partition(arr, left, right)
		sort(arr, left, idx-1)
		sort(arr, idx+1, right)
	}
}

func partition(arr []int, left, right int) int {
	p := arr[left]
	for left < right {
		for left < right && arr[right] >= p {
			right--
		}
		arr[left] = arr[right]
		for left < right && arr[left] <= p {
			left++
		}
		arr[right] = arr[left]
	}
	//复位
	arr[left] = p
	return left
}
