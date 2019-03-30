package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){
	s := "Hello,世界"
	for i:= 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

}


