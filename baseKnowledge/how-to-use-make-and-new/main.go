package main

import (
	"fmt"
)

type Student struct {
	name string
	age  int
}

func main() {
	//structDemo()
	//sliceDemo()
	mapDemo()
}

func mapDemo() {
	map1 := make(map[string]int)
	map1["xiaowang"] = 25
	fmt.Println(map1)
}

func sliceDemo() {
	slice1 := make([]int, 5, 10)
	fmt.Println(slice1)

}

func structDemo() {
	s1 := Student{
		name: "xiaowang",
		age:  12,
	}
	changeStructDemo(&s1)
	fmt.Println(s1)
}

func changeStructDemo(s *Student) {
	s.name = "XIAOWANG"
}
