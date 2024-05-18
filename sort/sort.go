package sort

type Order string

const (
	ASC  Order = "ASC"
	DESC Order = "DESC"
)

func BubbleSort(arr []int, order Order) {
	for i := len(arr); i > 0; i-- {
		for j := 0; j < i-1; j++ {
			if order == DESC {
				if arr[j] < arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			} else {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
	}
}

func SelectionSort(arr []int, order Order) {
	for i := 0; i < len(arr); i++ {
		foundIndex := i
		for j := i + 1; j < len(arr); j++ {
			if order == DESC {
				if arr[foundIndex] < arr[j] {
					foundIndex = j
				}
			} else {
				if arr[foundIndex] > arr[j] {
					foundIndex = j
				}
			}
		}
		arr[i], arr[foundIndex] = arr[foundIndex], arr[i]
	}
}

func InsertionSort(arr []int, order Order) {
	if order == DESC {
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] < arr[i+1] {
				for j := i; j >= 0; j-- {
					if arr[j] < arr[j+1] {
						arr[j], arr[j+1] = arr[j+1], arr[j]
					}
				}
			}
		}
	} else {
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				for j := i; j >= 0; j-- {
					if arr[j] > arr[j+1] {
						arr[j], arr[j+1] = arr[j+1], arr[j]
					}
				}
			}
		}
	}
}

func QuickSort(arr []int, order Order) {
	copy(arr, quickSort(arr, order))
}

func quickSort(arr []int, order Order) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]

	var left, right []int
	if order == DESC {
		for i := 1; i < len(arr); i++ {
			if arr[i] >= pivot {
				left = append(left, arr[i])
			} else {
				right = append(right, arr[i])
			}
		}
	} else {
		for i := 1; i < len(arr); i++ {
			if arr[i] <= pivot {
				left = append(left, arr[i])
			} else {
				right = append(right, arr[i])
			}
		}
	}
	var result []int
	result = append(result, quickSort(left, order)...)
	result = append(result, pivot)
	result = append(result, quickSort(right, order)...)
	return result
}
