package main

import "fmt"

const (
	BEIJING  = iota*2 + 1 //iota = 0
	SHANGHAI              //iota = 1
	SHENZHEN              //iota = 2

	XINXIANG = iota + 999
)

func main() {
	const a = 10
	fmt.Println("a = ", a)
	fmt.Printf("BEIJING is %d\n", BEIJING)
	fmt.Printf("SHANGHAI is %d\n", SHANGHAI)
	fmt.Printf("SHENZHEN is %d\n", SHENZHEN)
	fmt.Printf("XINXIANG is %d\n", XINXIANG)
}
