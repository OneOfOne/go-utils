package math

import "math/big"

var cache = []uint64{
	1,                   // 0
	1,                   // 1
	2,                   // 2
	6,                   // 3
	24,                  // 4
	120,                 // 5
	720,                 // 6
	5040,                // 7
	40320,               // 8
	362880,              // 9
	3628800,             // 10
	39916800,            // 11
	479001600,           // 12
	6227020800,          // 13
	87178291200,         // 14
	1307674368000,       // 15
	20922789888000,      // 16
	355687428096000,     // 17
	6402373705728000,    // 18
	121645100408832000,  // 19
	2432902008176640000, // 20
}

// Factorial returns the factorial n!, if n > 20 it returns 0, use FactorialBig instead
func Factorial(n uint) uint64 {
	if n < uint(len(cache)) {
		return cache[n]
	}
	return 0
}

var (
	one    = big.NewInt(1)
	twenty = big.NewInt(20)
)

func FactorialBig(n uint64) (r *big.Int) {
	//fmt.Println("n = ", n)
	bn := new(big.Int).SetUint64(n)
	r = big.NewInt(1)

	if bn.Cmp(one) <= 0 {
		return
	}

	if bn.Cmp(twenty) <= 0 {
		return r.SetUint64(Factorial(uint(n)))
	}

	for i := big.NewInt(2); i.Cmp(bn) <= 0; i.Add(i, one) {
		r.Mul(r, i)
	}

	return
}
