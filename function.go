package main

import "fmt"

func main() {
	// info("Gopher", "Hello", 11)
	// info("Nattapan", "Go Home", 15)

	// fmt.Println(today())

	a, b := swap("Hello", "World")
	println(a, b)
}

// Go type อยู่ข้างหลัง
// func info(name string, msg string) {
// 	fmt.Printf("My name is %s\n", name)
// 	fmt.Printf("My message is %s\n", msg)

// }

// Go ถ้า Parameters ติดกัน type เหมือนกันไม่ต้องเขียน 2 อัน
func info(name, msg string, age int) {
	fmt.Printf("My name is %s\n", name)
	fmt.Printf("My message is %s\n", msg)
	fmt.Print("My age is ", age, "\n")

}

func today() string { //หลังสุด return type
	return "มื้อนี้"
}

// Go สามารถ Return multiple value ได้
func swap(x, y string) (string, string) {
	return y, x
}
