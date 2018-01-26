package mergeSort

func merge(left []int, right []int) []int {
	result := []int{}
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	for ; l < len(left); l++ {
		result = append(result, left[l])
	}
	for ; r < len(right); r++ {
		result = append(result, right[r])
	}
	return result
}

func mergeSortGoroutine(data []int, rChan chan []int) {
	if len(data) <= 1 {
		rChan <- data
		return
	}
	mid := len(data) / 2
	x := make(chan []int)
	y := make(chan []int)

	go mergeSortGoroutine(data[:mid], x)
	go mergeSortGoroutine(data[mid:], y)

	ld := <-x
	rd := <-y

	close(x)
	close(y)
	rChan <- merge(ld, rd)
	return
}

func mergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	mid := len(data) / 2

	l := mergeSort(data[:mid])
	r := mergeSort(data[mid:])

	return merge(l, r)
}
