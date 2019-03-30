package main

import "fmt"

func main(){
	s := "Hello,Summer!"
	fmt.Println(len(s))
	fmt.Println(s[:5])
	fmt.Println(s[6:])
	fmt.Println("嗯 "+s)
}

