
package jpl

type Array []Matrix

func (x Array) Dim() Vector {
   t := x[0].Dim()
   var z = make(Vector, len(t)+1)
   copy(z[1:], t)
   z[0] = Integer(len(x))
   return z
}

