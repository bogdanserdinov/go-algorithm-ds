package main

import "fmt"

//const max int = 10

type stack struct{
	s []int
	top int
}

func createNewStack() *stack{
	s := new(stack)
	s.top = -1
	return s
}

func (st *stack) push(num int){
	st.top++
	st.s = append(st.s,num)
}

func (st *stack) isEmpty() bool{
	if st.top == -1{
		return true
	} else{
		return false
	}
}

func (st *stack) getLength() int{
	return len(st.s)
}

func (st *stack) pop() {
	res := st.isEmpty()
	if res == true {
		fmt.Println("stack is empty")
		return
	}else{
		st.s = st.s[:len(st.s)-1]
		st.top--
	}
}
func (st *stack) print(){
	for i := 0; i < st.top+1;i++{
		if i == st.top+1{
			fmt.Print(st.s[i])
			break
		}
		fmt.Print(st.s[i],"->")
	}
}

func main(){
	stack := createNewStack()

	stack.push(5)
	stack.push(10)
	stack.pop()
	stack.push(15)

	fmt.Println(stack.getLength())
	fmt.Println(stack.isEmpty())
	stack.print()
}