package main

import (
	"fmt"
	"os"
	"testing"
)

/**
1、普通正常测试
	- `go test`
	- `go test -v` # -v 参数会显示每个用例的测试结果，另外 -cover 参数可以查看覆盖率。
	- `go test -run TestAdd -v` # 如果只想运行其中的一个用例。
PS：
	Errorf 遇错不停，还会继续执行其他的测试用例
*/
func TestAdd(t *testing.T) {
	if ans := Add(1, 2); ans != 3 {
		t.Errorf("1 + 2 expected be 3, but %d got", ans)
	}
	if ans := Add(-10, -20); ans != -30 {
		t.Errorf("-10 + -20 expected be -30, but %d got", ans)
	}
}

/**
2、子测试(Subtests)
	- $ go test -run TestMul/zhengshu -v
PS：
	Fatal 遇错即停
*/
func TestMul(t *testing.T) {
	t.Run("zhengshu", func(t *testing.T) {
		if Mul(2, 3) != 6 {
			t.Fatal("fail")
		}
	})
	t.Run("fushu", func(t *testing.T) {
		if Mul(-2, -3) != 6 {
			t.Fatal("fail")
		}
	})
}

/**
3、多个子测试的场景，更推荐如下的写法(table-driven tests)
*/
func TestSub(t *testing.T) {
	cases := []struct {
		Name           string
		A, B, Expected int
	}{
		{"pos", 2, 3, -1},
		{"neg", 2, -3, 5},
		{"zero", 2, 0, 2},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := Sub(c.A, c.B); ans != c.Expected {
				t.Fatalf("%d * %d expected %d, but %d got",
					c.A, c.B, c.Expected, ans)
			}
		})
	}
}

/**
4、帮助函数(helpers)
	对一些重复的逻辑，抽取出来作为公共的帮助函数(helpers)，
	可以增加测试代码的可读性和可维护性。
	借助帮助函数，可以让测试用例的主逻辑看起来更清晰。

t.Helper()
	Go 语言在 1.9 版本中引入了 t.Helper()，
	用于标注该函数是帮助函数，报错时将输出帮助函数调用者的信息，而不是帮助函数的内部信息。
*/
type calcCase struct {
	A, Expected int
}

func createDoubleTestCase(t *testing.T, c *calcCase) {
	t.Helper()
	if ans := Double(c.A); ans != c.Expected {
		t.Fatalf("2 * %d expected %d, but %d got",
			c.A, c.Expected, ans)
	}
}
func TestDouble(t *testing.T) {
	createDoubleTestCase(t, &calcCase{2, 4})
	createDoubleTestCase(t, &calcCase{4, 8})
	createDoubleTestCase(t, &calcCase{100, 200})
}

/**
5、准备(setup)和回收(teardown)工作
*/
func setup() {
	fmt.Println("Before all tests")
}

func teardown() {
	fmt.Println("After all tests")
}

func Test1(t *testing.T) {
	fmt.Println("I'm test1")
}

func Test2(t *testing.T) {
	fmt.Println("I'm test2")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
