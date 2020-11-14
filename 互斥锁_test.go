package pos

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
)

// Mutex是一个互斥锁 可以创建为其他结构体的字段 0值为解锁状态 Mutex类型的锁和线程无关，可以由不同的线程加锁和解锁
// WaitGroup 用于等待一组线程的结束 父线程调用add方法设定应等待的线程的数量 每个被等待的线程在结束时应调用done方法
func TestShouldReturnReceipt_lock(t *testing.T) {
	var mutex sync.Mutex
	wait := sync.WaitGroup{}
	fmt.Println("lOCKED")
	mutex.Lock()
	for i:=1; i<3; i++{
		wait.Add(1)
		go func(i int) {
			fmt.Println("not lock:", i)
			mutex.Lock()
			fmt.Println("lock:", i)
		}(i)
	}
	// WaitGroup
	// Add 内部计数加delta 可以是负数 内部计数器为0时，wait方法阻塞等待的所有线程都被释放，计数器小于0，方法panic
	// Dond 减少计数器的值
	// Wait 阻塞直到计数器的值减为0
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _,url := range urls{
		wg.Add(1)
		go func(url string){
			defer wg.Done()
			http.Get(url)
		}(url)
	}
	wg.Wait()
	//
}