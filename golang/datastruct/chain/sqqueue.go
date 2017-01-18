package chain

type SqQueue struct {
	data  [MAXSIZE]QElemType
	front int
	rear  int
}

func NewSqQueue(data [MAXSIZE]QElemType, values ...int) SqQueue {
	if len(values) == 0 {
		values[0] = 0
		values[1] = 0
	} else if len(values) == 1 {
		values[1] = values[0]
	}
	return SqQueue{data, values[0], values[1]}
}

func (Q *SqQueue) QueueLength() int {
	return (Q.rear - Q.front + MAXSIZE) % MAXSIZE
}

func (Q *SqQueue) EnQueue(e QElemType) bool {
	if (Q.rear+1)%MAXSIZE == Q.front { // 队列已满
		return false
	}
	Q.data[Q.rear] = e
	Q.rear = (Q.rear + 1) % MAXSIZE
	return true
}

func (Q *SqQueue) DeQueue(e *QElemType) bool {
	if Q.front == Q.rear { // 队列为空
		return false
	}
	*e = Q.data[Q.front]
	Q.front = (Q.front + 1) % MAXSIZE
	return true
}
