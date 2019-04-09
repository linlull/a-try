package main

import (
	"fmt"
	"strings"
)

type BitFlag int
const (
	Active BitFlag = 1<< iota
	Send // 1<< 1
	Receive //1<< 2
)


  func (flag BitFlag) String() string {
  	var flags []string
  	if flag & Active ==Active{
  		flags = append(flags, "Active")
	}
  	if flag & Send == Send {
  		flags = append(flags, "Send")
	}
  	if flag & Receive == Receive {
  		flags = append(flags, "Receive")
	}
  	if len(flags) > 0 {
  		return fmt.Sprintf("%d(%s)",int(flag), strings.Join(flags,"|"))
	}
  	return "0"
  }
func main(){
	var flag BitFlag = 0
	fmt.Println(flag.String())
}