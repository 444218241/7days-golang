package main

import (
	"fmt"
)

type Bird interface {
	fly()
	sing()
}

type Sparrow struct {
	name string
	age  int
}

func (s Sparrow) fly() {
	fmt.Println("I am flying.")
}

func (s Sparrow) sing() {
	fmt.Println("I can sing.")
}

type Parrot struct {
	name string
	age  int
	kind int
}

func (p Parrot) fly() {
	fmt.Println("I am flying.")
}
func (p Parrot) sing() {
	fmt.Println("I can sing.")
}

func main() {
	//staticAndDynamicTypeDemo()
	assertDemo()
}

func assertDemo() {
	/**
	由于「接口类型」的变量的「动态类型」是变化的，
	类型断言：对「接口类型」的变量进行类型检查。
	*/
	var b Bird // // Bird interface
	b = Sparrow{}
	b = Parrot{}

	value, ok := b.(Parrot)
	fmt.Println(value, ok)
	// { 0 0} true
	// 断言成功：value表示断言成功之后目标类型变量；
}

func staticAndDynamicTypeDemo() {
	/**
	在使用 fmt.Printf("%T\n") 获取一个变量的类型时，其实是调用了reflect包的方法进行获取的，
	reflect.TypeOf 获取的是接口变量的「动态类型」，
	reflect.valueOf() 获取的是接口变量的「动态值」。
	PS: reflect.TypeOf 函数定义 func TypeOf(i interface{}) Type{}
	*/

	// 1、静态类型
	var age int
	fmt.Printf("%T\n", age) // int
	/**
	1。 age 调用 TypeOf 时，会进行类型的转换，将int型变量age转换为 interface 型，
	2。 在这个过程中会将变量 age 的类型（int）作为 调用 的动态类型，
	3。 age 的值（在这里是age 的零值0）作为 调用 的动态值。
	4。 因为 TypeOf() 获取的是age 的动态类型，所以这个时候展示出的类型为 int。
	*/

	// 2、动态类型
	var b Bird            // Bird interface
	fmt.Printf("%T\n", b) // <nil>
	/**
	因为一个「接口类型变量」在没有被赋值之前，它的「动态类型」和「动态值」都是 nil 。
	*/

	b = Sparrow{}         // Sparrow struct 实现了Bird interface
	fmt.Printf("%T\n", b) // main.Sparrow

	b = Parrot{}          // Parrot struct 实现了Bird interface
	fmt.Printf("%T\n", b) // main.Parrot
}
