package array

import (
	"fmt"
	"testing"
)

func TestBubSort(t *testing.T){
	eleType:=[]int{3,5,1,-7,4,9,-6,8,10,4}
	fmt.Println(eleType)
	bubbleSort(eleType)
	fmt.Println(eleType)
}

func TestQuickSort(t *testing.T){
	eleType:=[]int{3,5,1,-7,4,9,-6,8,10,4}
	fmt.Println(eleType)
	quickSort(eleType,0,len(eleType)-1)
	fmt.Println(eleType)
}
