//空接口
package main

import (
	"fmt"
)

func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {
     // 任何类型的变量传入都可以

	s := "Hello World"
	i := 55
	strt := struct {
		name string
	}{
		name: "Naveen R",
	}
	describe(s)
	describe(i)
	describe(strt)
}
