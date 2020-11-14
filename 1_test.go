package pos

import (
	"fmt"
	"testing"
)

type test1 struct {
	name string
}
func (t *test1)setName(name string){

	t.name = name
}
func a(t test1)test1{
	t.setName("aaa")
	return t
}
func TestShouldReturnReceipt(t *testing.T) {
	var t1 test1
	fmt.Println("1", t1.name)
	t1 = a(t1)
	fmt.Println("2", t1.name)
	// data := []int{1,2,3,4,5,6}
	// fmt.Printf("%p\n", &data[0])
	// s := data[:3]
	// fmt.Printf("%p\n", &s)
	// s2 := append(s, 100, 200)
	// s2[4] = 999
	// fmt.Printf("%p\n", &s2[0])
	// fmt.Println(data, s, s2)
	// data := []int{
	// 	3,
	// 	10,
	// 	81,
	// 	59,
	// 	0,
	// 	94,
	// }
	// res := []int{
	// 	1,
	// 	5,
	// 	40,
	// 	29,
	// 	0,
	// 	47,
	// }
	// for i, v := range data {
	// 	c := drink(v)
	// 	if c != res[i] {
	// 		t.Errorf("want %d, but got %d", res[i], c)
	// 	}
	// }
}
// add(1, 2) = 103
func add(x, y int)(z int){
	defer func(){
		z+=100
	}()
	z=x+y
	return
}
// (z=z+200)->(call defer)->(ret)
func add1(x, y int)(z int){
	defer func(){
		print(z)
	}()
	z= x+y
	return z+200
}
