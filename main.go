package main

import (
	"fmt"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	//fmt.Println("Even Numbers:")
	//printNumberEven()
	//fmt.Print("sum: ")
	//fmt.Print(summonArrayElement([]int{1, 2, 3, 4, 5}))

	Amax, _ := largestNum([]int{})
	fmt.Println("largest:", Amax)

	//using sequence
	fmt.Println("this is the first statement")
	fmt.Println("this is the second statement")
	x := 42 //third statement
	y := 5  // fourth statement
	fmt.Println(x, y)

	if x < 42 {
		fmt.Println("x is greater than 42.")
	} else if x == 42 {
		fmt.Println("x is equal to 42.")
	} else {
		fmt.Println("x is greater than 42.")
	}
}

func printNumberEven() {

	for i := 2; i <= 100; i += 2 {
		fmt.Println(i)
	}
	fmt.Println()

}
