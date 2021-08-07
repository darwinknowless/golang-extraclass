package main

import (
	"fmt"
	"math"
)

// 1. create interface
type hitung interface {
	// create function list
	luas() float64
	keliling() float64
}

// 2. create struct lingkaran
type lingkaran struct {
	diameter float64
}

// 3. create struct method
func (ling lingkaran) jariJari() float64 {
	return ling.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

// 4. create struct
type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

func main() {
	var bangunDatar hitung

	bangunDatar = persegi{10.0}
	fmt.Println("persegi")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())

	bangunDatar = lingkaran{diameter: 14.0}
	fmt.Println("lingkaran")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())
	fmt.Println("jari-jari :", bangunDatar.(lingkaran).jariJari())

	p := Person{Name: "agus"}
	Hello(p)

}

// create interface
type Name interface {
	GetName() string
}

// create struct
type Person struct {
	Name string
	Age  int
}

// create method struct
func Hello(name Name) {
	fmt.Println("Hello", name.GetName())
}

func (p Person) GetName() string {
	return p.Name
}
