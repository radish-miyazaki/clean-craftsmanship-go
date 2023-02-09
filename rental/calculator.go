package rental

type Calculator struct {
	Rentals []Rental
}

func NewCalculator() *Calculator {
	return &Calculator{}
}

func (c *Calculator) AddRental(title string, days int) {
	c.Rentals = append(c.Rentals, *NewRental(title, days))
}

func (c *Calculator) GetRentalFee() int {
	fee := 0
	for _, r := range c.Rentals {
		fee += r.GetFee()
	}
	return fee
}

func (c *Calculator) GetRentalPoints() int {
	points := 0
	for _, r := range c.Rentals {
		points += r.GetPoints()
	}
	return points
}
