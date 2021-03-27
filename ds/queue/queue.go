package main

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

func main() {
	new := newQueue()
}