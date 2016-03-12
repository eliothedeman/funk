package funk

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAdd(t *testing.T) {

	Convey("Given two Funks a, b", t, func() {
		a := func(f float64) float64 {
			return f - 5
		}
		b := func(f float64) float64 {
			return f + 5
		}
		Convey("When a is added to b.", func() {
			c := Add(a, b)

			Convey("The value of applying c onto 10 should be 20", func() {
				So(c(10), ShouldEqual, 20)
			})
		})
	})

}

func TestSub(t *testing.T) {

	Convey("Given two Funks a, b", t, func() {
		a := func(f float64) float64 {
			return f - 5
		}
		b := func(f float64) float64 {
			return f + 5
		}
		Convey("When b is subtracted form a", func() {
			c := Sub(a, b)

			Convey("The value of applying c onto 10 should be -10", func() {
				So(c(10), ShouldEqual, -10)
			})
		})
	})
}

func TestDiv(t *testing.T) {

	Convey("Given two Funks a, b", t, func() {
		a := func(f float64) float64 {
			return f * 2
		}
		b := func(f float64) float64 {
			return f * 4
		}
		Convey("When a is divided by b", func() {
			c := Div(a, b)

			Convey("The value of applying c onto 10 should be -10", func() {
				So(c(10), ShouldEqual, 0.5)
			})
		})
	})
}

func TestMult(t *testing.T) {

	Convey("Given two Funks a, b", t, func() {
		a := func(f float64) float64 {
			return f * 2
		}
		b := func(f float64) float64 {
			return f * 4
		}
		Convey("When a is multiplied by b", func() {
			c := Mult(a, b)

			Convey("The value of applying c onto 10 should be -10", func() {
				So(c(10), ShouldEqual, 800)
			})
		})
	})
}

func TestToCurve(t *testing.T) {

	Convey("Given a func 'a' that multiples x by 2", t, func() {
		a := func(x float64) float64 {
			return x * 2
		}

		Convey("When a is converted into a curve between 0 and 100 with a resolution 0f 0.5", func() {
			crv := ToCurve(a, 0, 100, 0.5)

			Convey("The Curve should have 200 points", func() {
				So(len(crv), ShouldEqual, 200)
				x := 0.0
				for x < 100 {
					So(crv[int(x*2)].Y, ShouldEqual, a(x))
					x += 0.5
				}
			})
		})

	})
}