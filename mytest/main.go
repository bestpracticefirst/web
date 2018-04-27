package main

import "fmt"

func f() (result int) {
	defer func() { result++ }()
	return 0
}
func f1() (r int) {
	t := 5
	defer func() {
		t=t+5
	}()
	return t
}
func f2() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
func main() {
	fmt.Println(f1())
}
