package channels

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

func Execute() {
	wg := &sync.WaitGroup{}
	IDsChan := make(chan string)
	FakeIDsChan := make(chan string)
	ClosedChannels := make(chan int)

	wg.Add(3)

	go generateIDs(wg, IDsChan, ClosedChannels)
	go generateFakeIDs(wg, FakeIDsChan, ClosedChannels)
	go logIDs(wg, IDsChan, FakeIDsChan, ClosedChannels)

	wg.Wait()
}

func generateFakeIDs(wg *sync.WaitGroup, fakeIDsChan chan<- string, closedChannels chan<- int) {

	for i := 0; i < 30; i++ {
		id := uuid.New()
		fakeIDsChan <- fmt.Sprintf("%d. %s", i+1, id.String())
	}
	close(fakeIDsChan)
	closedChannels <- 1
	wg.Done()
}

func generateIDs(wg *sync.WaitGroup, idsChan chan<- string, closedChannels chan<- int) {

	for i := 0; i < 50; i++ {
		id := uuid.New()
		idsChan <- fmt.Sprintf("%d. %s", i+1, id.String())
	}
	close(idsChan)
	closedChannels <- 1
	wg.Done()

}

func logIDs(wg *sync.WaitGroup, idsChan <-chan string, fakeIDsChan <-chan string, closedChannels chan int) {
	closedCounter := 0
	for {
		select {
		case id, ok := <-idsChan:
			if ok {
				fmt.Println("ID: ", id)
			}
		case id, ok := <-fakeIDsChan:
			if ok {
				fmt.Println("Fake ID: ", id)
			}
		case _, ok := <-closedChannels:
			if ok {
				closedCounter++
			}
		}
		if closedCounter == 2 {
			close(closedChannels)
			break
		}
	}
	wg.Done()
}
