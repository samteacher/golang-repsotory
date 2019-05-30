package other

type Nums int

func NumsToNum(nums Nums) int {

	// 类型转换
	// 在必要以及可行的情况下，一个类型的值可以被转换成另一种类型的值。由于 Go 语言不存在隐式类型转
	// 换，因此所有的转换都必须显式说明，就像调用一个函数一样

	return int(nums)

}
