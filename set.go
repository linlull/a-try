//2.自己实现的集合
//    利用Go实现一个集合，并且具有如下操作：
//    为集合增加一个元素：```func (set *Set) AddItem(item Type) (err error)```
//    查找集合中是否包含一个元素：```func (set *Set) Contain(item Type) (exist bool)```
//    获取集合中的元素的个数：```func (set *Set) Size() (size int)```
//    从集合中删除一个元素：```func (set *Set) Delete(item Type) (err error)```
//    返回一个包含所有元素的数组切片：```func (set *Set) Items() (sli []Type, err error)```
//    注意：以上函数中的Set和Type需要自己定义（使用type关键字
//    提示：可以使用map[Type]bool来进行实现
// 进阶：
//    实现两个集合的交集，并集，补集，差集
package main

import (
	"errors"
	"fmt"
)

type Type int
type Set map[Type]bool //映射的键不能重复，利用这个定义集合

func (set *Set) AddItem(item Type) error{
	_, ok := (*set)[item]//先判断元素是否存在
	if ok {
		return errors.New("already exists")
	} else {
		(*set)[item] = true
	}
	return nil
}


func (set *Set) Contain(item Type) bool{
	_, ok := (*set)[item]
	return ok
}

func (set *Set) Size() int{
	return len(*set)
}

func (set *Set) Delete(item Type) error{
	_, ok := (*set)[item]//把迭代出来的不需要的值赋给空变量
	if ok {
		delete(*set, item)
	}else{
		return errors.New("can't find it")
	}
	return nil
}

func (set *Set) Items() ([]Type, error){
	res := []Type{} //一个空的数组切片
	for key, _ := range(*set){
		res = append(res, key)

	}
	return res, nil
}

func And(setA, setB Set) Set{
	set := make(Set)
	for k, _ := range(setA){
		if _, ok := setB[k]; ok{
			set[k] = true//如果元素a在b中存在，则将元素放入set
		}
	}
	return set
}

func Or(setA, setB Set) Set{
	set := make(Set)
	for k,_ := range setA {
		set[k] = true}
		for k,_ := range setB{
		if _, ok := setA[k]; !ok{
			set[k] =true//如果元素a在b中不存在，则将元素放入set
		}
	}
	return set
}

func (setA *Set)Sub(setB Set) Set{
	set := *setA
	for k, _ := range setB {
	if _, ok := (*setA)[k]; ok{
		set.Delete(k)//如果元素在a中存在就删除该元素
	}
	}
    return set
}

func test() {
	seta := make(Set)
	setb := make(Set)
	seta.AddItem(1)
	seta.AddItem(2)
	seta.AddItem(5)
	setb.AddItem(2)
	setb.AddItem(11)
	setb.AddItem(7)
	fmt.Println("seta:", seta)
	fmt.Println("setb:", setb)
	fmt.Println("seta and setb:", And(seta, setb))
	fmt.Println("seta or setb:", Or(seta, setb))
	fmt.Println("seta sub setb:", seta.Sub(setb))
}

func main() {
	test()
}

