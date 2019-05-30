package other

import "fmt"

func Slice() {

	var array []int = []int{1, 7, 18, 34, 88, 90, 150}
	sourceSlice := []int{1, 7, 18, 34, 88, 90, 150}

	// 切片的五种定义方式：
	var slice []int
	slice2 := *new([]int)
	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := make([]int, 5, 10)
	slice5 := array[1:5]
	slice6 := sourceSlice[1:5]

	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(slice5)
	fmt.Println(slice6)

}
