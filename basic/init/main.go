package main

import (
	"fmt"
	_ "init/lib1" // 匿名导包 可以不必使用lib1提供的方法
	. "init/lib2" // .导包 使用时可省略包名
)

func main() {
	//fmt.Println(lib1.Lib1Test("lib1_test"))
	a := 1
	var p = &a
	fmt.Println(p, ":", *p)
	fmt.Println(Lib2Test("lib2_test"))

}
