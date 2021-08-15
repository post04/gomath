package rational

import "math/big"

// Simplify simplifies the rational number by dividing it's numerator and
// denominator by the GCD.
func (ev *Rational) Simplify() {
	if ev.Denominator < 0 {
		ev.Numerator *= -1
		ev.Denominator *= -1
	}

	currentNumerator := ev.Numerator
	currentDenominator := ev.Denominator

	if currentNumerator < 0 && currentDenominator > 0 {
		currentNumerator *= -1
	} else if currentDenominator < 0 && currentNumerator >= 0 {
		currentDenominator *= -1
	}

	n := big.NewInt(currentNumerator)
	d := big.NewInt(currentDenominator)

	gcd := new(big.Int).GCD(nil, nil, n, d).Int64()

	if gcd > 1 {
		ev.Numerator /= gcd
		ev.Denominator /= gcd
	}
}

// IsNatural determines whether the rational number is also natural.
func (ev Rational) IsNatural() bool {
	if ev.Numerator%ev.Denominator == 0 {
		return true
	}
	return false
}

// Float64 returns the float64 representation of a rational number.
func (ev Rational) Float64() float64 {
	return float64(ev.Numerator) / float64(ev.Denominator)
}

// Get returns a value.
func (ev Rational) Get() (numerator, denominator int64) {
	return ev.Numerator, ev.Denominator
}

// GetNumerator returns a numerator.
func (ev Rational) GetNumerator() int64 {
	return ev.Numerator
}

// GetDenominator returns a denominator.
func (ev Rational) GetDenominator() int64 {
	return ev.Denominator
}

// GetModule returns rational number's module.
func (ev Rational) GetModule() Rational {
	solveNegatives(&ev.Numerator, &ev.Denominator)
	if ev.LessThanNum(0) {
		ev = ev.MultiplyByNum(-1)
	}
	return ev
}

// IsNull determines whether the value is zero.
func (ev Rational) IsNull() (n bool) {
	if ev.Numerator == 0 {
		n = true
	}
	return
}

// RationalsAreNull determines whether the slice of Rationals contains only zero values.
func RationalsAreNull(l []Rational) (isNull bool) {
	isNull = true
	for _, v := range l {
		if !v.IsNull() {
			isNull = false
		}
	}
	return
}

func solveNegatives(n, d *int64) {
	if *d < 0 {
		*n *= -1
		*d *= -1
	}
}
