//编写一个就地处理函数，用于去除[]string slice中相邻的重复元素
package main

import "fmt"

func main() {
	s := []string{"h", "e", "l", "l", "o"}

	for i := 0; i<len(s)-1; i++  {
		if s[i] == s[i+1]{
			fmt.Println(remove(s, i))
		}
	}
}
func remove(slice []string , i int) []string {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}