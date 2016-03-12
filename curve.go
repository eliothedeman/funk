package funk

import (
	"math"
	"sort"
)

// A Curve is a collection of descrete X,Y pairs which lie on a Curve.
type Curve []Point

// A Point is an X,Y pair which desribes the location of a point on a graph
type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// FailFunk returns NaN no matter what is passed in.
func FailFunk(x float64) float64 {
	return math.NaN()
}

// Len returns the length of the curve
func (c Curve) Len() int {
	return len(c)
}

// Less returns true if i < j
func (c Curve) Less(i, j int) bool {
	return c[i].X < c[j].X
}

// Swap swaps index i with j
func (c Curve) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// ValAt returns the value at x on the curve
func (c Curve) ValAt(x float64) float64 {

	// TODO: Use smoothing to do the prediction.
	// This only handles liniar trends.

	// failure case, nothing to do
	if len(c) <= 1 {
		return FailFunk(x)
	}
	// special case: the x value is before our first point
	if x <= c[0].X {
		if x == c[0].X {
			return c[0].Y
		}

		// calculate the slope between the first two points
		slope := -1 * (c[1].Y - c[0].Y) / (c[1].X - c[0].X)

		return c[0].Y + ((c[0].X - x) * slope)

	}

	l := len(c) - 1

	// the x value is greater than our last point
	if x >= c[l].X {

		if x == c[l].X {
			return c[l].Y
		}

		slope := (c[l].Y - c[l-1].Y) / (c[l].X - c[l-1].X)

		return c[l].Y + ((x - c[l].X) * slope)

	}

	y := math.NaN()
	for i := 1; i < len(c); i++ {
		if x <= c[i].X {
			if c[i].X == x {
				return c[i].Y
			}

			slope := (c[i].Y - c[i-1].Y) / (c[i].X - c[i-1].X)
			y = c[i].Y + ((x - c[i].X) * slope)
			break
		}
	}
	return y

}

// NewCurve returns a curve created from the given slice
func NewCurve(p []Point) Curve {
	c := Curve(p)

	// TODO: remove any points that land at the same "X" value
	sort.Sort(c)
	return c
}

// ToFunk returns a Funk that will return the y value of any given x value on this curve
func (c Curve) ToFunk() Funk {

	// copy out the curve
	crv := make(Curve, len(c))
	copy(crv, c)

	return crv.ValAt
}
