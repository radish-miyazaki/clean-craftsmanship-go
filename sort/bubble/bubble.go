package bubble

// Sort bubble sort
func Sort(ls []int) []int {
	if len(ls) > 1 {
		for l := len(ls) - 1; l > 0; l-- {
			for i := 0; i < l; i++ {
				f := ls[i]
				s := ls[i+1]
				if f > s {
					ls[i] = s
					ls[i+1] = f
				}
			}
		}
	}

	return ls
}
