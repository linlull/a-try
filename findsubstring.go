//1.寻找子串
//    现在给定字符串stra、strb ，要求在stra中寻找到strb第一次出现的位置（若不存在则返回err）
//    题目要求：
//    编写函数进行解决，即文件中不应该只包括主函数，应当声明一个函数
//    （函数原型为：```func Find(stra, strb string) (position int, err error)```）即函数Find接受两个字符串stra和strb，返回两个值，第一个值为第一次出现的位置，第二个值为错误对象。若能查找成功，即strb能在stra中被找到，则返回位置和nil，若不能，则返回-1和err对象
//    提示：
//    使用errors.New(string)来进行创建error对象
//
//    进阶要求：
//    改写函数Find(string, string)为```func (stra *string)Find(strb string) (position int, err error)```
//
//    将函数的功能升级为，查找出strb在stra中出现的所有次数，使用数组返回。函数原型如下。
//    ```func (stra *string)FindAll(strb string) (positions []int, err error)```
package main

import (
	"errors"
	"fmt"
	"strings"
)

func Find(stra, strb string) (int, error){
	position := strings.Index(stra, strb)//内置函数strings.Index查找子串位置
	if position < 0 {
		return -1, errors.New("can't find it")//errors.New(strings)来返回错误信息
	}
	return position, nil
}

func test1(){
	a := "linlull"
	b := "lu"
	fmt.Println(Find(a,b))

}

type ss string
func (stra *ss)Find(strb string) (int, error){
	position := strings.Index(string(*stra), strb)//需要将自定义类型转为内置类型string才能对其使用内置函数！！！！！！
	if position < 0 {
		return -1, errors.New("can't find it")
	}
	return position, nil

}
func test2(){

	var stra ss = "linlull"//一定要明确声明a的类型不然会报错（unresolved reference Find)
	strb := "a"
	fmt.Println(stra.Find(strb))
}
func (stra *ss)FindAll(strb string) ([]int, error) {
	res := []int{} //设置一个空切片用于存放结果
	base := 0
	for strings.Index(string(*stra)[base:],strb)>=0{
	position := base + strings.Index(string(*stra)[base:],strb)//第一次查找到元素后，后面查找到的位置要加上上一次截去的长度
	res = append(res, position)
	base += len(strb)}
	if len(res) == 0 {
		return res, errors.New("can't find")
	}
	return res, nil
}
func test3(){
	var stra ss ="linlullu"
	strb := "l"
	fmt.Println(stra.FindAll(strb))
}
func main(){
	test1()
	test2()
	test3()
}