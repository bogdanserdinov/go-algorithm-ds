package main

func main(){
	arr := [...]int{1,2,3,4,5,6,7,8,9,10}
	guessNum := 10
	findNum(arr,guessNum)
}

func findNum (arr [10]int,num int) (int,int){		//returns index,number
	min := 0
	max := len(arr)

	mid := (min+max)/2

	if arr[mid] == num {
		return mid,arr[mid]
	}
}
