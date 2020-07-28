package exercise_fibonacci_closure

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	f0 := 0
	f1 := 1
	return func() int {
		ret := f0
		n := f0 + f1
		f0 = f1
		f1 = n
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
