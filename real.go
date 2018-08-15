
package jpl

import "math"

type Real float64
//type RealVector []Real

// Number interface

func(a Real) ToInteger() Integer { return Integer(int64(a)) }
func(a Real) ToReal() Real { return a }

//func(a Real) IMap(fn I1fI) Number { return Integer(fn(int64(a))) }
//func(a Real) RMap(fn R1fR) Number { return Real(fn(float64(a))) }

// Tensor interface

func(a Real) Dim() Vector { return Vector{} }
func(a Real) First() Tensor { return a }
func(a Real) Last() Tensor { return a }

func (x Real) Apply(f NNN, y Number) Real { return f(x, y).ToReal() }

func (x Real) Add(y Real) Real { return x+y }
func (x Real) Log() Real { return Real(math.Log(float64(x))) }

