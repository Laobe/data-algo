package chain

type SqStack struct {
	data [MAXSIZE]SElemType
	top  int
}

func NewSqStack(selem [MAXSIZE]SElemType, top int) SqStack {
	s := SqStack{selem, top}
	return s
}

func (S *SqStack) Push(e SElemType) (bool, error) {
	var err error
	if S.top == MAXSIZE-1 {
		return false, err
	}
	S.top++
	S.data[S.top] = e
	return true, nil
}

func (S *SqStack) Pop(e SElemType) (bool, error) {
	var err error
	if S.top == -1 {
		return false, err
	}
	S.data[S.top] = e
	e = S.data[S.top]
	S.top--
	return true, nil
}
