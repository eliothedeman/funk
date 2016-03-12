# funk
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
