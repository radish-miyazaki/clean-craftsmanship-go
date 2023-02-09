package rental

type Movie interface {
	GetFee(days int) int
	GetPoints(days int) int
}

func applyGracePeriod(amount int, days int, grace int) int {
	if days > grace {
		return amount + amount*(days-grace)
	}
	return amount
}

type RegularMovie struct {
	Title string
}

func (m *RegularMovie) GetPoints(days int) int {
	return applyGracePeriod(1, days, 3)
}

func (m *RegularMovie) GetFee(days int) int {
	return applyGracePeriod(150, days, 3)
}

type ChildrenMovie struct {
	Title string
}

func (m *ChildrenMovie) GetPoints(days int) int {
	return 1
}

func (m *ChildrenMovie) GetFee(days int) int {
	return days * 100
}
