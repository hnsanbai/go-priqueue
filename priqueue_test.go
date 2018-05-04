package priqueue

import (
	"testing"
	"math/rand"
)
var lastw int

func TestPush(t *testing.T) {
	notices := make(Notices, 0,10)
	t.Log(notices)
	for len(notices) < 60 {
		d := rand.Intn(1000-100)+100
		notice := &Notice{
			Weight:			d,
			PhoneNumber:	13000000000,
			Message:		"我是通知消息",
		}
		t.Logf("add %d\n", d)
		notices.Push(notice)
	}
	
	for len(notices) > 0 {
		n := notices.Pop()
		ne := n.(*Notice)
		t.Logf("Pop %d\n", ne.Weight)
		
		if lastw > ne.Weight {
			t.Fatalf("expected weight: %d and %d", lastw, ne.Weight)
		}
	}
}