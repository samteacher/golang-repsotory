package other

func Test(){

	//fmt.Println("Hello world！！！")
	//fmt.Printf("世界，你好！")

	//if len(os.Args)!=0{
	//	// args 第一个片 是文件路径
	//	fmt.Println(os.Args[0])
	//}
	//// 第二个参数是， 用户输入的参数 例如 go run osdemo01.go 123
	//fmt.Println(os.Args[1])

	//组成程序的函数、变量、常量、类型的声明语句（分别由关键字 func,var,const,type定义）。这些内容的声明顺序并不重要

	// os.Args变量是一个字符串（string）的切片（slice）（译注：slice和Python语言中的切片类似，
	// 是一个简版的动态数组），切片是Go语言的基础概念，稍后详细介绍。现在先把切片s当	作数组元素序列,
	// 序列的成长度动态变化, 用 s[i]  访问单个元素，用 s[m:n]  获取子序列(译注：和python里的语法差不多)。
	// 序列的元素数目为len(s)。和大多数编程语言类似，区间索引时，Go言里也采用左闭右开形式, 即，区间包括第一个索引元素，
	// 不包括最后一个, 因为这样可以简化逻辑。（译注：比如a = [1, 2, 3, 4, 5], a[0:3] = [1, 2, 3]，不包含最后一个元素）。
	// 比如s[m:n]这个切片，0 ≤ m ≤ n ≤ len(s)，包含n-m个元素

	//var s, sep string
	//// 外部入参使用os.Args访问该参数
	//for i := 1; i < len(os.Args); i++ {
	//	s += sep + os.Args[i]
	//	sep = " "
	//}
	//fmt.Printf("%s -- %d number\n", "os.Args的个数是：", len(os.Args)-1)


	//Go语言只有for循环这一种循环语句。for循环有多种形式，其中一种如下所示：

	//for initialization; condition; post {
	//	// zero or more statements
	//}
	//for循环三个部分不需括号包围。大括号强制要求, 左大括号必须和post语句在同一行

	//initialization语句是可选的，在循环开始前执行。initalization如果存在，必须是一条简单语句
	//（simple statement），即，短变量声明、自增语句、赋值语句或函数调用。 condition  是一
	//个布尔表达式（boolean expression），其值在每次循环迭代开始时计算。如果为 true  则执
	//行循环体语句。 post  语句在循环体执行结束后执行，之后再次对 conditon求值。 condition  值为 false  时，循环结束。
	//for循环的这三个部分每个都可以省略，如果省略 initialization  和 post  ，分号也可以省略：
	//for condition {
	//	// ...
	//}

	//如果连 condition  也省略了，像下面这样：
	//// a traditional infinite loop
	//for {
	//	// ...
	//}

	//这就变成一个无限循环，尽管如此，还可以用其他方式终止循环, 如一条 break 或 return语句。


	// for循环的另一种形式, 在某种数据类型的区间（range）上遍历
	//s, sep := "", ""
	// range 遍历
	//for _, arg := range os.Args[:4] {
	//	s += sep + arg
	//	sep = " "
	//}
	//fmt.Println(s)


	// 声明一个变量有好几种方式，下面这些都等价：

	//// 短变量声明，最简洁，但只能用在函数内部,而不能用于包变量
	//str := ""
	//// 依赖于字符串的默认初始化零值机制，被初始化为""
	//var str1 = ""
	//// 用得很少,除非同时声明多个变量
	//var str2 string
	//// 显式地标明变量的类型，当变量类型与初值类型相同时，类型冗余，但如果两者类型不同，变量类型就必须了
	//var str3 string = ""
	//
	//
	//fmt.Println(str)
	//fmt.Println(str1)
	//fmt.Println(str2)
	//fmt.Println(str3)

	//fmt.Println(strings.Join(os.Args[1:], "-- "))


	// 变量声明的一般语法如下：
	// var 变量名字 类型 = 表达式

	//其中“类型”或“= 表达式”两个部分可以省略其中的一个。如果省略的是类型信息，那么将根据
	//初始化表达式来推导变量的类型信息。如果初始化表达式被省略，那么将用零值初始化该变
	//量。 数值类型变量对应的零值是0，布尔类型变量对应的零值是false，字符串类型对应的零
	//值是空字符串，接口或引用类型（包括slice、map、chan和函数）变量对应的零值是nil。数
	//组或结构体等聚合类型对应的零值是每个元素或字段都是对应该类型的零值

	// var num int
	// var str string
	// var foo float32
	// var maps map[int] int
	//
	// fmt.Println(num)
	// fmt.Println(str)
	// fmt.Println(foo)
	// fmt.Println(maps)

}
