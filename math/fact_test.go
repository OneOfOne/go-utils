package math

import "testing"

const fact100 = "93326215443944152681699238856266700490715968264381621468592963895217599993229915608941463976156518286253697920827223758251185210916864000000000000000000000000"

func TestFactorialBig20(t *testing.T) {
	if FactorialBig(20).Uint64() != Factorial(20) {
		t.Errorf("FactorialBig(20) != Factorial(20)")
	}
}

func TestFactorialBig100(t *testing.T) {
	f100 := FactorialBig(100)
	if f100.String() != fact100 {
		t.Errorf("Factorial(100) returned %q, expected %q.", f100, fact100)
	}
}
