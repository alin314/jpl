
package jpl

type Number interface {
   ToInteger() Integer
   ToReal() Real
   //Apply(NNN, Number) Number

   // Tensor interface
   Tensor
}

//type NN func(Number) Number
//type NNN func(Number, Number) Number
type N1fb func(Number) bool
type N2fb func(Number, Number) bool

var ADD NNN = func(x, y Number) Number { return x.ToReal()+y.ToReal() }
var SUB NNN = func(x, y Number) Number { return x.ToReal()-y.ToReal() }
var MUL NNN = func(x, y Number) Number { return x.ToReal()*y.ToReal() }
var DIV NNN = func(x, y Number) Number { return x.ToReal()/y.ToReal() }

var ADDr RRR = func(x, y Real) Real { return x+y }
var ADDi III = func(x, y Integer) Integer { return x+y }
var GT N2fb = func(x, y Number) bool { return x.ToReal()>y.ToReal() }
var LT N2fb = func(x, y Number) bool { return x.ToReal()<y.ToReal() }

// reduce binary to monad functions
/*
func (fn NNN) Left(x Number) NN {
   return func(y Number) Number { return fn(x, y) }
}
func (fn NNN) Right(y Number) NN {
   return func(x Number) Number { return fn(x, y) }
}
*/

// reduce binary to monad predicates
func (fn N2fb) Left(x Number) N1fb {
   return func(y Number) bool { return fn(x, y) }
}
func (fn N2fb) Right(y Number) N1fb {
   return func(x Number) bool { return fn(x, y) }
}

