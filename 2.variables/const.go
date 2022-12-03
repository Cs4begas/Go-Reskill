package main

import "fmt"

func main() {
	// var day string = "Monday"
	// fmt.Println("day :", day)

	// day = "ม่วนเดย์"
	// fmt.Println("day :", day) // เปลี่ยนค่า ตกกะปิ

	// const day string = "Monday"
	// fmt.Println("day :", day)

	// day = "ม่วนเดย์" // แตก ไม่สามารถเปลี่ยนค่าได้
	// fmt.Println("day :", day)

	/* ประกาศแบบปกติ
	const Sunday = 0
	const Monday = 1
	const Tuesday = 2
	const Wednesday = 3
	const Thursday = 4
	const Friday = 5
	const Saturday = 6
	*/

	/* ประกาศแบบ const เดียว
	const (
		Sunday    = 0
		Monday    = 1
		Tuesday   = 2
		Wednesday = 3
		Thursday  = 4
		Friday    = 5
		Saturday  = 6
	)
	*/

	/* ประกาศแบบ iota
	 */
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println("Sunday :", Sunday)
	fmt.Println("Monday :", Monday)
	fmt.Println("Tuesday :", Tuesday)
	fmt.Println("Wednesday :", Wednesday)
	fmt.Println("Thursday :", Thursday)
	fmt.Println("Friday :", Friday)
	fmt.Println("Saturday :", Saturday)

}
