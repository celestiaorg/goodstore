package goodstore

import "sync"

type memoryMap struct {
	sync.RWMutex
	m map[string][]byte
}

type memoryMapItem struct {
	key   []byte
	value []byte
}

func newMemoryMap() *memoryMap {
	mm := &memoryMap{}
	mm.reset()
	return mm
}

func (mm *memoryMap) get(key []byte) ([]byte, error) {
	mm.RLock()
	defer mm.RUnlock()

	if value, ok := mm.m[string(key)]; ok {
		return value, nil
	}
	return nil, &InvalidKeyError{Key: key}
}

func (mm *memoryMap) set(key []byte, value []byte) {
	mm.Lock()
	defer mm.Unlock()
	mm.m[string(key)] = value
}

func (mm *memoryMap) delete(key []byte) error {
	mm.Lock()
	defer mm.Unlock()
	_, ok := mm.m[string(key)]
	if ok {
		delete(mm.m, string(key))
		return nil
	}
	return &InvalidKeyError{Key: key}
}

func (mm *memoryMap) reset() {
	mm.Lock()
	defer mm.Unlock()
	mm.m = make(map[string][]byte)
}

func (mm *memoryMap) iter() <-chan memoryMapItem {
	c := make(chan memoryMapItem)
	go func() {
		mm.RLock()
		defer mm.RUnlock()

		for k, v := range mm.m {
			c <- memoryMapItem{[]byte(k), v}
		}

		close(c)
	}()

	return c
}
