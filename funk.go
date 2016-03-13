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

// Then returns a new Funk that calls n on the result of f
func (f Funk) Then(n Funk) Funk {
	return func(x float64) float64 {
		return f(n(x))
	}
}

// Pipe runs the functions in seperate goroutines and links them with channels.
func (f Funk) Pipe(n Funk) Funk {
	in := make(chan float64)
	link := make(chan float64)
	out := make(chan float64)

	// run f
	go func() {
		var x float64
		for {
			x = <-in
			link <- f(x)
		}
	}()

	// run n
	go func() {
		var x float64
		for {
			x = <-link
			out <- n(x)
		}
	}()

	return func(x float64) float64 {
		in <- x
		return <-out
	}
}

// Pipe a list of Funk's together
func Pipe(f ...Funk) Funk {

	if len(f) == 0 {
		return FailFunk
	}

	x := f[0]

	for i := 1; i < len(f); i++ {
		x = x.Then(f[i])
	}

	return x
}
