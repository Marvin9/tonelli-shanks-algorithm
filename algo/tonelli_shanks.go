package algo

import "fmt"

/*
- solve for r, in congurence(==)
	(r)^2 = n (mod p)  ~~~  sqrt(n modulo p)
	 p is prime.

- non-zero n & odd prime p, n has square root if and only if:

	(n)^((p-1)/2) == 1 (mod p)

- square root of n modulo p
*/

func bigPow(base, power, mod int64) int64 {
	var ans int64 = 1
	var i int64
	for i = 1; i <= power; i++ {
		ans *= base
		if mod != 0 {
			ans %= mod
		}
	}
	return ans
}

func debug(msg interface{}) {
	fmt.Println(msg)
}

func eulerCriterion(n, p int64) int64 {
	return bigPow(n, (p-1)/2, p)
}

// TonelliShanks - to solve for r in a congruence of the form r2 ≡ n (mod p),
// where p is a prime: that is, to find a square root of n modulo p.
// https://en.wikipedia.org/wiki/Tonelli%E2%80%93Shanks_algorithm
// it returns one solution if exists otherwise -1
func TonelliShanks(n, p int64) int64 {
	if eulerCriterion(n, p) != 1 {
		return -1
	}

	// debug("In euler's criteria")

	// step: 1
	ded := p - 1
	var Q, S int64 = 0, 0
	for ded > 0 {
		if ded%2 == 0 {
			S++
			ded /= 2
		} else {
			Q = ded
			break
		}
	}

	// debug(fmt.Sprintf("Q:%v, S:%v", Q, S))

	// step: 2
	var z int64
	var zNotInEulerCriteria int64 = 2
	for eulerCriterion(zNotInEulerCriteria, p) == 1 {
		zNotInEulerCriteria++
	}
	z = zNotInEulerCriteria

	// debug(fmt.Sprintf("Z:%v", z))

	// step: 3
	m := S
	c := bigPow(z, Q, p)
	t := bigPow(n, Q, p)
	R := bigPow(n, (Q+1)/2, p)

	for {
		if t == 0 {
			return 0
		}
		if t == 1 {
			return R
		}
		var nextI int64 = 1
		for bigPow(t, 1<<nextI, p) != 1 && nextI < m {
			// debug(bigPow(t, 1<<nextI, p))
			nextI++
		}
		// debug(fmt.Sprintf("nextI: %v", nextI))

		if nextI == m {
			fmt.Println("Unexpected error in algorithm, possible negative shifting.")
			return -1
		}

		b := bigPow(c, 1<<(m-nextI-1), p)
		bSqr := bigPow(b, 2, p)
		m = nextI
		c = bSqr
		t = (t * bSqr) % p
		R = (R * b) % p

		// debug(fmt.Sprintf("After iteration:\nnextI:%v, b:%v, m:%v, c:%v, t:%v, R:%v", nextI, b, m, c, t, R))
	}
}
