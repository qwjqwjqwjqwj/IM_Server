package main

import (
	"fmt"
)

type myint int

// 父类
type Book struct {
	name  string
	price int
}

func (this Book) GetName() string {
	return this.name
}
func (this *Book) SetName(newName string) {
	this.name = newName
}

// 结构体嵌入
type MagicBook struct {
	Book //匿名嵌入 实例可直接使用Book的方法和属性
	//b Book //有名嵌入 假如说为b Book那么mb1 := MagicBook{name:"xxx"}`，必须写mb1 := MagicBook{b: Book{name:""}}
	power int
}

func (this MagicBook) GetName() {
	fmt.Printf("     name of this MagicBook's name is %s\n", this.name)
}

func main() {

	var a myint = 10
	fmt.Printf("a =%d,  type of a is %T \n", a, a)

	var b1 Book
	b1.name = "《Go语言开发基础》"
	b1.price = 100
	b1 = Book{name: "《Go语言开发基础》", price: 100}
	fmt.Println(b1.GetName())

	fmt.Printf("b1's name is %s\n", b1.GetName())
	fmt.Printf("b1 = %#v   location of b1 it %p\n", b1, &b1)

	var b2 Book
	b2 = b1 //值传递，b2是b1的副本，b2和b1是两个不同的对象
	b2.SetName("《数字信号处理》")
	b2.price = 200
	fmt.Printf("b2 = %#v   location of b2 it %p\n", b2, &b2)

	// var b3 Book
	// b3.SetName(b2.GetName())
	// fmt.Printf("b3's name is %s\n", b3.GetName())

	mb1 := MagicBook{name: "《计算机网络》", price: 300, power: 1000}

	fmt.Printf("mb1 = %#v   location of mb1 it %p\n", mb1, &mb1)
	mb1.GetName()
}
