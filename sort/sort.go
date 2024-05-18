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
