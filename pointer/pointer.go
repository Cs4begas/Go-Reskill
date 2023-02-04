package main

import "fmt"

func main() {
	d := 2
	doubleWithOutPointer(d)
	fmt.Printf("d without pointer = %d\n", d)
	double(&d)
	fmt.Printf("d = %d\n", d)

}

func double(d *int) {
	*d = *d * 2
}

func doubleWithOutPointer(d int) {
	d = d * 2
}
