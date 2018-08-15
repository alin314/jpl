
package jpl

type NN func(Number) Number
type NNN func(Number, Number) Number
type NVV func(Number, Vector) Vector
type RRR func(Real, Real) Real
type III func(Integer, Integer) Integer

type VN func(Vector) Number
type VV func(Vector) Vector
type VNV func(Vector, Number) Vector
type VVV func(Vector, Vector) Vector

func (f VNV) Hook(g VN) VV {
   return func(y Vector) Vector {
      return f(y, g(y))
   }
}

func (f VVV) Hook(g VV) VV {
   return func(y Vector) Vector {
      return f(y, g(y))
   }
}
func (f VVV) Apply(x, y Vector) Vector { return f(x, y) }
func (f VV) Apply(y Vector) Vector { return f(y) }

// reduce binary to monad functions
func (f NNN) Left(x Number) NN {
   return func(y Number) Number { return f(x, y) }
}
func (f NNN) Right(y Number) NN {
   return func(x Number) Number { return f(x, y) }
}
func (f NVV) Left(x Number) VV {
   return func(y Vector) Vector { return f(x, y) }
}

func (f NNN) Apply(x, y Vector) Vector { return x.Apply(f, y) }
func (f NN) Map(y Vector) Vector { return y.Map(f) }

var ADDV VVV = func (x Vector, y Vector) Vector { return x.Apply(ADD, y) }
var ADDV2 VVV = func (x Vector, y Vector) Vector { return ADD.Apply(x, y) }
var DIVf VNV = func(x Vector, y Number) Vector { return x.IMap(DIV.Right(y)) }
var DIVf2 VNV = func(x Vector, y Number) Vector { return DIV.Right(y).Map(x) }
var DIVFST VV = DIVf.Hook(Vector.First)

