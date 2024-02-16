package main

func mergeSort(arr []int) {
	n := len(arr)
	if n == 1 {
		return
	}

	mid := n / 2
	l := make([]int, mid)
	r := make([]int, n-mid)

	copy(l, arr[:mid])
	copy(r, arr[mid:])

	mergeSort(l)
	mergeSort(r)
	merge(arr, l, r)
}

func merge(arr []int, l []int, r []int) {
	left := len(l)
	right := len(r)
	i := 0
	j := 0
	idx := 0

	for i < left && j < right {
		if l[i] < r[j] {
			arr[idx] = l[i]
			i++
		} else {
			arr[idx] = r[j]
			j++
		}
		idx++
	}

	for ll := i; ll < left; ll++ {
		arr[idx] = l[ll]
		idx++
	}

	for rr := j; rr < right; rr++ {
		arr[idx] = r[rr]
		idx++
	}
}
