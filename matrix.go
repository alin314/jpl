
package jpl

import "fmt"
//import "os"

type Matrix []Vector

func (x Matrix) Dim() Vector { return Vector{Integer(len(x)), Integer(len(x[0]))} }
func (x Matrix) First() Vector { return x[0] }
func (x Matrix) Last() Vector { return x[len(x)-1] }

func (a Matrix) Copy() Matrix {
   var x = make(Matrix, len(a))
   for i, y := range a {
      copy(x[i], y)
   }
   return x
}

func (x Matrix) Sel(y Vector) (Matrix, error) {
   var t Matrix
   n := len(x)
   for _, s := range y {
      k := int(s.ToInteger())
      if k < 0 || k > n-1 {
         return nil, fmt.Errorf("matrix.sel: index out of bound: k = %d", k)
      }
      t = append(t, x[k])
   }
   return t, nil
}

func (x Matrix) Seln(y Vector, d int) (Matrix, error) {
   //k := x.Dim()-d
   k := 2-d

   //if k < 0 || k > x.Dim() {
   if k < 0 || k > 2 {
      return nil, fmt.Errorf("matrix.seln: dimension out of bound: d = %d", d)
   }
   if k == 0 {
      return x.Sel(y)
   }
   var z Matrix
   for _, t := range x {
      tt, err := t.Seln(y, d-1)
      if err != nil {
         return nil, fmt.Errorf("matrix.seln: dimension out of bound: d = %d; %v", d, err)
      }
      z = append(z, tt)
   }
   return z, nil
}

// does not change a itself
func (x Matrix) Reduce(fn NNN, d int) Vector {
   switch d {
   case 0:
      if len(x) < 1 {
         panic("array too short")
      }
      var s Vector = x[0].Copy()
      for _, y := range x[1:] {
         //s = fn(s, y)
         s.IApply(fn, y)
      }
      return s
   case 1:
      var s Vector
      for _, xv := range x {
         s = append(s, xv.Reduce(fn))
      }
      return s
   default:
      panic("Matrix: Reduce d not recognized")
   }
}

func (x Matrix) Map(f VV) Matrix {
   var z Matrix
   for _, xt := range x {
      z = append(z, f(xt))
   }
   return z
}

func (x Matrix) Apply(f NNN, y Vector) Matrix {
   var z Matrix
   for i, xt := range x {
      z = append(z, xt.Map(f.Right(y[i])))
   }
   return z
}

func (x Matrix) Transpose() Matrix {
   var m = len(x)
   var xx = make(Matrix, len(x[0]))
   for i := range xx {
      xx[i] = make(Vector, m)
   }

   for j, xt := range x {
      for i := range xt {
         xx[i][j] = xt[i]
      }
   }
   return xx
}

