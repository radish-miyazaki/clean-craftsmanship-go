package sine

import "math"

type TaylorCalculator struct{}

func (s *TaylorCalculator) Sin(r float64) float64 {
	var sin float64 = 0
	for n := 1; n < 10; n++ {
		sin += s.CalculateTerm(r, n)
	}
	return sin
}

func (s *TaylorCalculator) Cos(r float64) float64 {
	return s.Sin((r + math.Pi) / 2)
}

func (s *TaylorCalculator) CalculateTerm(r float64, n int) float64 {
	sign := float64(s.calcSign(n - 1))
	power := s.calcPower(r, 2*n-1)
	factorial := s.calcFactorial(2*n - 1)
	return sign * power / factorial
}

func (*TaylorCalculator) calcFactorial(n int) float64 {
	var fac float64 = 1
	for i := 1; i <= n; i++ {
		fac *= float64(i)
	}
	return fac
}

func (*TaylorCalculator) calcPower(r float64, n int) float64 {
	var power float64 = 1
	for i := 0; i < n; i++ {
		power *= r
	}
	return power
}

func (*TaylorCalculator) calcSign(n int) int {
	sign := 1
	for i := 0; i < n; i++ {
		sign *= -1
	}
	return sign
}

type TaylorCalculatorSpy struct {
	facN   int
	powerR float64
	powerN int
	signN  int
	real   TaylorCalculator
}

func NewTaylorCalculatorSpy(real TaylorCalculator) *TaylorCalculatorSpy {
	return &TaylorCalculatorSpy{
		real: real,
	}
}

func (s *TaylorCalculatorSpy) GetR() float64 {
	return s.powerR
}

func (s *TaylorCalculatorSpy) GetRPower() int {
	return s.powerN
}

func (s *TaylorCalculatorSpy) GetFac() int {
	return s.facN
}

func (s *TaylorCalculatorSpy) GetSignPower() int {
	return s.signN
}

func (s *TaylorCalculatorSpy) calcFactorial(n int) float64 {
	s.facN = n
	return s.real.calcFactorial(n)
}

func (s *TaylorCalculatorSpy) calcPower(r float64, n int) float64 {
	s.powerR = r
	s.powerN = n
	return s.real.calcPower(r, n)
}

func (s *TaylorCalculatorSpy) calcSign(n int) int {
	s.signN = n
	return s.real.calcSign(n)
}

func (s *TaylorCalculatorSpy) CalculateTerm(r float64, n int) float64 {
	sign := float64(s.calcSign(n - 1))
	power := s.calcPower(r, 2*n-1)
	factorial := s.calcFactorial(2*n - 1)
	return sign * power / factorial
}
