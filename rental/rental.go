package rental

type Rental struct {
	Movie Movie
	Days  int
}

func NewRental(title string, days int) *Rental {
	return &Rental{
		Movie: GetType(title),
		Days:  days,
	}
}

func (r *Rental) GetPoints() int {
	return r.Movie.GetPoints(r.Days)
}

func (r *Rental) GetFee() int {
	return r.Movie.GetFee(r.Days)
}
