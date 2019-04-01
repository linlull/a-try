package main

import "fmt"


//给基础类型添加方法
type long int
func (a long) Add2(b long) long{
	return a +b
}
func main(){
	var a long = 2

	fmt.Println(a.Add2(3))
}
