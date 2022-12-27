package data

import "fmt"

func loopGoroutineVal(values []string) {
	for index, val := range values {
		go func() {
			fmt.Println(index) // want "loopgoroutinecheck: loop variable `index` captured by func literal in go statement might have unexpected values"
			fmt.Println(val)   // want "loopgoroutinecheck: loop variable `val` captured by func literal in go statement might have unexpected values"
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i) // want "loopgoroutinecheck: loop variable `i` captured by func literal in go statement might have unexpected values"
		}()
	}
}
