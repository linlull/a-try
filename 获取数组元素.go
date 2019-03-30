package main

import "fmt"

func main(){
	var a [3]int = [3]int{1, 2, 3}
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	//输出索引和元素
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	//把索引丢掉只输出元素
	for _, v :=range a {
		fmt.Printf("%d\n", v)
	}
}