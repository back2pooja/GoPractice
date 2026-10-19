package main

import (
	"fmt"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	//fmt.Println("Even Numbers:")
	//printNumberEven()
	fmt.Print("sum: ")
	fmt.Print(summonArrayElement([]int{1, 2, 3, 4, 5}))

}

func printNumberEven() {

	for i := 2; i <= 100; i += 2 {
		fmt.Println(i)
	}
	fmt.Println()

}
