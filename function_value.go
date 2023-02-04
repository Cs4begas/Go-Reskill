package main

import "fmt"

var name string = "Goku"

// แบบที่ 1 var plus = func(a int, b int) int { return a + b } // function value (แค่ value มันคือ function)
// แบบที่ 2 var plus func(a int, b int) int = func(a int, b int) int { return a + b } // บอก type ก็ได้

func main() {
	fmt.Printf("value %v\n", name)
	fmt.Printf("type %T\n", name)

	plus := func(a int, b int) int { return a + b } // function value (แค่ value มันคือ function)

	p := plus(1, 2)
	fmt.Println("1+2 =", p)
	fmt.Printf("type %T\n", plus)

}

// func plus(a int, b int) int {
// 	return a + b
// }

// Question : function value เอาไว้ทำอะไรถามพี่หน๋องใน discords
