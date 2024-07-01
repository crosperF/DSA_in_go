package main

import (
	"fmt"
)

type ArrayDs struct {
	array []int
}

func (a *ArrayDs) push(val int) {
	a.array = append(a.array, val)
}

func (a *ArrayDs) pop() int {
	if len(a.array) < 1 {
		return -1
	}
	popped := a.array[len(a.array)-1]
	a.array = a.array[:len(a.array)-1]
	return popped
}

func (a *ArrayDs) last() int {
	if len(a.array) < 1 {
		return -1
	}
	return a.array[len(a.array)-1]
}

func (a *ArrayDs) front() int {
	if len(a.array) < 1 {
		return -1
	}
	return a.array[0]
}

func (a *ArrayDs) PrintArr() {
	arr := a.array
	for _, num := range arr {
		fmt.Println(num)
	}
}

func CreateArray() *ArrayDs {
	a := ArrayDs{
		array: []int{},
	}
	return &a
}
