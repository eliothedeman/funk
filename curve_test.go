package funk

import (
	"fmt"
	"math"
	"sort"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewCurve(t *testing.T) {
	Convey("With a slice of Points", t, func() {
		s := []Point{
			{0, 0},
			{1, 1},
			{2, 2},
			{3, 3},
			{4, 4},
		}

		Convey("When converted into a curve", func() {
			c := NewCurve(s)

			Convey("Then they should be the same length", func() {
				So(len(s), ShouldEqual, len(c))
			})

			Convey("Then they should have the same value at the same index", func() {
				for i := range s {
					So(c.ValAt(s[i].X), ShouldEqual, s[i].Y)
				}
			})

			Convey("Then a value at -1 should be -1", func() {
				So(c.ValAt(-1), ShouldEqual, -1)
			})

			Convey("Then a value at 100 should be 100", func() {
				So(c.ValAt(100), ShouldEqual, 100)
			})

			Convey("Then a value at 0.5 should be 0.5", func() {
				So(c.ValAt(0.5), ShouldEqual, 0.5)
			})
		})
	})
}

func TestOutOfOrderPoints(t *testing.T) {
	s := []Point{
		{100, 0},
		{200, 0},
		{-1, 44},
		{32, 19},
	}
	Convey(fmt.Sprintf("When an unsorted slice of poitns %v", s), t, func() {

		Convey("With a curve created from thos points", func() {
			c := NewCurve(s)

			Convey("Then the new curve should be sorted", func() {
				So(sort.IsSorted(c), ShouldBeTrue)
			})
		})
	})
}

func TestEmptyCurve(t *testing.T) {
	Convey("With an empty curve", t, func() {
		c := Curve{}

		Convey("When we look for a value at index 1", func() {

			v := c.ValAt(1)

			Convey("Then we should get NaN", func() {
				So(math.IsNaN(v), ShouldBeTrue)
			})
		})
	})
}

func TestConvertToFunc(t *testing.T) {
	c := Curve{
		{0, 0},
		{1, 2},
		{2, 4},
	}
	Convey(fmt.Sprintf("Given a curve %v", c), t, func() {

		Convey("When the curve is converted to a function", func() {
			f := c.ToFunk()

			Convey("Then the value at 2 should be 4", func() {
				So(f(2), ShouldEqual, 4)
			})

			Convey("Then the value at 100 should be 200", func() {
				So(f(100), ShouldEqual, 200)
			})
		})
	})

}
