package main //有main函数一定是main包

import (
	"fmt"
	"maps"
	"strconv"
)

func f1(a int, b string) (string, int) {

	return b + strconv.Itoa(a), a
}

func printArray(array []int) {
	array[0] = 100
	for a, b := range array {
		fmt.Println("index = ", a, "value = ", b)
	}
}
func main() {

	//变量声明
	var a int = 10
	var a1 = 10
	a2 := 10
	var a3 string
	fmt.Printf("%v", a)
	fmt.Println("a = ", a)
	fmt.Printf("type of a is %T\n", a)
	fmt.Println("a1,a2,a3=", a1, a2, a3, "//")

	x, y := 1, "hello"
	fmt.Println("x = ", x)
	fmt.Println("y = ", y)

	//函数调用
	fmt.Println(f1(1, "func_back"))

	//固定数组
	var array1 [10]int = [10]int{1, 2, 3, 4}
	fmt.Println("array1 = ", array1)
	var array2 = [4]string{0: "league", 2: "of", 3: "legends"}
	fmt.Println("array2 = ", array2)

	//动态数组 (切片
	array3 := []int{1, 2, 3, 4}
	printArray(array3) //引用传递
	fmt.Printf("array3:   %#v \n", array3)

	var slice1 []int
	//slice1[0] = 100 //报错，切片没有分配内存

	slice1 = make([]int, 3)
	slice1 = make([]int, 3, 5) //分配内存 预分配5个位置 但只能读写前三个值
	fmt.Println("slice1 = ", slice1)
	fmt.Printf("slice1 lenght = %d ,slice1 cap =%d \n ", len(slice1), cap(slice1))

	slice1 = append(slice1, 1)
	fmt.Printf("slice1 append lenght = %d ,slice1 append cap =%d \n ", len(slice1), cap(slice1))

	slice1 = append(slice1, 77, 88)
	fmt.Printf("slice1 over append lenght = %d ,slice1 over append cap =%d \n ", len(slice1), cap(slice1))

	//切片是引用类型，修改一个切片会影响到另一个切片
	s2 := slice1
	s2[0] = 999
	fmt.Printf("slice1 = %#v ,s2 = %#v \n ", slice1, s2)

	//map
	var m1 map[int]string
	m1 = make(map[int]string)
	m1[0] = "league"
	m1[2] = "of"
	m1[4] = "legends"
	fmt.Println("m1 = ", m1)
	m2 := maps.Clone(m1)
	m2[0] = "LOL"
	delete(m1, 0)
	fmt.Println("m1 = ", m1, "m2 = ", m2)

}
