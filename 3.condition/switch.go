package main

import "time"

func main() {
	today := time.Monday

	switch today { // switch ใน go ไม่ต้องเขียน break เข้า เคสไหนจะเข้า case นั้นแล้วจบเลย
	//fall through คือ ถ้าเข้า case นั้น จะเข้า case ถัดไปด้วย
	case time.Monday:
		println("Today is Monday")
		fallthrough
	case time.Saturday, time.Sunday: // เสาร์หรืออาทิตย์
		println("Today is weekend")
	default:
		println("Today is not Monday")
	}
}
