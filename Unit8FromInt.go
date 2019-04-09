package main

import (
	"fmt"
	"math"
)

func Uint8FromInt(x int) (uint8, error){
	if 0 <=x && x <= math.MaxUint8 {
		return uint8(x), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", x)
}
func main(){
	var x int = 1
	b, err := Uint8FromInt(x)
	fmt.Print(b,err)
}

