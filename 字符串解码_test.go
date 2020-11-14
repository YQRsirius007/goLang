package pos

import "testing"

func TestShouldReturnReceipt_decode(t *testing.T) {
	// data := []string{
	// 	"3[a]2[bc]",
	// 	"3[a2[c]]",
	// }
	//
	// res := []string{
	// 	"aaabcbc",
	// 	"accaccacc",
	// }
	// for i,v := range data{
	// 	c := decodeString(v)
	// 	if c!=res[i] {
	// 		t.Errorf("want %s, but got %s",res[i], c)
	// 	}
	// }

	if "abcccbcccd" != decodeString2("a(b(c)<3>)<2>d") {
		t.Errorf("want %s, but got %s", "abcccbcccd", decodeString2("a(b(c)<3>)<2>d"))
	}
}
