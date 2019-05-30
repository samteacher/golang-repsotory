// 本人自己封闭的一些常用的工具及方法

package strx

import (
	"fmt"
	"reflect"
)

// 如果某个函数的入参是interface{}，有下面几种方式可以获取入参的方法：
// 1.fmt
func Typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

// 2.反射
func Typeof2(v interface{}) string {
	return reflect.TypeOf(v).String()
}

// 3.类型断言
func Typeof3(v interface{}) string {
	switch t := v.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	//... etc
	default:
		_ = t
		return "unknown"
	}
}
