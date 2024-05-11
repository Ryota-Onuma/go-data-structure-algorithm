package main

import (
	"fmt"

	"github.com/Ryota-Onuma/go-data-structure-algorithm/data/hash_table"
)

func main() {
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
