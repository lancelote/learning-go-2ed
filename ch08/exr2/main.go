package main

import (
	"fmt"
	"strconv"
)

type Printable interface {
	fmt.Stringer
	~int | ~float64
}

type PrintInt int

func (pi PrintInt) String() string {
	return strconv.Itoa(int(pi))
}

type PrintFloat float64

func (pf PrintFloat) String() string {
	return strconv.FormatFloat(float64(pf), 'f', -1, 64)
}

func Print[T Printable](t T) {
	fmt.Println(t)
}

func main() {
	var pi PrintInt = 41
	Print(pi)

	var pf PrintFloat = 42.5
	Print(pf)
}
