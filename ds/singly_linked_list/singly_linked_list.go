package main

import "fmt"

type node struct {
	data int
	next *node
}

var head = &node{}		//head of the list

func CreateNewList(data int) *node{
	newList := new(node)

	newList.data = data
	newList.next = nil
	head.next = newList

	return newList
}

func (n *node) InsertInTheBeginning(data int){
	newItem := new(node)
	newItem.data = data
	newItem.next = head.next
	head.next = newItem
}
func (n *node) InsertInTheMiddle(data , index int){
	i := 0
	newItem := new(node)
	newItem.data = data

	item := head.next
	for item.next != nil {
		i++
		if index == i{
			break
		}
		item = item.next
	}

	newItem.next = item.next
	item.next = newItem
}

func (n *node) InsertInTheEnd(data int){
	newItem := new(node)
	newItem.data = data
	newItem.next = nil

	item := head.next
	for item.next != nil {
		item = item.next
	}
	item.next = newItem

}

func (n *node) DeleteFirst(){
	head.next = head.next.next
}

func (n *node) DeleteLast(){
	item := head.next
	for item.next.next != nil {
		item = item.next
	}
	item.next = nil
}

func (n *node) DeleteFromMiddle(index int){
	i := 0
	item := head.next
	for item.next != nil {
		i++
		if index == i+1{
			item.next = item.next.next
		}
		item = item.next
	}

}

func (n *node) print(){
	item := head.next
	for item != nil {
		fmt.Print(item.data,"->")
		item = item.next
	}
}

func main(){
	newList := CreateNewList(5)
	newList.InsertInTheBeginning(10)
	newList.InsertInTheBeginning(12)
	newList.InsertInTheBeginning(14)
	newList.InsertInTheBeginning(16)

	newList.InsertInTheBeginning(11)
	newList.DeleteFirst()

	newList.InsertInTheMiddle(100,3)

	newList.DeleteLast()
	newList.InsertInTheEnd(7)

	newList.DeleteFromMiddle(4)
	newList.print()
}
