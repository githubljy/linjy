package main

import(
"fmt"
)

func assert(i interface{}){
    s,ok := i.(int)
    fmt.Println(s,ok)
}

func main(){
  var s interface{} = 5.4
  assert(s)
}
