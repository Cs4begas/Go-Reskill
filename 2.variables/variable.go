package main

import "fmt"

// var name string = "Gopher นะจ๊ะ"
// name := "Gopher นะจ๊ะ" แตก เพราะ := ใช้ได้เฉพาะภายใน function body
// var ไว้ประกาศตัวแปร
func main() {
	//var name string = "Gopher นะจ๊ะ"
	// name := "Gopher นะจ๊ะ"
	word := "Gopher นะจ๊ะ"
	word += "Test"
	fmt.Printf("Hello, %v !!!!!!! \n", word)
	var name string // ประกาศตัวแปรเฉย ๆ Go จะให้ค่า Default เสมอ Default String = "" Int = 0 Boolean = False
	fmt.Printf("Hello, %v !!!!!!! \n", name)
	fmt.Printf("type: %T \n", name)  // print Type ของ variable %T
	fmt.Printf("name : %q \n", name) // print ค่าของ variable %q (quote จะเน้นคำที่อยู่ในตัวแปร)

	// !!! Go ไม่อนุญาติประกาศตัวแปรแล้วไม่ใช้ !!!!  รันแล้วจะแตก

}

/*
1. ประกาศ ไว้นอก function กับใน function ต่างกันที่
	นอก function อยู่ภายใต้ package เรียกว่า package scope
		ทุก function ที่อยุ่ภายใต้ package จะเห็นตัวแปร package scope ทั้งหมด
	{} คือ block scope ของ Go
2. Go สามารถ infer type ได้ (หมายความว่าเราไม่จำเป็นต้องประกาศ type ของตัวแปรได้ ถ้าเรามีค่าระบุทางด้านขวา)
3. ประกาศแบบ shortHand (ไม่มี var ใช้ :=) !!!! := (เรียกว่า colon function ไม่สามารถใช้นอก function ได้)
4. ประกาศตัวแปรเฉย ๆ (ตอนรันจะได้ empty string)
*/

/* Basic Type in Go
1. bool
2. string
3. int int8 int16 int32 int64
4. uint uint8 uint16 uint32 uint64 uintptr
5. byte // alias for uint8
6. rune // alias for int32 (rune คือ 1 ตัวอักษรจริง ๆ 1 char แต่ไม่เหมือน char ในภาษาอื่น ๆ)
	represent a Unicode code point
7. float32 float64

8. complex64 complex128 (Complex number)
*/

// func fucnum() {
// 	fmt.Println(name)
// }
