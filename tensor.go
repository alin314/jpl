
package jpl

type Tensor interface {
   Dim() Vector
   //Sel(Vector) (Tensor, error)
   //Seln(Vector, int) (Tensor, error)

   //First() Tensor
   //Last() Tensor
   //Add(Tensor) Tensor
}

