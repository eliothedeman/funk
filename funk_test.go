package funk

import (
	"math"
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

func TestThen(t *testing.T) {

	Convey("Given a Funk that multipies it's input by 2", t, func() {
		a := Funk(func(f float64) float64 {
			return f * 2
		})
		Convey("When a is succeeded by a Funk that devides x by 2", func() {
			f := a.Then(func(x float64) float64 {
				return x / 2

			})

			Convey("Then calling f on any input should return the same value", func() {
				So(f(100), ShouldEqual, 100)
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

func TestPipe(t *testing.T) {
	Convey("Given two functions a, b", t, func() {
		a := func(x float64) float64 {
			return x * 2
		}

		b := func(x float64) float64 {
			return x * 5
		}

		Convey("When the two functions are linked together via Pipe into c", func() {
			c := Funk(a).Pipe(b)

			Convey("Then the output of c should equal b(a)", func() {

				So(c(5), ShouldEqual, b(a(5)))
			})
		})
	})
}

func TestPipeFunk(t *testing.T) {
	Convey("Given three functions a, b, c", t, func() {
		a := func(x float64) float64 {
			return x * 2
		}

		b := func(x float64) float64 {
			return x * 5
		}

		c := func(x float64) float64 {
			return x / 100
		}
		Convey("When the three functions are linked together via Pipe into d", func() {
			d := Pipe(a, b, c)

			Convey("Then the output of d should equal c(b(a))", func() {

				So(d(5), ShouldEqual, c(b(a(5))))
			})

			Convey("Then an empty pipe should always return NaN", func() {
				x := Pipe()
				So(math.IsNaN(x(5)), ShouldBeTrue)
			})
		})

	})

}
