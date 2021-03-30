package main

import (
	"fmt"
	"log"
)

type deque struct{
	arr [20]int
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
	if (d.front == 0 && d.rear == d.size - 1) || (d.front == d.rear + 1){
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

func (d *deque) insertRear(data int) {
	if d.isFull(){
		log.Println("deque is full")
		return
	}

	if d.rear == -1{
		d.front = 0
		d.rear = 0
	}else if d.rear == d.size - 1{
		d.rear = 0
	}else {
		d.rear += 1
	}

	d.arr[d.rear] = data
}

func(d *deque) deleteFront(){
	if d.isEmpty(){
		log.Println("deque is empty")
		return
	}

	if d.rear == d.front{	//if 1 element
		d.rear = 0
		d.front = 0
	}else if d.front == d.size - 1{
		d.front = 0
	}else {
		d.front += 1
	}
	d.arr[d.front] = 0
}

func (d *deque) deleteRear() {
	if d.isEmpty(){
		log.Println("deque is empty")
		return
	}

	if d.front == d.rear{
		d.front = 0
		d.rear = 0
	}else if d.rear == 0{
		d.rear = d.size - 1
	}else {
		d.rear -= d.rear
	}

	d.arr[d.rear] = 0
}

func(d *deque) getFront(){
	if d.isEmpty(){
		log.Println("deque is empty")
		return
	}
	fmt.Println(d.arr[d.front])
}

func(d *deque) getRear(){
	if d.isEmpty() || d.rear < 0{
		log.Println("deque is empty")
		return
	}
	fmt.Println(d.arr[d.rear])
}

func main(){
	size := 10
	d := initializeDeque(size)
	d.insertFront(5)
	d.insertFront(10)
	d.getFront()
	d.deleteFront()
	d.getFront()

	d.insertRear(1)
	d.insertRear(2)
	d.getRear()
	d.deleteRear()
	d.getRear()

	fmt.Println(d)
}
