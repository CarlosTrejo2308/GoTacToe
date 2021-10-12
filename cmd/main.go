package main

type Player int

const (
	empty Player = iota
	ia
	human
)

type Board [3][3]int

func main() {
	println("Hello, world.")
}
