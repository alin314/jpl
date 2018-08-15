
package jpl

import "fmt"

func main() {
   var x Vector = Iota(10)
   fmt.Println(x)
   fmt.Println(x.Reduce(ADD))
   fmt.Println(x.Map(ADD.Right(Integer(1))))

   x = x.Apply(ADD, Iota(5))
   fmt.Println(x)
   fmt.Println(x.Reduce(SUB))
   fmt.Println(append(x, Iota(3)...))

   fmt.Println(x)
   fmt.Println(x.Filter(GT.Right(Integer(0))))
   fmt.Println(x.Filter(GT.Right(Integer(2))))
   fmt.Println(x.Filter(GT.Left(Integer(2))))

   fmt.Println(x.Map(SUB.Left(Integer(2))))
   fmt.Println(x.Map(SUB.Right(Integer(2))))
}
