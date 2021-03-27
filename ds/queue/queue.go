package main

import "fmt"

const max int = 5	//max size of queue

type queue struct {
	front int
	rear int
	items []int
}

func newQueue() *queue{
	res := new(queue)
	res.rear = -1
	res.front = -1

	return res
}

func(q *queue) isFull() bool{
	if q.rear == max - 1{
		return true
	}else {
		return false
	}
}

func(q *queue) isEmpty() bool{
	if	q.front == -1 && q.rear == -1{
		return true
	}else{
		return false
	}
}

func(q *queue) enqueue(data int) {
	if q.isFull(){
		fmt.Println("queue is full, could not insert new data")
		return
	}else {
		if q.front == -1 {
			q.front++
		}
		q.items = append(q.items, data)
		q.rear++
	}
}

func(q *queue) dequeue(){
	if q.isEmpty(){
		fmt.Println("queue is empty, could not delete data")
		return
	}else{
		if q.front >= q.rear{	//if in queue was only 1 element
			q.front = -1
			q.rear = -1
		}else{
			q.front++
		}
	}
}

func(q *queue) peek() int{
	if q.isEmpty(){
		fmt.Println("queue is empty, could not peek first item")
		return -1
	}else{
		return q.items[q.front]
	}
}

func(q *queue) print(){
	for i:= q.front;i <= q.rear;i++{
		if i == q.rear{
			fmt.Print(q.items[i])
			break
		}
		fmt.Print(q.items[i],"<-")
	}
}

func main() {
	q := newQueue()
	q.enqueue(10)
	q.enqueue(11)
	q.enqueue(12)
	q.enqueue(13)

	first := q.peek()
	fmt.Println("first in queue -> ",first )
	q.print()

	fmt.Println()

	q.dequeue()
	fmt.Println("---after dequeue---")
	q.print()
}
