package other

import (
	"fmt"
	"sort"
)

func MapSort(){

	var m = map[string]int{"hello": 0, "morning": 1, "my": 2, "girl": 3}

	var keys []string

	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	fmt.Println(keys)

	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", m[k])
	}

}
