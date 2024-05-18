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
