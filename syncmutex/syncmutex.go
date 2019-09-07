package main

import (
	"fmt"
	"sync"
	"time"
)

type throne struct {
	sittingOrder []string
	mux          sync.Mutex
}

func (t *throne) sitdown(person string) {
	t.mux.Lock() //comment these out and sometimes a person will be missing
	t.sittingOrder = append(t.sittingOrder, person)
	t.mux.Unlock()
}

func main() {
	t := throne{sittingOrder: make([]string, 0)}
	for _, character := range []string{"akira", "tetsuo", "kei", "yamagata"} {
		go t.sitdown(character)
	}

	time.Sleep(time.Second)
	fmt.Println(t.sittingOrder)
}
