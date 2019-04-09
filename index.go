package main

import (
	"fmt"
	"strings"
)

func main(){
	line := "Today is Tuesday"
	i := strings.Index(line," ")
	firstword := line[:i]

	j := strings.LastIndex(line, " ")
	lastword := line[j+1:]

	fmt.Println(firstword,lastword)
}
