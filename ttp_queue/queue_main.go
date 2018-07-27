package main

import (
	"appliedgo/intqueue"
	"appliedgo/queue"
	"fmt"
)

func main() {
	fmt.Println("This is queue program")
	q := queue.Queue{}
	q.PutAny(1)
	q.PutAny(2)
	q.PutAny(3)
	q.PutAny(4)

	elem, err := q.GetAny()
	if err == nil {
		fmt.Println(elem)
	}

	elem1, err1 := q.GetAny()
	if err1 == nil {
		fmt.Println(elem1)
	}

	elem2, err2 := q.GetAny()
	if err2 == nil {
		fmt.Println(elem2)
	}

	elem3, err3 := q.GetAny()
	if err3 == nil {
		fmt.Println(elem3)
	}

	elem4, err4 := q.GetAny()
	if err4 == nil {
		fmt.Println(elem4)
	}
	forInt()
}

func forInt() {
	fmt.Println("This is queue program For int")
	q := intqueue.IntegerQueue{}
	q.PutAny(1)
	q.PutAny(2)
	q.PutAny(3)
	q.PutAny(4)

	elem, err := q.GetAny()
	if err == nil {
		fmt.Println(elem)
	}

	elem1, err1 := q.GetAny()
	if err1 == nil {
		fmt.Println(elem1)
	}

	elem2, err2 := q.GetAny()
	if err2 == nil {
		fmt.Println(elem2)
	}

	elem3, err3 := q.GetAny()
	if err3 == nil {
		fmt.Println(elem3)
	}

	elem4, err4 := q.GetAny()
	if err4 == nil {
		fmt.Println(elem4)
	}
}
