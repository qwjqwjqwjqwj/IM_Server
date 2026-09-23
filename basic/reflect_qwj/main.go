package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `name:"user's name"`
	Id   int    `id:"user's id" ex:"it's the only" `
}

func main() {

	//var u1 User = User{"lihua", 128124}
	//Do_Type_Value(u1)

	u2 := &User{"liqing", 4780700}
	find_Tag(u2)
}
func reflect_Example() {
	//--------------------reflet.Type的Elem()用法--------------------
	//.Elem()，得先确认"它是不是容器"。
	/*
		Kind		示例类型				Elem () 返回
		Ptr 指针	*int、*User			指针指向的底层类型 int / User
		Slice 切片	[]string			切片元素类型 string
		Array 数组	[10]int64			数组元素类型 int64
		Map 映射	map[string]bool		map 的 value 类型 bool（key 需要用 Key()）
		Chan 通道	chan float64		channel 里传递的元素类型 float64
	*/
	a := &User{}
	b := reflect.TypeOf(a)
	fmt.Println(b)        //*User
	fmt.Println(b.Elem()) //main.User

	//--------------------reflect.Value的Elem()--------------------
	x := 10
	v := reflect.ValueOf(&x) // 这是个指针的值
	fmt.Println(v)
	v.Elem() // 10  ← 指针指向的实际值
	fmt.Println(v.Elem())
}

// ----------reflect
func Do_Type_Value(input interface{}) {
	ty := reflect.TypeOf(input)
	fmt.Println("type of Input is :", ty)            //main.User
	fmt.Println("Number of type is:", ty.NumField()) //获取字段个数

	vl := reflect.ValueOf(input) //     {lihua 128124}
	fmt.Println("value of u1 is :", vl)

	for i := 0; i < ty.NumField(); i++ {
		field := ty.Field(i)
		value := vl.Field(i)
		fmt.Println(field.Name, field.Type, value)
	}

}

// ----------tag
func find_Tag(str interface{}) {
	t := reflect.TypeOf(str).Elem()

	for i := 0; i < t.NumField(); i++ {
		x := t.Field(i).Tag
		fmt.Println(x)
		y := x.Get("ex")
		fmt.Println("ex:", y)

	}
}
