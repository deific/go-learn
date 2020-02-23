package main

import "fmt"

/**
管理类
*/
type message struct {
	// 属性
	notifyUser string
	// 定义
	content string
}

/**
用户类
*/
type user struct {
	/* 用户名 */
	name string
	/* 邮件 */
	email      string
	ext        int
	privileged bool
}

/**
管理类
*/
type admin struct {
	// 属性
	person user
	// 定义
	level string
}

/**
管理类
*/
type adminUser struct {
	// 属性
	user
	// 定义
	level string
}

func main() {

	sally := user{"娃哈哈", "as", 123, true}
	//sally.notify()
	sendNotification(&sally)
	sally.showName()

}

//////////////////////////////////////面向接口
// 接口定义
type notifier interface {
	// 方法
	notify()
}

// 接口实现
func (u *user) notify() {
	u.name = "接收者：" + u.name
	fmt.Println("通知：", u.name)
}

// 接口使用
func sendNotification(notifier2 notifier) {
	notifier2.notify()
}

////////////////////////////////////////////////面向对象
func faceObject() {
	sally := user{"娃哈哈", "as", 123, true}
	sally2 := &user{"娃哈哈", "as", 123, true}
	sally.showName()
	fmt.Println(sally.name)

	sally.changeAndShowName()
	fmt.Println(sally.name)

	sally2.changeAndShowName()
	fmt.Println(sally2.name)

	aq := admin{sally, "leve1"}
	aq.person.showName()

	aq1 := adminUser{user: user{"娃哈哈", "as", 123, true}, level: "la"}
	aq1.showName()
	aq1.user.showName()
}

func (u user) showName() {
	fmt.Println(u.name)
}

func (u *user) changeAndShowName() {
	u.name = "我改名了：" + u.name
	fmt.Println(u.name)
}

func object() {
	// 1.实例化
	var bill user
	fmt.Printf("user is:", bill)

	// 2.创建并赋值
	lisa := user{
		name:       "张总",
		email:      "xx@xx.com",
		ext:        123,
		privileged: false,
	}

	fmt.Println(lisa)

	// 3.直接赋值
	lisa2 := user{"李四", "ss@mi.com", 123, true}
	fmt.Println(lisa2)

	// 4.嵌套初始化
	fred := admin{
		person: lisa2,
		level:  "as",
	}

	fmt.Println(fred)
}
