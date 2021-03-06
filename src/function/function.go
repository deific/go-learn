package main

import (
	"fmt"
	"runtime"
	"strings"
)

func main() {
	a, b := testAdd(1, 2)
	fmt.Println("a + b=", a)
	fmt.Println("a * b=", b)

	// 函数也可以是一种类型
	var mathA func(a, b int) (a1, b1 int)
	mathA = testAdd
	mathA(1, 2)

	// 匿名函数作为回调
	testStrChain(func(str []string) {
		fmt.Println("testStrChain ", str)
	})

	// 结构体实现接口
	testStructInterface()
	// 方法实现接口
	testFunctionInterface()

	// 闭包
	testClosure()

	// panic 和 recover
	testRecoverFunction()

}

// 声明有返回值及返回名称
func testAdd(a, b int) (addRet, mulRet int) {
	addRet = a + b
	mulRet = a * b
	mulRet++
	return
}

// 字符串链式处理
func StrProcess(list []string, chain []func(string) string) {
	// 遍历每一个字符
	for index, str := range list {
		// 第一个需要处理的字符
		result := str

		// 调用处理链处理
		for _, proc := range chain {
			result = proc(result)
		}
		list[index] = result
	}
}

// 自定义字符串处理，删除前缀
func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

// 测试字符串链式处理
func testStrChain(callback func([]string)) {
	list := []string{
		"go world1",
		"go world2",
		"go world3",
		"go world4",
	}

	chain := []func(string) string{
		removePrefix,
		strings.ToLower,
	}

	StrProcess(list, chain)

	// 匿名函数,赋值给一个变量
	//var print func(list []string)
	print := func(list []string) {
		fmt.Println("str chain :", list)
	}
	print(list)
	callback(list)
}

// 函数实现接口
// 定义接口
type Invoker interface {
	Call(interface{})
}

// 定义结构体并实现接口
type Simple struct {
	result string
}

// 接口实现方法
func (s *Simple) Call(str interface{}) {
	fmt.Println(str)
	s.result = " function interface"
}

// 结构体实现接口
func testStructInterface() {
	// 定义接口变量
	var invoker Invoker
	// 实例化结构体
	//s := Simple{result:""}
	s := new(Simple)

	// 将实例化后的结构体赋值到接口变量上
	invoker = s
	// 使用接口变量调用结构体Simple的实现
	// 类似java中使用接口类型引用实现实例
	invoker.Call("hello ")
}

// 函数实现接口

// 定义函数类型
type FuncCaller func(interface{})

func (f FuncCaller) Call(p interface{}) {
	f(p)
}

// 函数的声明不能直接实现接口，需要将函数定义为类型后，使用类型实现结构体。当类型方法被调用时，还需要调用函数本体。
func testFunctionInterface() {
	// 定义接口变量
	var invoker Invoker
	invoker = FuncCaller(func(p interface{}) {
		fmt.Println("testFunctionInterface", p)
	})
	invoker.Call("hello")
}

// 匿名函数可以用于闭包
func Accmulate(value int) func() int {
	return func() int {
		value++
		return value
	}
}

func testClosure() {
	accmulator := Accmulate(1)
	// 累计加一
	fmt.Println(accmulator())
	fmt.Println(accmulator())
	fmt.Println(accmulator())
	fmt.Printf("%p\n", accmulator)

	accmulator2 := Accmulate(20)
	fmt.Println(accmulator())
	fmt.Println(accmulator2())
	fmt.Printf("%p\n", accmulator2)
}

// panic和recover
type panicContext struct {
	function string
}

// 保护函数
func protectMan(entry func()) {

	defer func() {
		// 发生宕机时，获取panic上下文并打印
		err := recover()
		switch err.(type) {
		case runtime.Error:
			fmt.Println("runtime error:", err)
		default:
			fmt.Println("error:", err)
		}
	}()

	entry()
}

// 异常恢复
func testRecoverFunction() {
	fmt.Println("被保护程序，如果发生宕机会由保护程序恢复程序，类似try-catch机制")
	fmt.Println("运行前")
	protectMan(func() {
		// 手动宕机异常
		fmt.Println("手动宕机前")
		panic(&panicContext{"手动触发宕机"})
		fmt.Println("手动宕机后")
	})

	protectMan(func() {
		fmt.Println("空指针赋值前")
		var a *int
		*a = 1
		fmt.Println("空指针赋值后")
	})
	fmt.Println("运行后")
}
