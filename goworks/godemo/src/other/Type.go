package other

import "fmt"

type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

func Type() {

	CToF(AbsoluteZeroC)

	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0) // "true"
	fmt.Println(f >= 0) // "true"
	//fmt.Println(c == f) // compile error: type mismatch
	fmt.Println(c == Celsius(f)) // "true"!

	// 类型别名
	// 当你在使用某个类型时，你可以给它起另一个名字，然后你就可以在你的代码中使用新的名字（用于简化名称或解决名称冲突）。
	// 在typeTZint中，TZ就是int类型的新名称（用于表示程序中的时区），然后就可以使用TZ来操作int类型的数据
	type TZ int
	var a, b TZ = 3, 4
	d := a + b
	fmt.Printf("c has the value: %d", d) // 输出：c has the value: 7

}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
