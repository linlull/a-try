package main

import (
	"fmt"
	"sort"
)

func main(){
	ages := map[string]int{
		"zlice":   31,
		"linlu":   20,
		"charlie": 34,
        "bob": 0,
	}
	//用string函数进行键的排序
		var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
	//辨别map中一个元素是不存在还是其本身是零
	age, ok := ages["bob"]
	if !ok{
		fmt.Printf("bob不在map中,age=%d", age)
	}
	//判断两个map是否有同样的键和值
	ages2 := map[string]int{
		"zlice":   31,
		"linlu":   20,
		"charlie": 34,
		"bob": 0,
	}
		fmt.Println(equal(ages, ages2))
}
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x{
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true

}
