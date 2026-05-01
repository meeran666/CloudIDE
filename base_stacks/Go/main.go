package main

import (
	"fmt"
)

// -------- Bubble Sort --------
func bubbleSort(arr []int) []int {
	n := len(arr)
	res := make([]int, n)
	copy(res, arr)

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if res[j] > res[j+1] {
				res[j], res[j+1] = res[j+1], res[j]
			}
		}
	}
	return res
}

// -------- Selection Sort --------
func selectionSort(arr []int) []int {
	n := len(arr)
	res := make([]int, n)
	copy(res, arr)

	for i := 0; i < n; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if res[j] < res[minIdx] {
				minIdx = j
			}
		}
		res[i], res[minIdx] = res[minIdx], res[i]
	}
	return res
}

// -------- Insertion Sort --------
func insertionSort(arr []int) []int {
	res := make([]int, len(arr))
	copy(res, arr)

	for i := 1; i < len(res); i++ {
		key := res[i]
		j := i - 1

		for j >= 0 && res[j] > key {
			res[j+1] = res[j]
			j--
		}
		res[j+1] = key
	}
	return res
}

// -------- Merge Sort --------
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

// -------- Quick Sort --------
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)-1]
	left := []int{}
	right := []int{}

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, pivot), right...)
}

// -------- Main --------
func main() {
	data := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Original:", data)

	fmt.Println("Bubble Sort:", bubbleSort(data))
	fmt.Println("Selection Sort:", selectionSort(data))
	fmt.Println("Insertion Sort:", insertionSort(data))
	fmt.Println("Merge Sort:", mergeSort(data))
	fmt.Println("Quick Sort:", quickSort(data))
}
