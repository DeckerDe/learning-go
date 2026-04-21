package main

import "fmt"

type number interface {
	~int | ~float64
}

func double[T number](x T) T {
	return x * 2
}

type Printable interface {
	number
	fmt.Stringer
}

func printNum[T Printable](num T) {
	fmt.Println(num)
}

type Element[T comparable] struct {
	val  T
	next *Element[T]
}

type LinkedList[T comparable] struct {
	head Element[T]
	last Element[T]
}

func (l *LinkedList[T]) Add(member T) {
	newElement := &Element[T]{
		val: member,
	}
	l.last.next = newElement
	l.last = *newElement
}

func (l *LinkedList[T]) Insert(member T, pos int) {
	currentElement := l.head
	for i := 0; i <= pos; i++ {
		if i == pos-1 {
			newElement := &Element[T]{
				val:  member,
				next: currentElement.next,
			}
			currentElement.next = newElement
		}
	}
}

func (l LinkedList[T]) Index(member T) int {
	pos, currentElement := 0, l.head
	for {
		if currentElement.val == member {
			return pos
		}
		if currentElement.next == nil {
			return -1
		}
		pos++
	}
}

func main() {

}
