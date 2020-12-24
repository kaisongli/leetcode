package main

import (
	"fmt"
	"strings"
)

type Stu struct {
	P Person
	School string
}
type Person struct {
	Name string
	Age int
}
func main(){
	//ch := make(chan int)
	//go send(ch)
	//go func() {
	//	defer fmt.Println("子go程结束")
	//
	//	fmt.Println("子go程正在运行……")
	//
	//	ch <- 666 //666发送到c
	//}()
	//num := <- ch
	//fmt.Println(num)

	//var ans []int
	//fmt.Println(ans)
	//fmt.Println(ans == nil)

	dmsServiceName := "dmsServiceName"
	fmt.Println(strings.Join([]string{dmsServiceName, "dmsServiceName"}, ","))

}

func send(ch chan int)  {
	defer fmt.Println("子协程结束....")
	ch <- 555
	fmt.Println("子协程开始....")

}