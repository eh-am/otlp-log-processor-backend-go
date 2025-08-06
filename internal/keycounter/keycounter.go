package keycounter

import (
	"fmt"
	"io"
	"sort"
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

func (kc *KeyCounter) Add(value string) {
	kc.mu.Lock()
	defer kc.mu.Unlock()

	// TODO: when we converted to string using %v we end up getting <nil>s
	if len(value) <= 0 || value == "<nil>" {
		value = "unknown"
	}
	kc.count[value]++
}

func (kc *KeyCounter) Flush() {
	kc.mu.Lock()
	defer kc.mu.Unlock()

	// sort by alphabetical order
	var keys []string
	for key := range kc.count {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, key := range keys {
		value := kc.count[key]

		// TODO: handle this appropriately
		if key == "unknown" {
			fmt.Fprintf(kc.reporter, `unknown - %d%s`, value, "\n")
		} else {
			fmt.Fprintf(kc.reporter, `"%s" - %d%s`, key, value, "\n")
		}

		// TODO: maybe creating a new one from scratch is faster?
		delete(kc.count, key)
	}
}
