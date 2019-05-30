package main

// 初始化时启动，只执行一次
//func init(){
//	fmt.Println("this is Init()...")
//}

import (
	"../other"
	"../utils/strx"
	"fmt"
	"reflect"
)

// 在 main里的main函数也很特殊，它是整个程序执行时的入口
func main() {

	// 基础练习
	//other.Test()

	// main主程序导入其它
	//other.Statement()

	// Map知识点总结
	//other.Map()

	// map排序
	//other.MapSort()

	// 切片的创建
	//other.Slice()

	// 赋值
	//other.Test2()

	// type 类型
	//other.Type()

	// 复数
	//other.Complex()

	//other.NumsToNum(10)

	// 常量
	other.Const()

	v := "hello world"
	fmt.Println(strx.Typeof(v))

	// 使用反射
	fmt.Println("value typeOf : ", reflect.TypeOf(v))

	// 随机数
	//other.RandNumber()

	x := "hello"
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
	}


}
