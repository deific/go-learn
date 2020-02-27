package main

import "fmt"

func main() {
	ifExpress()

	forExpress()

	forRangeExpress()

	switchExpress()

	gotoExpress()
}

// if 控制语句
func ifExpress() {

	isTrue := true
	// { 必须与if在一行，否则报错
	if isTrue {
		fmt.Println("if ", isTrue)
	} else if isTrue == false {
		fmt.Println("if ", isTrue)
	} else {
		fmt.Println("if ", isTrue)
	}

	var checkTrue = checkResult

	// 新起一个线程
	go func() {
		_ = checkTrue
	}()

	// 先执行表达式，再用if判断
	if isOk, value := checkTrue(); isOk {
		fmt.Println("first execute expression, after check if: ", isOk, value)
	}
}

func checkResult() (bool, int) {
	return true, 1
}

// for循环控制语句
func forExpress() {

	step := 4

	for ; step > 0; step-- {

		if step == 2 {
			goto doSomething
		}
		fmt.Println("for item is :", step)
	}

	// 标签语句块
doSomething:
	fmt.Println("here is 9X9：")
	for x := 1; x <= 9; x++ {
		for y := 1; y <= x; y++ {
			fmt.Printf("%d*%d=%d ", y, x, x*y)
		}
		fmt.Println()
	}
}

// for range 表达式
func forRangeExpress() {

	// 遍历数组
	for index, value := range [...]int{1, 2, 3, 4, 5} {
		fmt.Println("array index and value:", index, value)
	}

	// 遍历切片
	slices := []int{1, 2, 3, 4, 5}
	slices = append(slices, 6)
	for index, value := range slices {
		fmt.Println("slice index and value:", index, value)
	}

	// 遍历字符串
	str := "AAAAbbb乙"
	for index, value := range str {
		fmt.Println("character is:", index, value)
	}

	// 遍历map
	strMap := make(map[string]string)
	strMap["a"] = "甲"
	strMap["b"] = "乙"
	strMap["c"] = "丙"

	intMap := map[int]string{
		1: "我加",
		2: "住在",
	}

	for key, value := range intMap {
		fmt.Println("map item is:", key, value)
	}

	// 遍历通道消息
	c := make(chan string)

	go func() {
		c <- "下次1"
		c <- "下次2"
		c <- "下次3"
		c <- "下次4"

		// 关闭通道
		close(c)
	}()

	for m := range c {
		fmt.Println("channel msg is:", m)
	}

}

func switchExpress() {

	a := "hell0"

	// switch 每次只能选择一个分支执行，支持多值选择
	switch a {
	case "hello", "hell0":
		fmt.Println("this is hello")
	case "world":
		fmt.Println("this is world")
	default:
		fmt.Println("this is default")
	}

}

// goto 语句
func gotoExpress() {

	// 直接跳出所有for循环
	for x := 1; x < 100; x++ {
		for y := 1; y < 100; y++ {
			if x*y > 1000 {
				goto ok
			}
		}
	}

ok:
	fmt.Println("we are finish")
}
