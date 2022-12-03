package main

import "fmt"

func main() {
	num := 31
	// ใน ภาษา Go if ไม่มี วงเล็บอะไรไม่สำคัญ Go ตัดออกหมด
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else if num == 31 {
		fmt.Println(num, "is 31")
	} else {
		fmt.Println(num, "is odd")
	}
}
