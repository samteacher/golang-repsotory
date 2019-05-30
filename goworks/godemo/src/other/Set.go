package other

import "fmt"

// golang的map是不允许键重复的
func Set()  {

	m := map[string]string{"1": "one","2": "two","3": "three",}
	fmt.Println(m)//输出map[1:one 2:two 3:three]

	// golang的map是不允许键重复的
	//m2 := map[string]string{"1": "one","2": "two","1": "one","3": "three",}
	//fmt.Println(m2)

}
