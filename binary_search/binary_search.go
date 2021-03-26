package main

import (
	"errors"
	"fmt"
	"log"
	"sort"
)

func main(){
	arr := []int{1,2,3,4,5,6,7,8,9,10}
	sort.Ints(arr)
	guessNum := 1
	index,num,err := findNum(arr,guessNum)
	if err != nil{
		log.Println(err.Error())
	}
	fmt.Println("result ->\tindex = ",index,"num = ",num)
}

func findNum (arr []int,num int) (int,int,error){		//returns index,number
	min := 0
	max := len(arr)-1

	//mid := (min+max)/2
	var mid int

	for min <= max{
		mid = (min+max)/2
		if arr[mid] == num {
			return mid,arr[mid],nil

		}
		if	arr[mid] > num {
			min = mid + 1
			//mid = (min+max)/2
		}
		if	arr[mid] < num {
			max = mid - 1
			//mid = (min+max)/2
		}
	}
	err := errors.New("could not find guess number")
	return 0,0,err

}
