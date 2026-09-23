package main

import "fmt"

// 定义父类接口  （指针类型
type Animal interface { //接口方法如果实例有指针接收者 那么只能用指针实现
	//如果接口方法均为值接收者 那么指针和非指针类型都可以实现
	Speak() string
	Name() string
}

type Empty interface{} //空接口

// 实现接口 - Dog
type Dog struct {
	name string
}

func (Dog) Speak() string {
	return "汪汪"
}

func (d *Dog) Name() string {
	return d.name
}

// 实现接口 - Cat
type Cat struct {
	name string
}

func (c Cat) Speak() string {
	return "喵喵"
}

func (c Cat) Name() string {
	return c.name
}

func main() {
	// 接口变量可以持有任何实现了该接口的类型
	var a Animal

	//a = Dog{name: "旺财"}
	a = &Dog{name: "旺财"} //两种均可以

	//a.SetName("小黑") // 这里会报错，因为接口类型的变量不能调用实现类的非接口方法
	fmt.Println(a.Name(), "说:", a.Speak())
	fmt.Printf("%T\n", a)

	a = Cat{name: "咪咪"}
	fmt.Println(a.Name(), "说:", a.Speak())

	//接口切片
	animals := []Animal{
		&Dog{name: "旺财"},
		Cat{name: "咪咪"},
	}
	for _, animal := range animals {
		fmt.Println(animal.Name(), ":", animal.Speak())
	}

	var e Empty
	e = 1
	value, bool := e.(int)
	if bool {
		fmt.Println("e是int类型，值为：", value)
	} else {
		fmt.Println("e不是int类型")
	}
}
