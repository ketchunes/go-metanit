package main

import "fmt"

func main() {
	// variables
	// var a int8 = -1
	// var b uint8 = 2
	// var c byte = 3 // byte - синоним типа uint8
	// var d int16 = -4
	// var f uint16 = 5
	// var g int32 = -6
	// var h rune = -7 // rune - синоним типа int32
	// var j uint32 = 8
	// var k int64 = -9
	// var l uint64 = 10
	// var m int = 102
	// var n uint = 105

	// fmt.Println("a: ", a)
	// fmt.Println("b: ", b)
	// fmt.Println("c: ", c)
	// fmt.Println("d: ", d)
	// fmt.Println("f: ", f)
	// fmt.Println("g: ", g)
	// fmt.Println("h: ", h)
	// fmt.Println("j: ", j)
	// fmt.Println("k: ", k)
	// fmt.Println("l: ", l)
	// fmt.Println("m: ", m)
	// fmt.Println("n: ", n)

	//const

	const pi float32 = 21.33
	fmt.Println(pi)
	const a = 11
	const b = a
	fmt.Println(a)
	const (
		lol = 32
		c
		d
		v
		s = 3
	)
	fmt.Println(lol, c, d, v, s)
	const (
		c0 = iota
		c1
		c2 = iota
	)
	fmt.Println("C0:", c0)
	fmt.Println("C1:", c1)
	fmt.Println("C2:", c2)

	const (
		c3 = iota
	)
	fmt.Println("C3:", c3)
}
