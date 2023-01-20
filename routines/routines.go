package routines

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Book struct {
	ID       int
	Title    string
	Finished bool
}

var books = []*Book{
	{0, "dune 1", false},
	{1, "dune", false},
	{2, "El perfume", false},
	{3, "HP", false},
	{4, "HP 2", false},
}

func ExecuteBooks() {
	start := time.Now()
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go readBook(i, wg, m)
	}
	wg.Wait()
	fmt.Printf("%dms\n", time.Since(start).Milliseconds())
}

func readBook(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
	finishBook(id, m)
	delay := rand.Intn(800)
	time.Sleep(time.Millisecond * time.Duration(delay))
	wg.Done()
}

func findBook(id int, m *sync.RWMutex) (int, *Book) {
	index := -1
	var book *Book
	m.RLock()
	for i, b := range books {
		if b.ID == id {
			index = i
			book = b
		}
	}
	m.RUnlock()
	return index, book
}

func finishBook(id int, m *sync.RWMutex) {
	i, book := findBook(id, m)
	if i < 0 {
		return
	}
	m.Lock()
	book.Finished = true
	books[i] = book
	m.Unlock()
	fmt.Printf("Finished book: %s\n", book.Title)
}

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
