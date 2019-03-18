//接口的内部表现 一个接口可以被认为是由一个元组（类型，值）在内部表示的。type是接口的基础具体类型，value是具体类型的值
package main

import (
	"fmt"
)

type Test interface {
	Tester()
}

type MyFloat float64 

func (m MyFloat) Tester() {
	fmt.Println(m)
}

func describe(t Test) {
	fmt.Printf("Interface type: %T , value:%v \n",t,t)
}

func main(){

	var t Test
	f := MyFloat(89.777)
	t = f
	describe(t)
	t.Tester()

}
