package generator

func NewID(i int) (id string) {
	var c int
	var res []int
	l := len(list)
	if i <= l {
		return list[i]
	}
	for {
		c = i % l
		i /= l

		res = append(res, c, i)
		if i < l {
			break
		}
	}

	for j := len(res) - 1; j >= 0; j-- {
		id += list[res[j]]
	}

	return
}
