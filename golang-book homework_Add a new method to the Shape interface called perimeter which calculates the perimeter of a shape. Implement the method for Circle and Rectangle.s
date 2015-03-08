package main

import ("fmt"
        "match"
)

type Circle struct {
      x, y, r float64
      }
      
type Rectangle struct{
x1, y1, x2, y2 float64
}

func (c *Circle) perimeter () float64{
return math.Pi * c.r * 2
}

func distance (x1,y1,x2,y2 float64) float64{
a:=x2 - x1
b:=y2 - y1
return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) perimeter () float64{
l:= distance(r.x1, r.y1, r.x1, r.y2)
w:= distance(r.x1, r.y1, r.x2, r.y1)
return 1*2 + w*2
}

type Shape interface{
perimeter() float64
}

func main(){

c:=Rectangle{4,5}
fmt.Println(c.perimeter())
}
