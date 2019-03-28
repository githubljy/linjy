package main

import "time"
func main(){
	go func() {
		sum:=0
		for i := 0;i< 10 ; i++ {
			sum += i;
	time.Sleep(1*time.Second)
			println(i)
		}

	}()

	for i:=10 ; i < 20 ; i ++{
	time.Sleep(1*time.Second)
		println(i)
	}

}
