package main

import "fmt"

/**
1、new
	只分配内存。
	new(T) 对应类型的零值，并且返回它的地址。
2、make
	lice、map、channel的初始化(而不是零值)。
*/

func main() {
	newDemo()
	makeDemo()
}

func newDemo() {
	/**
	new(T) 对应类型的零值，并且返回它的地址
	*/
	num := new(int8)
	fmt.Println(num, *num)
	// 0xc0000ae002 0
}

func makeDemo() {
	userinfo := make(map[string]int)
	fmt.Println(userinfo) // map[]
	userinfo["chenjian"] = 31
	fmt.Println(userinfo) // map[chenjian:31]
}
