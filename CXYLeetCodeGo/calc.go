package CXYLeetCodeGo

import "fmt"
import "rsc.io/quote"

func main() {
	fmt.Println("Hello World!")
	fmt.Println(add(1,2))
	dp := [3]int{1,2,3}
	fmt.Println(dp)
	fmt.Println(quote.Hello())
}

func add(num1 int,num2 int) int {
	return num1+num2
}