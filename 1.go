package pos

import "fmt"

// exec/internal/role/ddf_manager.go
// d-discReqPara
// d.getTopoDiscReq().Get
// EQ GT GE LT LE HasPrefix IsPrefixOf IsMaxPrefixOf In Intersection

func test_1(){
	var a interface{} = "abc"
	b, ok := a.(string)
	r := []rune("abc你好")
	fmt.Println(b, ok)
	fmt.Println(r[3])
}