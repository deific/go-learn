package main

import "fmt"

func main() {
	arrayType()
}

// 数组和切片
func arrayType() {

	// 数组长度固定，不可改变长度，可以改变成员值
	// 切片长度是动态的

	// 数组的声明方式：必须指定数组长度
	// 初始化使用类型的默认值初始化
	var strList = [...]string{"1", "2"}
	// 数组长度不可改变
	//strList = append(strList, "aa")
	fmt.Println("strList = ", strList)

	// 切片是都动态长度，声明时不指定长度
	var str2List []string
	str2List = append(str2List, strList[0])
	// append不能添加数组
	//str2List = append(str2List, strList...)
	// 转换成切片可以append
	str2List = append(str2List, strList[:]...)
	fmt.Println("str2List = ", str2List)
	str2List = append(str2List, "asddd")
	fmt.Println("str2List = ", str2List)

	// make创建长度不为零的切片
	var str3List = make([]string, 3)
	str3List[0] = "A3"
	str3List[1] = "B3"
	str3List[2] = "C3"
	str3List = append(str3List, "a3", "b3", "c3")
	fmt.Println("str3List = ", str3List, len(str3List))
	// make创建时，除了可以指定初始长度外，还可以指定每次需要扩展长度时，分配多少空间
	var str4List = make([]string, 1, 3)
	fmt.Println("str4List = ", str4List, len(str4List))
	// append方法可以添加切片
	str4List = append(str4List, str3List...)
	// 实际已经预分配了空间，但不影响size
	fmt.Println("str4List = ", str4List, len(str4List))

	// copy复制切片,必须分配且有足够的空间，否则空间不足只能复制空间所能容纳的元素
	var str5List = make([]string, 10)
	//var str5List = []string{""}
	str5List[0] = "1"
	str5List = append(str5List, "2")
	fmt.Println("str5List = ", str5List)
	count := copy(str5List, str4List)
	fmt.Println("str5List = ", str5List, count)

	// 删除切片元素
	var str6List = make([]string, 10)
	copy(str6List, str5List)
	fmt.Println("str6List = ", str6List)
	// 删除C3
	str6List = append(str6List[0:3], str6List[5:]...)
	fmt.Println("str6List = ", str6List)
}
