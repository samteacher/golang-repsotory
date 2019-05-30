package other

import "fmt"

func Map()  {

	//var numbers map[string] int
	//fmt.Println(numbers)

	/**
		声明Map类型的几种方式
	 */

	type personInfo struct {
		ID string
		Name string
		Address string
	}
	// 1.
	rating := map[string] float32 {"C":5, "Go":4.5, "Python":4.5, "C++":2 }
	myMap := map[string] personInfo{"1234": personInfo{"1", "Jack", "Room 101,..."}}
	fmt.Println(rating)
	fmt.Println(myMap)


	// 2.
	numbers2 := make(map[string] int)
	fmt.Println(numbers2)

	// 3.


}
