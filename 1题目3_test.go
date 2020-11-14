package pos

import (
	"fmt"
	"testing"
)

func TestShouldReturnReceipt_timu3(t *testing.T) {
	arr := []string{"acc "}
	// sort.Strings(arr)
	arr = append(arr[:0], arr[1:]...)
	fmt.Println(arr)

	m := Constructor_1()
	a1 := m.addEvent(2, "save a apple", 2, 1)
	fmt.Println("add", a1, m.list[:10])
	a2 := m.addEvent(3, "Save a apple", 2, 2)
	fmt.Println("add", a2, m.list[:10])
	a3 := m.addEvent(3, "d a apple", 3, 2)
	fmt.Println("add", a3, m.list[:10])
	a4 := m.addEvent(3, "d a apple ", 3, 2)
	fmt.Println("add", a4, m.list[:10])
	f1 := m.finishEvent(3, "d a apple")
	fmt.Println("finish", f1, m.list[:10])
	f2 := m.finishEvent(3, "d a apple")
	fmt.Println("finish", f2, m.list[:10])
	f3 := m.finishEvent(2, "d a apple")
	fmt.Println("finish", f3, m.list[:10])
	r1 := m.removeEvent(2, "d a apple")
	fmt.Println("remove", r1, m.list[3])
	r2 := m.removeEvent(3, "d a apple ")
	fmt.Println("remove", r2, m.list[:10])
	q1 := m.queryTodo(0, 4)
	fmt.Println("query", q1)
	// fmt.Println("remove", r1)
	// r1 :=m.removeEvent(2, "drink water")
	// fmt.Println("remove", r1)
	// m.finishEvent(1, "drink water")
	// m.addEvent(0, "eat a apple", 5, 2)
	// m.addEvent(2, "eat a apple", 1, 1)
	// q1 := m.queryTodo(0, 4)
	// fmt.Println("query", q1)
	// q2 := m.queryTodo(10, 11)
	// fmt.Println("query", q2)
	// f1 := m.finishEvent(2, "eat a apple")
	// fmt.Println("finish", f1)
	// a3 := m.addEvent(2, "save a apple", 1, 1)
	// fmt.Println("add", a3)
}
