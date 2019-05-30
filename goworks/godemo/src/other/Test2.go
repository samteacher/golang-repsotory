package other

import (
	"fmt"
	"os"
)

// 如果表达式太复杂的话，应该尽量避免过度使用元组赋值；因为每个变量单独赋值语句的写法可读性会更好

func Test2() {

	file, err := os.Open("D://foo.txt")
	if err != nil {
		fmt.Println("Open file Failed：", err)
		return
	}
	defer func() {
		file.Close()
	}()

	var b []byte = make([]byte, 4096)
	n, err := file.Read(b)

	if err != nil {
		fmt.Println("Open file Failed", err)
	}

	data := string(b[:n])
	fmt.Println(data)


	// 无组赋值
	i, j, k := 2, 3, 5
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

	fmt.Println(addNum(10, 2))
}

func addNum(x, y int) int {
	// 无组赋值 计算
	y, x = x*y, y-x
	return x
}
