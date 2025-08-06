package keycounter

import (
	"fmt"
	"io"
	"sync"
)

type KeyCounter struct {
	// TODO: not entirely sure why we would use it (debugging?), but may as well store it
	keyName string
	// TODO: is int enough?
	count    map[string]int
	reporter io.Writer
	mu       sync.Mutex
}

func NewKeyCounter(keyName string, reporter io.Writer) *KeyCounter {
	return &KeyCounter{
		keyName:  keyName,
		count:    make(map[string]int),
		reporter: reporter,
	}
}

func (kc *KeyCounter) Add(key string) {
	kc.mu.Lock()
	defer kc.mu.Unlock()

	kc.count[key]++
}

func (kc *KeyCounter) Flush() {
	kc.mu.Lock()
	defer kc.mu.Unlock()

	for k, v := range kc.count {
		fmt.Fprintf(kc.reporter, `"%s" - %d\n`, k, v)

		// TODO: maybe creating a new one from scratch is faster?
		delete(kc.count, k)
	}
}
