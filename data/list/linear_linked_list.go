package list

import (
	"fmt"
)

type LinerLinkedList struct {
	head *node
}

type node struct {
	value any
	next  *node
}

func NewLinerLinkedList() *LinerLinkedList {
	return &LinerLinkedList{}
}

func (l *LinerLinkedList) String() string {
	if l.head == nil {
		return "nil"
	}
	listLen := l.Len()
	elements := make([]any, 0, listLen)

	current := l.head
	for i := 0; i < listLen; i++ {
		elements = append(elements, current.value)
		current = current.next
	}

	return fmt.Sprintf("%v", elements)
}

func (l *LinerLinkedList) Append(v any) {
	n := &node{value: v}
	if l.head == nil {
		l.head = n
		return
	}

	current := l.head

	for current.next != nil {
		current = current.next
	}

	current.next = n
}

func (l *LinerLinkedList) ValueOf(i int) (any, bool) {
	const ok = true
	if l.head == nil {
		return nil, !ok
	}
	current := l.head

	// 探すのにO(n)
	for j := 0; j <= i; j++ {
		if j == i {
			return current.value, ok
		}
		if current.next == nil {
			return nil, !ok
		}
		current = current.next
	}
	return nil, !ok
}

func (l *LinerLinkedList) Len() int {
	if l.head == nil {
		return 0
	}
	count := 0
	current := l.head
	for current.next != nil {
		count++
		current = current.next
	}

	count++
	return count
}

func (l *LinerLinkedList) findNodeByIndex(i int) *node {
	if l.head == nil {
		return nil
	}
	current := l.head

	// 探すのにO(n)
	for j := 0; j <= i; j++ {
		if j == i {
			return current
		}
		if current.next == nil {
			return nil
		}
		current = current.next
	}
	return current
}

func (l *LinerLinkedList) Insert(v any, i int) error {
	if i < 0 {
		return fmt.Errorf("i must be greater than 0, but got %d", i)
	}
	listLen := l.Len()
	if i > listLen {
		return fmt.Errorf("expected i is %d or less, but got %d", listLen, i)
	}

	if i == listLen {
		l.Append(v)
		return nil
	}
	prevNode := l.findNodeByIndex(i - 1)
	nextNode := l.findNodeByIndex(i) // 現在iのインデックスに存在するnode

	if prevNode == nil {
		return fmt.Errorf("cannot find node at index %d", i-1)
	}

	if nextNode == nil {
		return fmt.Errorf("cannot find node at index %d", i)
	}

	newNode := &node{value: v, next: nextNode}
	prevNode.next = newNode

	return nil
}

func (l *LinerLinkedList) Remove(i int) error {
	if i < 0 {
		return fmt.Errorf("i must be greater than 0, but got %d", i)
	}
	listLen := l.Len()
	if i >= listLen {
		return fmt.Errorf("expected i is less than %d, but got %d", listLen, i)
	}

	if i == 0 {
		l.head = l.head.next
		return nil
	}

	if i == listLen-1 {
		prevNode := l.findNodeByIndex(i - 1)
		if prevNode == nil {
			return fmt.Errorf("cannot find node at index %d", i-1)
		}
		prevNode.next = nil
		return nil
	}

	prevNode := l.findNodeByIndex(i - 1)
	if prevNode == nil {
		return fmt.Errorf("cannot find node at index %d", i-1)
	}

	if prevNode.next == nil {
		return fmt.Errorf("cannot find node at index %d", i)
	}

	prevNode.next = prevNode.next.next

	return nil
}
