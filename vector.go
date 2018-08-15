
package jpl

import "fmt"

type Vector []Number

func RealVector(n int) Vector {
   var x = make(Vector, n)
   for i := range x {
      x[i] = Real(0.0)
   }
   return x
}

func IntegerVector(n int) Vector {
   var x = make(Vector, n)
   for i := range x {
      x[i] = Integer(0)
   }
   return x
}

func Iota(n int) Vector {
   var x = make(Vector, n)
   for i := range x {
      x[i] = Integer(i)
   }
   return x
}

func Repeat(y Number, n int) Vector {
   var x = make(Vector, n)
   for i := range x {
      x[i] = Number(y)
   }
   return x
}

// Methods

func (x Vector) Dim() Vector { return Vector{Integer(len(x))} }
func (x Vector) Size() int { return len(x) }
func (x Vector) First() Number { return x[0] }
func (x Vector) Last() Number { return x[len(x)-1] }

func (a Vector) Copy() Vector {
   var x = make(Vector, len(a))
   copy(x, a)
   return x
}

/*
func (x Vector) Add(y Number) Vector {
   z := x.Copy()
   for i, xt := range x {
      //z[i] = xt.Add(y)
      z[i] = xt.Apply(ADD, y)
   }
   return z
}

func (x Vector) Addv(y Vector) Vector {
   if len(x) != len(y) {
      panic("Add: length error")
   }

   z := x.Copy()
   for i, xt := range x {
      //z[i] = xt.Add(y[i])
      z[i] = xt.Apply(ADD, y[i])
   }
   return z
}
*/

func (x Vector) Sel(y Vector) (Vector, error) {
   var z Vector
   n := len(x)
   for _, t := range y {
      k := int(t.ToInteger())
      if k < 0 || k > n-1 {
         return nil, fmt.Errorf("array.sel: index out of bound: k = %d", k)
      }
      z = append(z, x[k])
   }
   return z, nil
}

func (x Vector) Seln(y Vector, d int) (Vector, error) {
   //k := x.Dim()-d
   k := 1-d

   if k != 1 {
      return nil, fmt.Errorf("array.seln: dimension out of bound: d = %d", d)
   }
   return x.Sel(y)
}

func (a Vector) Reduce(fn NNN) Number {
   if len(a) < 1 {
      panic("array too short")
   }
   var s Number = a[0]
   for _, y := range a[1:] {
      s = fn(s, y)
   }
   return s
}

// Map: on a copy
func (a Vector) Map(fn NN) Vector {
   var x = make(Vector, len(a))
   for i, y := range a {
      x[i] = fn(y)
   }
   return x
}

// IMap: in place
func (a Vector) IMap(fn NN) Vector {
   for i, y := range a {
      a[i] = fn(y)
   }
   return a
}

// Apply: on a copy
func (x Vector) Apply(fn NNN, y Vector) Vector {
   if len(x) != len(y) {
      panic("arrays of unequal lengths")
   }
   var t = make(Vector, len(x))
   for i := range x {
      t[i] = fn(x[i], y[i])
   }
   return t
}

// IApply: in place
func (a *Vector) IApply(fn NNN, b Vector) *Vector {
   if len(*a) != len(b) {
      panic("arrays of unequal lengths")
   }
   for i := range *a {
      (*a)[i] = fn((*a)[i], b[i])
   }
   return a
}

func (a Vector) Filter(fn N1fb) Vector {
   var x Vector
   for _, y := range a {
      if fn(y) {
         x = append(x, y)
      }
   }
   return x
}

func (x Vector) Dot(y Vector) Number {
   if len(x) != len(y) {
      panic("Vector: Dot")
   }
   switch x[0].(type) {
   case Real:
      var s Real = 0
      for j, xt := range x {
         s += xt.ToReal()*y[j].ToReal()
      }
      return s
   case Integer:
      var s Integer = 0
      for j, xt := range x {
         s += xt.ToInteger()*y[j].ToInteger()
      }
      return s
   default:
      panic("Vector: Dot")
   }
}

