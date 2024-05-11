package hash_table

import (
	"fmt"
	"hash/fnv"
	"log"
	"strings"
)

const (
	bucketCount        = 10
	backetElementCount = 8
)

func New() *hashTable {
	return &hashTable{
		buckets: make([][backetElementCount]bucketElement, bucketCount),
	}
}

type hashTable struct {
	buckets [][backetElementCount]bucketElement
}

type bucketElement struct {
	key   string
	value any
}

func (e bucketElement) IsEmpty() bool {
	return e.key == ""
}

func (h *hashTable) String() string {
	var str string
	for i, bucket := range h.buckets {
		var strs []string
		for _, element := range bucket {
			if !element.IsEmpty() {
				strs = append(strs, fmt.Sprintf("%s -> %v", element.key, element.value))
			}
		}
		str += fmt.Sprintf("[%d] %s\n", i, strings.Join(strs, ", "))
	}
	return str
}

func (h *hashTable) Insert(key string, value any) {
	bucketIndex := calcBucketIndex(key, len(h.buckets))

	foundIndex := -1
	for i, element := range h.buckets[bucketIndex] {
		if element.key == key {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		for i, element := range h.buckets[bucketIndex] {
			if element.IsEmpty() {
				h.buckets[bucketIndex][i] = bucketElement{key: key, value: value}
				return
			}
		}
		// TODO: ここでバケットがいっぱいの場合の処理を追加する
		// TODO: リンクリストを使ってバケットを拡張する
		log.Fatal("バケットがいっぱいです")
	} else {
		h.buckets[bucketIndex][foundIndex] = bucketElement{key: key, value: value}
	}
}

func (h *hashTable) Get(key string) (any, bool) {
	const ok = true
	bucketIndex := calcBucketIndex(key, len(h.buckets))
	for _, element := range h.buckets[bucketIndex] {
		if element.key == key {
			return element.value, ok
		}
	}
	return nil, !ok
}

func (h *hashTable) Delete(key string) {
	bucketIndex := calcBucketIndex(key, len(h.buckets))

	deleteIndex := -1
	for i, element := range h.buckets[bucketIndex] {
		if element.key == key {
			deleteIndex = i
			break
		}
	}

	if deleteIndex != -1 {
		newBucket := [backetElementCount]bucketElement{}
		newBucketIdx := 0
		for i, element := range h.buckets[bucketIndex] {
			if i != deleteIndex {
				newBucket[newBucketIdx] = element
				newBucketIdx++
			}
		}
		h.buckets[bucketIndex] = newBucket
	}
}

// バケットのインデックスを計算する
func calcBucketIndex(key string, bucketCount int) int {
	hash := fnv.New64()
	hash.Write([]byte(key))
	hashed := hash.Sum64()
	index := uint(hashed) % uint(bucketCount)
	return int(index)
}
