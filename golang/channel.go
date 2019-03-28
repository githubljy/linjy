package main

import "runtime"

func main() {

	c:=make(chan struct{})
	ci:=make(chan int,8)
	go func(i chan struct{},j chan int){
		for i:= 0 ; i < 8; i++ {
			ci<-i
		}
		close(ci)
		c<- struct {}{}
		
	}(c,ci)
	println("numgoruntine:",runtime.NumGoroutine())

	<-c
	println("numgoruntine:",runtime.NumGoroutine())
	/*
	for v:= range ci {
		println(v)
	}*/
	
		data:=<-ci
		println(data)
		data =<-ci
		println(data)
		println(<-ci)
		<-ci
		println(<-ci)
	
}
