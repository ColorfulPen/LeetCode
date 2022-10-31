package calc

import "fmt"

// func main() {
// 	fmt.Println("Hello World!")
// 	fmt.Println(add(1,2))
// 	dp := [3]int{1,2,3}
// 	fmt.Println(dp)
// 	fmt.Println(quote.Hello())
// }

func Add(num1 int, num2 int) int {
	fmt.Println("hello")
	return num1 + num2
}

func add(num1 int, num2 int) int {
	fmt.Println("hello")
	return num1 + num2
}

func Get(index int) (ret int) {
	defer func ()  {
		if r:=recover();r!=nil {
			fmt.Println("spmm+    ",r)
			ret=-1
		}
	}()
	arr:=[3]int{1,2,3}
	return  arr[index]
}
