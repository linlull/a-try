
//2.自己实现的集合
//
//    题目描述：
//    利用Go实现一个集合，并且具有如下操作：
//    为集合增加一个元素：```func (set *Set) AddItem(item Type) (err error)```
//    查找集合中是否包含一个元素：```func (set *Set) Contain(item Type) (exist bool)```
//    获取集合中的元素的个数：```func (set *Set) Size() (size int)```
//    从集合中删除一个元素：```func (set *Set) Delete(item Type) (err error)```
//    返回一个包含所有元素的数组切片：```func (set *Set) Items() (sli []Type, err error)```
//    注意：以上函数中的Set和Type需要自己定义（使用type关键字
//    提示：可以使用数组切片来进行实现
package main

import "fmt"

type Set struct{
	s []Item
}
//增加一个元素

var set  Set

func main(){
	set.s = []Item{1, 2, 3}

	fmt.Printf("添加元素%d", 4, set.AddItem(4))
	fmt.Printf("是否存在%d %t\n", 2,set.Contain(2))
	fmt.Printf("集合长度为%d",set.Size())
}
type Item int
func (set *Set) AddItem(n Item) (err error){
	set.s = append(set.s, n)
	fmt.Println(set.s)
	return err
}
//查找是否包含某个元素
func (set *Set) Contain(m Item) (exist bool){
	for i := 0; i <= len(set.s); i++{
		if set.s[i] == m{
			return true
		}
	}
	return false
}
func (set *Set) Size() (size int){
	return len(set.s)
}
