package priqueue

type Notice struct {
	Weight		int			// 值大者，优先发送
	index		int
	PhoneNumber	int64		// 对方手机号码
	Message		string		// 发送的提示信息
	
}

type Notices []*Notice


func (q *Notices) Push(n interface{}) {
	length := len(*q)
	capacity := cap(*q)
	if length + 1 > capacity {
		tq := make(Notices, length, capacity * 2)
		copy(tq, *q)
		*q = tq
	}
	*q = (*q)[0 : length + 1]
	ne := n.(*Notice)
	ne.index = length
	(*q)[length] = ne
	q.up(length)
}

func (q *Notices) Pop() interface{} {
	length := len(*q)
	capacity := cap(*q)
	(*q)[0], (*q)[length-1] = (*q)[length-1], (*q)[0]
	q.down(0, length-1)
	
	if length < (capacity / 2) && capacity > 25 {
		tq := make(Notices, length, capacity / 2)
		copy(tq, *q)
		*q = tq
	}
	ne := (*q)[length - 1]
	ne.index = -1
	*q = (*q)[0 : length - 1]
	return ne
}

func (q *Notices) up(length int) {
	for {
		i := (length-1)/2
		if length == i || (*q)[i].Weight > (*q)[length].Weight {
			break
		}
		(*q)[i], (*q)[length] = (*q)[length], (*q)[i]
		length = i
	}
}

func (q *Notices) down(i, length int) {
	index := i
	for {
		left := 2 * index + 1
		if left >= length || left < 0 {
			break
		}
		swap := left;
		if right := left + 1; right < length && (*q)[right].Weight > (*q)[left].Weight {
			swap = right
		}
		
		if (*q)[swap].Weight < (*q)[index].Weight {
			break
		}
		
		(*q)[index], (*q)[swap] = (*q)[swap], (*q)[index]
		index = swap
	}
}