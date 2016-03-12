# Funk [![Build Status](https://travis-ci.org/eliothedeman/funk.svg?branch=master)](https://travis-ci.org/eliothedeman/funk) [![Coverage Status](https://coveralls.io/repos/github/eliothedeman/funk/badge.svg?branch=master)](https://coveralls.io/github/eliothedeman/funk?branch=master)
Do arithmetic on functions. Treat slices like functions. Do what ever you want.

### Add functions together
```go
add1 := func(x float64) float64 {
	return x + 1
}

add2 := func(x float64) float64 {
	return x + 2
}

cmb := funk.Add(mult, div)
cmd(1) == add1(1) + add2(1)
cmd(1) == 5
```

### String functions together
```go
c := a.Then(b)
c(x) == a(b(x))
```


### Convert functions into descrite approxomations
```go
f := func(x float64) float64 {
	// ... do something complicated to x
	return x
}

// Convert f into a curve that is has points between -100 and 100
// With a resolution of 1
c := funk.ToCurve(f, -100, 100, 1)

// Encode to json
buff, _ := json.Marshal(buff)

// Decode
newCurve := funk.Curve{}
json.Unmarshal(buff, &newCurve)

// Convert back to a Funk
newFunk := newCurve.ToFunk()

// Now the new function will approxomate the original
f(x) =~ newFunk(x)
```
