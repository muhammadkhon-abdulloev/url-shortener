package generator

import "github.com/muhammadkhon-abdulloev/url-shortener/pkg/types"

func NewID(i int) (id string) {
	var c int
	var res []int
	l := len(types.List)
	if i < l {
		return types.List[i]
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
		id += types.List[res[j]]
	}

	return
}
