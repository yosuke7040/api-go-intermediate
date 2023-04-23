package main

import (
	"fmt"
	"sync"
)

// func cutIngredient() {
// 	fmt.Println("start cutting ingredients")
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("finish cutting ingredient!")
// }
// func boilWater() {
// 	fmt.Println("start boiling water")
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("The water has boiled!")
// }

// func main() {
// 	go cutIngredient()
// 	boilWater()
// }

var i int
var mu sync.Mutex
var wg sync.WaitGroup

func MyFunc1() {
	defer wg.Done()
	fmt.Println("MyFunc1 start")

	mu.Lock()
	i += 1
	mu.Unlock()

	fmt.Println("MyFunc1 finish")
}

func MyFunc2() {
	defer wg.Done()
	fmt.Println("MyFunc2 start")

	mu.Lock()
	i -= 1
	mu.Unlock()

	fmt.Println("MyFunc2 finish")
}

func main() {
	i = 0

	wg.Add(2)
	go MyFunc1()
	go MyFunc2()

	wg.Wait()
	fmt.Printf("final result: i = %d\n", i)
}
