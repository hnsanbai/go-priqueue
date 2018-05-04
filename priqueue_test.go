package priqueue

import (
	"testing"
	"math/rand"
	"runtime"
	"sync"
	"time"
	"sort"
)
func TestPush(t *testing.T) {
	runtime.GOMAXPROCS(4)
	
	notices	:= make(Notices, 0, 180)
	var wg sync.WaitGroup
	var lock sync.Mutex
	
	wg.Add(4)
	
	for i := 0; i < 3; i++ {
		go func() {
			for j := 0; j < 60; j++ {
				d := rand.Intn(10000-100)+100
				notice := &Notice{
					Weight:			d,
					PhoneNumber:	13000000000,
					Message:		"我是通知消息",
				}
				t.Logf("add %d\n", d)
				
				lock.Lock()
				notices.Push(notice)
				lock.Unlock()
			}
			wg.Done()
		}()
	}
	
	go func() {
		counter := 0
		for {
			if counter > 10 {
				break
			}
			
			lock.Lock()
			var ne *Notice
			index := 0
			length := len(notices)
			if length > 0 {
				n := notices.Pop()
				ne = n.(*Notice)
				length = len(notices);
				index = sort.Search(len(notices), func(i int) bool { return notices[i].Weight > ne.Weight })
			}
			lock.Unlock()
			
			if length > 0 {
				if index < length {
					t.Fatalf("expected weight: %d, index: %d, length: %d", ne.Weight, index, length )
				}
				t.Logf("Pop %d\n", ne.Weight)
			} else {
				time.Sleep(100 * time.Millisecond)
				counter++
			}
			
		}
		wg.Done()
		t.Log("end-----------------------")
	}()
	
	wg.Wait()
}