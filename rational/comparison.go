package rational

// GreaterThan returns true if a rational is greater than the passed one.
func (ev Rational) GreaterThan(e Rational) bool {
	solveNegatives(&ev.Numerator, &ev.Denominator)
	solveNegatives(&e.Numerator, &e.Denominator)
	if e.GetDenominator() == 0 {
		return ev.Float64() > 0
	}
	return ev.Numerator*e.Denominator > e.Numerator*ev.Denominator
}

// GreaterThanNum returns true if a rational is greater than the passed integer.
func (ev Rational) GreaterThanNum(i int64) bool {
	return ev.GreaterThan(New(i, 1))
}

// LessThan returns true if a rational is less than the passed one.
func (ev Rational) LessThan(e Rational) bool {
	solveNegatives(&ev.Numerator, &ev.Denominator)
	solveNegatives(&e.Numerator, &e.Denominator)
	if e.GetDenominator() == 0 {
		return ev.Float64() < 0
	}
	return ev.Numerator*e.Denominator < e.Numerator*ev.Denominator
}

// LessThanNum returns true if a rational is less than the passed integer.
func (ev Rational) LessThanNum(i int64) bool {
	return ev.LessThan(New(i, 1))
}
