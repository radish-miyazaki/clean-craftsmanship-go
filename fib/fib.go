package fib

import (
	"math/big"
)

func Fib(n int) big.Int {
	fm1 := *big.NewInt(1)
	fm2 := *big.NewInt(1)

	for i := n; i > 1; i-- {
		f := big.NewInt(0)
		f = f.Add(&fm1, &fm2)
		fm2 = fm1
		fm1 = *f
	}

	return fm1
}
