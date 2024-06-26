package main

import (
	"fmt"

	"github.com/Ryota-Onuma/go-data-structure-algorithm/data/hash_table"
	"github.com/Ryota-Onuma/go-data-structure-algorithm/data/list"
	"github.com/Ryota-Onuma/go-data-structure-algorithm/sort"
)

func main() {
	tryHashTable()
	tryLinerLinkedList()
	tryBubbleSort()
	trySelectionSort()
	tryInsertionSort()
	tryQuickSort()
	tryMergeSort()
}

func tryHashTable() {
	fmt.Println("\n--------TryHashTable--------\n")
	hash := hash_table.New()
	hash.Insert("お腹がすいた", "ご飯を食べる")
	hash.Insert("ラーメン", "美味しい")
	hash.Insert("ラーメン", "食べたい")
	hash.Insert("ラーメン2", "食べたい2")
	v, ok := hash.Get("ラーメン")
	fmt.Println(v, ok)
	hash.Delete("ラーメン")
	v2, ok2 := hash.Get("ラーメン")
	fmt.Println(v2, ok2)
	fmt.Println(hash)
}

func tryLinerLinkedList() {
	fmt.Println("\n--------TryLinerLinkedList--------\n")
	l := list.NewLinerLinkedList()
	elments := [10]int{5, 2, 4, 6, 1, 3, 8, 7, 10, 9}
	for _, e := range elments {
		l.Append(e)
	}

	fmt.Println(l.ValueOf(9))
	fmt.Println(l.Len())
	fmt.Println(l)
	if err := l.Insert(100, 8); err != nil {
		fmt.Println(err)
	}
	fmt.Println(l)
	if err := l.Remove(l.Len()-1); err != nil {
		fmt.Println(err)
	}
	fmt.Println(l)
}

func tryBubbleSort() {
	fmt.Println("\n--------TryBubbleSort--------\n")
	arr := []int{5, 3, 8, 4, 2, 1, 9, 7, 6, 10, 15, 20, 13, 12, 11, 14, 19, 18, 17, 16}
	fmt.Println("Try ASC")
	fmt.Println("Before:", arr)
	sort.BubbleSort(arr, sort.ASC)
	fmt.Println("After:", arr)
	fmt.Println("Try DESC")
	fmt.Println("Before:", arr)
	sort.BubbleSort(arr, sort.DESC)
	fmt.Println("After:", arr)
}

func trySelectionSort() {
	fmt.Println("\n--------TrySelectionSort--------\n")
	arr := []int{5, 3, 8, 4, 2, 1, 9, 7, 6, 10, 15, 20, 13, 12, 11, 14, 19, 18, 17, 16}
	fmt.Println("Try ASC")
	fmt.Println("Before:", arr)
	sort.SelectionSort(arr, sort.ASC)
	fmt.Println("After:", arr)
	fmt.Println("Try DESC")
	fmt.Println("Before:", arr)
	sort.SelectionSort(arr, sort.DESC)
	fmt.Println("After:", arr)
}

func tryInsertionSort() {
	fmt.Println("\n--------TryInsertionSort--------\n")
	arr := []int{5, 3, 8, 4, 2, 1, 9, 7, 6, 10, 15, 20, 13, 12, 11, 14, 19, 18, 17, 16}
	fmt.Println("Try ASC")
	fmt.Println("Before:", arr)
	sort.InsertionSort(arr, sort.ASC)
	fmt.Println("After:", arr)
	fmt.Println("Try DESC")
	fmt.Println("Before:", arr)
	sort.InsertionSort(arr, sort.DESC)
	fmt.Println("After:", arr)
}

func tryQuickSort() {
	fmt.Println("\n--------TryQuickSort--------\n")
	arr := []int{5, 3, 8, 4, 2, 1, 9, 7, 6, 10, 15, 20, 13, 12, 11, 14, 19, 18, 17, 16}
	fmt.Println("Try ASC")
	fmt.Println("Before:", arr)
	sort.QuickSort(arr, sort.ASC)
	fmt.Println("After:", arr)
	fmt.Println("Try DESC")
	fmt.Println("Before:", arr)
	sort.QuickSort(arr, sort.DESC)
	fmt.Println("After:", arr)
}

func tryMergeSort() {
	fmt.Println("\n--------TryMergeSort--------\n")
	arr := []int{5, 3, 8, 4, 2, 1, 9, 7, 6, 10, 15, 20, 13, 12, 11, 14, 19, 18, 17, 16}
	fmt.Println("Try ASC")
	fmt.Println("Before:", arr)
	sort.MergeSort(arr, sort.ASC)
	fmt.Println("After:", arr)
	fmt.Println("Try DESC")
	fmt.Println("Before:", arr)
	sort.MergeSort(arr, sort.DESC)
	fmt.Println("After:", arr)
}
