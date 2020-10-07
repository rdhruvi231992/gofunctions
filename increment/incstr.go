package main

import "fmt"

func main() {
	Counter := &Counter{10}

	// j = 10
	for i := 0; i < 10; i++ {
		Counter.increment()
		fmt.Println("currentValue ", Counter.currentValue())
		// fmt.Println(j)
		// fmt.Println()
	}

}

//Counter struct
type Counter struct {
	count int
}

//currentValue()
func (set *Counter) currentValue() int {
	return set.count

}

func (set *Counter) increment() {

	set.count++
}
