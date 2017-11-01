package main

import (
	. "fmt"
)

type person struct {
	name string
	age  int
}

var P person

func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age
}

func main() {
	P.name = "冴草 きい"
	P.age = 16

	Me := person{"say3no", 27}
	// key:valueによって明示すれば順序はどうでもよくなる
	You := person{age: 99999, name: "God"}
	// ポインタもつくれる
	He := new(person)

	Printf("Name: %s, Age: %v \n", P.name, P.age)
	Printf("Name: %s, Age: %v \n", Me.name, Me.age)
	Printf("Name: %s, Age: %v \n", You.name, You.age)
	Printf("Name: %s, Age: %v \n", He.name, He.age)

	pmOlder, pmdiff := Older(P, Me)
	pyOlder, pydiff := Older(P, You)
	myOlder, mydiff := Older(Me, You)

	Printf("Of %s and %s, %s is older by %d years\n",
		P.name, Me.name, pmOlder.name, pmdiff)

	Printf("Of %s and %s, %s is older by %d years\n",
		P.name, You.name, pyOlder.name, pydiff)

	Printf("Of %s and %s, %s is older by %d years\n",
		P.name, Me.name, myOlder.name, mydiff)

}
