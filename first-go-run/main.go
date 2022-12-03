package main

import "fmt" // build in Library จัดการเกี่ยวกับการ Print IO ต่าง ๆ

func main() {
	// ...
	/* ชื่อ Package.ชื่อ Function Go บอกจริง ๆ มี ; (semicolon) แต่ complier ไปเติมเองตอน compile ให้ */
	fmt.Println("Hello, World!") // Print line

	fmt.Printf("Hello, %s !!!!!!! \n", "Gopher") // Print Template String

	fmt.Printf("Hello, %d !!!!!!! \n", 123) // Print Template Integer

	// % พระเจ้า คือ %v print ค่าอะไรก็ได้
	fmt.Printf("Hello, %v !!!!!!! \n", 123) // Print Template %v

	fmt.Printf("Hello, %v : %v !!!!!!! \n", 123, "Gopher") // Print Template %v
}
