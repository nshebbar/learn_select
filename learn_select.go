package main

import (
	"fmt"
	//"time"
	"sync"
)

// func multiples(i int) chan int { definition changed
//	func multiples(i int) (chan int, chan struct{}) {	
//	out := make(chan int)
//	done := make(chan struct{})
//	curVal := 0
//	go func() {
//		for {
//			select {
//			case out <- curVal * i:
//				curVal ++
//			case <-done:
//				fmt.Println("goroutine shutting down")
//				return
//			}
//
//		}
//	}()
//	return out, done
//}

func main() {
	in := make(chan int)
	in2 := make(chan int)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		for i:= 0; i < 10; i++ {
			in <- i
		}
		close(in)
		wg.Done()
	}()

	go func() {
		for i := 100; i < 110; i++ {
			in2 <- i	
		} 
		close(in2)
		wg.Done()
	}()

	go func() { 
		count := 0
		for count < 2 {
			select {
			case i, ok := <-in:
				if !ok {
					count++
					in = nil
					continue
				}
				fmt.Println("from in, result is", i*2)
			case i, ok := <-in2:
				if !ok {
					count++
					in2 = nil
					continue
				}
				fmt.Println("from in2, result is", i+2)
			}
		}
}()
	//twosCh, done := multiples(2)
	//for v := range twosCh {
	//	if v > 20 {
	//		break
	//	}
	//	fmt.Println(v)
	//}
	//close(done)
	//in := make(chan int)
//	out := make(chan int, 1) next without buffer
	//out := make(chan int)
	//time.Sleep(1 * time.Second)
//}

	// out <- 1

// this with select, doesn't cause a deadlock!
//	select {
//	case in <- 2:
//		fmt.Println("wrote 2 to in")
//	case v := <-out:
//		fmt.Println("read", v, "from out")
//	default:
//		fmt.Println("nothing else works")
//	}
//}
}