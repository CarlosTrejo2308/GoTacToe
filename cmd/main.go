package main

import "fmt"

type Player int

const (
	empty Player = iota
	ia
	human
)

type Board [3][3]Player

func (b Board) String() string {
	represent := map[Player]string{
		empty: " ",
		ia:    "X",
		human: "O",
	}

	toPrint := ""

	for i, row := range b {
		toPrint += fmt.Sprintf("%s | %s | %s", represent[row[0]], represent[row[1]], represent[row[2]])
		if i < 2 {
			toPrint += "\n---------\n"
		}
	}

	return toPrint
}

func main() {
	b := Board{}
	b[0][0] = ia
	b[0][1] = human
	b[0][2] = ia
	b[1][0] = empty
	b[1][1] = ia
	b[1][2] = human
	b[2][0] = ia
	b[2][1] = human
	b[2][2] = ia

	fmt.Println(b)
}
