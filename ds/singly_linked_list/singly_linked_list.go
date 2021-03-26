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
	newList.InsertInTheEnd(7)

	newList.print()
}
