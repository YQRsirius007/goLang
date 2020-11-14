package pos

//
// import (
// 	"fmt"
// 	"math"
// 	"strings"
// )
//
// func getNormalDeviceNum(start, end int) int {
// 	order := []string{}
// 	board1 := []string{}
//
// 	success := map[string]int{}
// 	fail := map[string]int{}
// 	for i:=0; i<n ;i++{
// 		arr := strings.Fields(records[i])
// 		fmt.Println(arr[0], arr[0]=="C", arr[0]=="W")
// 		if arr[0]=="C"{
// 			order = append(order, arr[1])
// 			flag := false
// 			for _,s := range board1 {
// 				fmt.Println(arr[1], s, strings.HasPrefix(arr[1], s[:len(s)-1]))
// 				if arr[1]==s || (strings.HasSuffix(s, "*")&&strings.HasPrefix(arr[1], s[:len(s)-1])){
// 					success[arr[1]]++
// 					flag = true
// 					break
// 				}
// 			}
// 			if !flag{
// 				fail[arr[1]]++
// 			}
// 		}
// 		if arr[0]=="W"{
// 			board1 = append(board1, arr[1])
// 		}
// 	}
// 	for _,s := range order{
// 		fmt.Printf("%s\t%d\t%t", s, success[s], fail[s])
// 	}
// }
// func recordHas() bool{
//
// }
// func passNum(num int)int{
// 	ret := 0
// 	temp:= num
// 	cur := 1
// 	flag8 := false
// 	for temp>0{
// 		cur = cur*(temp%10)
// 		if flag8{
// 			cur --
// 			flag8 = false
// 		}
//
// 		if temp%10>=4{
// 			cur --
// 		}
// 		if temp%10>=8{
// 			flag8 = true
// 		}
// 		temp = temp/10
// 	}
// 	ret = cur
// 	return ret
// }
