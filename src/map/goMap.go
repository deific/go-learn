package main

import (
	"fmt"
	"sync"
)

func main() {
	mapDefine()
	//unSafeMap()
	safeMap()
}

func mapDefine() {
	// 内置make
	m := make(map[string]string)

	// 赋值
	index := [...]string{"1", "2", "3", "4", "5", "6", "7", "8"}
	for i := range index {
		fmt.Println("range i:", i, index[i])
		m[index[i]] = index[i] + "A"
	}

	// 遍历
	for key, value := range m {
		fmt.Println("map key and value:", key, value)
	}

	// 删除元素
	delete(m, "2A") // 不存在的key
	fmt.Println("delete after:", m)
	delete(m, "2") // 存在的key
	fmt.Println("delete after:", m)
}

// 并发安全
func unSafeMap() {

	m := make(map[int]int)

	// 新起线程循环写
	go func() {
		for {
			m[1] = 1
		}
	}()

	// 新起线程循环读
	go func() {
		// 不停的读取
		for {
			_ = m[1]
		}
	}()

	for {
		// 无限循环，让程序不退出
		_ = 1
	}
}

// 安全map
func safeMap() {
	// 直接声明定义
	var safeMap sync.Map

	// 循环写
	//go func() {
	//	for {
	//		safeMap.Store(1, 1)
	//	}
	//}()

	safeMap.Store(1, 1)
	// 获取key对于的值，如果不存在则设置值，err=false
	value, err := safeMap.LoadOrStore(2, 3)
	fmt.Println("safeMap ", value, err)

	//go func() {
	//	for {
	//		_,_ = safeMap.Load(1)
	//	}
	//}()

	// 遍历
	safeMap.Range(func(key, value interface{}) bool {
		fmt.Println("safe map:", key, value)
		return true
	})

}
