package chain

type String string

func (S String) Index(T String, pos int) int {
	var n, m, i int
	var sub String
	if pos > 0 {
		n = len(S)
		m = len(T)
		i = pos
		for i <= n-m+1 {
			end := i + m
			if end >= n {
				end = n
			}

			sub = S[i:end]
			if sub != T {
				i++
			} else {
				return i
			}
		}

	}
	return 0
}
