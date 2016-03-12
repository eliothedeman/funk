package funk

// A Funk is anything that takes an X value and returns a Y.
// A function... like on a graph...
type Funk func(float64) float64

// Add returns a Funk that will add the result of a() and b ()
func Add(a, b Funk) Funk {
	return func(f float64) float64 {
		return a(f) + b(f)
	}
}

// Sub returns a Funk that subtracts the result of a() and b()
func Sub(a, b Funk) Funk {
	return func(f float64) float64 {
		return a(f) - b(f)
	}
}

// Div returns a Funk that divides the result of a() by b()
func Div(a, b Funk) Funk {
	return func(f float64) float64 {
		return a(f) / b(f)
	}
}

// Mult returns a Funk that will multiple the results of a() and b()
func Mult(a, b Funk) Funk {
	return func(f float64) float64 {
		return a(f) * b(f)
	}
}

// ToCurve returns a Curve that approxomates the given Funk
func ToCurve(f Funk, min, max, res float64) Curve {
	c := make(Curve, 0, int((max-min)/res))
	x := 0.0
	for x < max {
		c = append(c, Point{
			X: x,
			Y: f(x),
		})
		x += res
	}

	return c
}

// Tnen returns a new Funk that calls n on the result of f
func (f Funk) Then(n Funk) Funk {
	return func(x float64) float64 {
		return f(n(x))
	}
}
