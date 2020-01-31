package numberOfIslands

type unionFind struct {
	roots map[int]int
	count int
}

func (u *unionFind) Init(arr []int) {
	if u.roots == nil {
		u.roots = map[int]int{}
	}
	for _, v := range arr {
		u.roots[v] = v
		u.count++
	}
}

func (u *unionFind) Find(value int) int {
	if u.roots[value] != value {
		u.roots[value] = u.Find(u.roots[value])
	}
	return u.roots[value]
}

func (u *unionFind) Union(left, right int) {
	LRoot, RRoot := u.Find(left), u.Find(right)
	if LRoot != RRoot {
		u.roots[LRoot] = RRoot
		u.count--
	}
}

func (u *unionFind) Count() int {
	return u.count
}
