package main

import (
	"fmt"
	"log"
)

type deque struct{
	arr [10]int
	front int
	rear int
	size int
}

func initializeDeque(size int) *deque{
	d := new(deque)
	d.front = -1
	d.rear = -1
	d.size = size

	return d
}

func (d *deque) isFull() bool{
	if (d.front == 0 && d.rear == d.size - 1) || (d.front == d.rear - 1){
		return true
	}
	return false
}

func (d *deque) isEmpty() bool{
	return d.front == -1
}

func (d *deque) insertFront(data int){
	if d.isFull() {
		log.Println("deque is full")
		return
	}

	if d.front == -1{
		d.front = 0
		d.rear = 0
	}else if d.front == 0{
		d.front = d.size - 1
	}else {
		d.front -= 1
	}

	d.arr[d.front] = data
}

func main(){
	size := 5
	d := initializeDeque(size)
	d.insertFront(5)
	d.insertFront(10)
	d.insertFront(11)

	fmt.Println(d)
}
