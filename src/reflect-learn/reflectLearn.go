package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	reflectTypeOf()
	testTypeAndKind()
	testReflectMethod()
	testStructJson()
}

// 通过typeOf获取对象类型
func reflectTypeOf() {
	var a int
	typeOfA := reflect.TypeOf(a)
	fmt.Println(typeOfA.Name(), typeOfA.Kind(), typeOfA.PkgPath())
}

// 定义类型
type Enum int

const (
	ZERO1 Enum = 0
	ZERO2      = 0
)

type cat struct {
}
type dog struct {
	Name  string
	color string
	desc  string `msg:"这是一个描述"`
}

// 测试type和kind
func testTypeAndKind() {

	typeOfCat := reflect.TypeOf(cat{})
	fmt.Println("typeOfCat :", typeOfCat.Name(), typeOfCat.Kind(), typeOfCat.PkgPath())

	typeOfZero1 := reflect.TypeOf(ZERO1)
	fmt.Println("typeOfZero1 :", typeOfZero1.Name(), typeOfZero1.Kind(), typeOfZero1.PkgPath())

	typeOfZero2 := reflect.TypeOf(ZERO2)
	fmt.Println("typeOfZero2 :", typeOfZero2.Name(), typeOfZero2.Kind(), typeOfZero2.PkgPath())

	// 指针变量
	cat := &cat{}
	// 获取到的时候指针的种类，指针无类型
	typeOfCat2 := reflect.TypeOf(cat)
	fmt.Println("typeOfCat2 & :", typeOfCat2.Name(), typeOfCat2.Kind(), typeOfCat2.PkgPath())
	// 获取原始类型和种类
	typeOfCat2 = typeOfCat2.Elem()
	fmt.Println("typeOfCat2 *:", typeOfCat2.Name(), typeOfCat2.Kind(), typeOfCat2.PkgPath())

	dog1 := dog{
		Name:  "wangwang",
		color: "red",
		desc:  "aa",
	}

	typeOfDog := reflect.TypeOf(dog1)
	for i := 0; i < typeOfDog.NumField(); i++ {
		field := typeOfDog.Field(i)
		fmt.Println("typeOfDog field :", field.Name, field.Type, field.Tag)
	}

	// 反射获取值
	valueOfDog := reflect.ValueOf(&dog1)
	fmt.Println("valueOfDog :", valueOfDog, valueOfDog.Type(), valueOfDog.IsValid())
	// 修改值
	valueOfDog = valueOfDog.Elem()
	valueOfDog.Field(0).SetString("bbbb")
	fmt.Println("valueOfDog :", valueOfDog, valueOfDog.Type(), valueOfDog.IsValid())

	// 反射创建实例
	// 先根据反射获取类型，再通过new实例化
	dogType := reflect.TypeOf(dog1)
	dogTypeIns := reflect.New(dogType)

	fmt.Println("dogTypeIns :", dogTypeIns, dogTypeIns.Type(), dogTypeIns.IsValid())
}

func add(a, b int) int {
	return a + b
}

// 测试反射调用
func testReflectMethod() {
	// 反射函数
	funcValue := reflect.ValueOf(add)
	// 使用Value切片构造函数参数
	paramList := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(29)}
	// 通过Call方法调用函数
	retList := funcValue.Call(paramList)

	fmt.Println("testReflectMethod:", retList[0].Int())
}

// 测试结构体和JSON
func testStructJson() {
	// 技能
	type Skill struct {
		Name  string
		Level int
	}
	// 角色
	type Actor struct {
		Name   string
		Age    int
		Skills []Skill
	}

	// 初始化数据
	a := Actor{
		Name: "cow boy",
		Age:  37,
		Skills: []Skill{
			{Name: "Roll and roll", Level: 1},
			{Name: "Flash your dog eye", Level: 2},
			{Name: "Time to have lunch", Level: 3},
		},
	}

	// 转换为json
	if aJson, err := MarshalJson(a); err == nil {
		fmt.Println("json:", aJson)
	} else {
		fmt.Println("error:", err)
	}
}

// 转换为json字符串
func MarshalJson(v interface{}) (string, error) {
	var buffer bytes.Buffer
	// 将任意值转为json并输出为json格式
	if err := toJson(&buffer, reflect.ValueOf(v)); err == nil {
		return buffer.String(), nil
	} else {
		return "", err
	}
}

// 使用反射机制转为json格式
func toJson(buffer *bytes.Buffer, value reflect.Value) error {
	// 根据反射的值，判断种类和类型
	switch value.Kind() {
	case reflect.String:
		buffer.WriteString(strconv.Quote(value.String()))
	case reflect.Int:
		buffer.WriteString(strconv.FormatInt(value.Int(), 10))
	case reflect.Slice:
		return writeSlice(buffer, value)
	case reflect.Struct:
		return writeStruct(buffer, value)
	default:
		buffer.WriteString(strconv.Quote(value.String()))
	}
	return nil
}

// 转换切片
func writeSlice(buffer *bytes.Buffer, value reflect.Value) error {
	// 切片数据开始
	buffer.WriteString("[")
	// 遍历切片元素
	for s := 0; s < value.Len(); s++ {
		sliceValue := value.Index(s)
		toJson(buffer, sliceValue)
		// 每个元素写一个逗号，最后一个元素不需要
		if s < value.Len()-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("]")
	return nil
}

// 转换结构体
func writeStruct(buffer *bytes.Buffer, value reflect.Value) error {
	// 获取值的对象类型
	valueType := value.Type()
	fmt.Println("valueType:", valueType)
	// 写入结构体的大括号
	buffer.WriteString("{")

	// 遍历结构体的所有值
	for i := 0; i < value.NumField(); i++ {
		// 获取每个字段的值
		fieldValue := value.Field(i)
		// 获取每个字段的类型
		fieldType := valueType.Field(i)

		// 写入每个字段的双引号
		buffer.WriteString("\"")
		// 写入每个字段的名称
		buffer.WriteString(fieldType.Name)
		// 写入每个字段的双引号
		buffer.WriteString("\":")
		toJson(buffer, fieldValue)
		// 每个元素写一个逗号，最后一个元素不需要
		if i < value.NumField()-1 {
			buffer.WriteString(",")
		}
	}
	// 写入结构体的大括号
	buffer.WriteString("}")
	return nil
}
