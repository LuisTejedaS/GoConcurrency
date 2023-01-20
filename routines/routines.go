package routines

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Execute() {
	waitG()
}

func waitG() {
	start := time.Now()
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go showGoRoutine(i, wg)
	}
	wg.Wait()
	fmt.Printf("%dms\n", time.Since(start).Milliseconds())
}

func showGoRoutine(id int, wg *sync.WaitGroup) {
	delay := rand.Intn(500)
	fmt.Printf("Goroutine #%d with %dms\n", id, delay)
	time.Sleep(time.Millisecond * time.Duration(delay))
	wg.Done()
}
