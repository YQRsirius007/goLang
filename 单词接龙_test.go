package pos

import (
	"fmt"
	"testing"
)
func TestShouldReturnReceipt_ladderLength(t *testing.T) {
	s := "abc"
	fmt.Println(s[:1])
	beginWord := []string{
		"hit",
		"red",
	}
	endWord := []string{
		"cog",
		"tax",
	}
	wordList := [][]string{
		{"hot","dot","dog","lot","log","cog"},
		{"ted","tex","red","tax","tad","den","rex","pee"},
	}
	res := []int{
		5,
		4,
	}
	for i:=1; i<2; i++ {
		c := ladderLength(beginWord[i], endWord[i], wordList[i])
		if c != res[i] {
			t.Errorf("want %v, but got %v", res[i], c)
		}
	}
}