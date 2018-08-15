
package jpl

type Integer int64
//type IntegerVector []Integer

// Number interface

func(a Integer) ToInteger() Integer { return a }
func(a Integer) ToReal() Real { return Real(float64(a)) }

//func(a Integer) IMap(fn I1fI) Number { return Integer(fn(int64(a))) }
//func(a Integer) RMap(fn R1fR) Number { return Real(fn(float64(a))) }

// Tensor interface

func(a Integer) Dim() Vector { return Vector{} }
func(a Integer) First() Tensor { return a }
func(a Integer) Last() Tensor { return a }

func(a Integer) Sel(y Vector) (Tensor, error) {
   return a, nil
}

func(a Integer) Seln(y Vector, d int) (Tensor, error) {
   return a, nil
}

func (x Integer) Apply(f NNN, y Number) Number { return f(x, y) }

