package quick

func Sort(ls []int) []int {
	if len(ls) <= 1 {
		return ls
	}

	m := ls[0]
	var lss []int
	var mds []int
	var grs []int
	for _, l := range ls {
		if m < l {
			grs = append(grs, l)
		} else if m > l {
			lss = append(lss, l)
		} else {
			mds = append(mds, l)
		}
	}

	var rs []int
	rs = append(rs, Sort(lss)...)
	rs = append(rs, mds...)
	rs = append(rs, Sort(grs)...)
	return rs
}
