package main

import (
	"fmt"

)

func main(){
	a := "linlu"
	b := "lin"
	c := "lu"
	fmt.Printf("b是否为a的前缀 %t\n", HasPrefix(a, b))
	fmt.Printf("b是否为a的后缀 %t\n", HasSuffix(a, c))
	fmt.Printf("b是否为a的子串 %t\n",Contains(a, c))

}
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}
func HasSuffix(s, suffix string) bool {
	return len(s) >=  len(suffix) && s[len(s) - len(suffix):] == suffix
}
func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++{
		if HasPrefix(s[i:], substr){
			return true
		}
	}
	return false
}