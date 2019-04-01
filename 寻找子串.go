package main

import (
	"errors"
	"fmt"
)

var e = errors.New("cannot find")
func main(){
	stra := "linlull"
	strb := "lu"
	fmt.Println(find(stra, strb))
}
func find(a, b string) (position int, err error){
	for i:= 0; i < len(a); i++ {

		if b[0] == a[i] && a[i:i+len(b)] == b[:]{
			return i, err
		}
	}
	return -1, e
}