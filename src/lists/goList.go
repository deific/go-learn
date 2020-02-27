package main

import (
	list "container/list"
	"fmt"
)

func main() {
	listDefine()
}

func listDefine() {
	list := list.New() // 与list:= list.List{}等价

	// 列表尾部加入元素
	index := list.PushBack(1)
	list.PushBack("2")
	// 中间插入
	list.InsertAfter("1.5", index)

	fmt.Println("list length is :", list.Len())

	for i := list.Front(); i != nil; i = i.Next() {
		fmt.Println("list element:", i.Value)
	}
}
