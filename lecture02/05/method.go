package main

import (
	"fmt"
	"math"
)


type ages int
type money float32

// わかりましたか？簡単でしょう？このように自分のコードの中に意味のある型を定義することができるのです。
// 実際はエイリアスを定義しているだけです。
// Cのtypedefに似たようなもので、例えば上のagesはintの代わりになっています。


type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

// methodはレシーバのフィールドにアクセスできる。

func main() {
	var hoge ages = 3
	var unko money = 222

	fmt.Println("sdk:", hoge)
	fmt.Println("iaaa:", unko)


	r1 := Rectangle{12, 2}
	r2 := Rectangle{9, 4}
	fmt.Println("Area of r1 is:", r1.area())
	fmt.Println("Area of r2 is:", r2.area())

	c1 := Circle{10}
	c2 := Circle{25}
	fmt.Println("Area of r1 is:", c1.area())
	fmt.Println("Area of r2 is:", c2.area())

}


