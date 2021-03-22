package main

import "fmt"

type Employer interface {
	Calc() float32
}

type Programer struct {
	name  string
	base  float32
	extre float32
}

func NewProgramer(name string, base float32, extre float32) Programer {
	return Programer{
		name:  name,
		base:  base,
		extre: extre,
	}
}

func (p Programer) Calc() float32 {
	return p.base
}

type Sale struct {
	name  string
	base  float32
	extre float32
}

func NewSale(name string, base float32, extre float32) Sale {
	return Sale{
		name:  name,
		base:  base,
		extre: extre,
	}
}

func (s Sale) Calc() float32 {
	return s.base + s.extre*s.base*0.5
}

func calAll(e []Employer) float32 {
	var cost float32
	for _, v := range e {
		cost = cost + v.Calc()
	}
	return cost
}

func main() {

	p1 := NewProgramer("1", 1500, 0)
	p2 := NewProgramer("2", 1500, 0)
	p3 := NewProgramer("3", 1500, 0)

	s1 := NewSale("s1", 800, 2)
	s2 := NewSale("s2", 900, 2.5)
	s3 := NewSale("s3", 100, 3.5)
	s4 := NewSale("s4", 1000, 3.5)


	var employList []Employer
	employList = append(employList, p1)
	employList = append(employList, p2)
	employList = append(employList, p3)

	employList = append(employList, s1)
	employList = append(employList, s2)
	employList = append(employList, s3)
	employList = append(employList, s4)

	cost := calAll(employList)
	fmt.Println(cost)
}
