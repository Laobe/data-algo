package chain

type SqDoubleStack struct {
	data [MAXSIZE]SElemType
	top1 int
	top2 int
}

func NewSqDoubleStack(selem [MAXSIZE]SElemType, top1 int, top2 int) SqDoubleStack {
	s := SqDoubleStack{selem, top1, top2}
	return s
}

func (S *SqDoubleStack) Push(e SElemType, stackNumber int) (bool, error) {
	var err error
	if S.top1+1 == S.top2 {
		return false, err
	}
	if stackNumber == 1 {
		S.top1++
		S.data[S.top1] = e
	} else if stackNumber == 2 {
		S.top2--
		S.data[S.top2] = e
	}
	return true, nil
}

func (S *SqDoubleStack) Pop(e SElemType, stackNumber int) (bool, error) {
	var err error
	if stackNumber == 1 {
		if S.top1 == -1 {
			return false, err
		}
		S.data[S.top1] = e
		S.top1--
	} else if stackNumber == 2 {
		if S.top2 == MAXSIZE {
			return true, nil
		}
		S.data[S.top2] = e
		S.top2++
	}
	return true, nil
}
