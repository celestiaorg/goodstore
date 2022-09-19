package goodstore

import "sync"

type memoryMap struct {
	sync.Mutex
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
	if value, ok := mm.m[string(key)]; ok {
		return value, nil
	}
	return nil, &InvalidKeyError{Key: key}
}

func (mm *memoryMap) set(key []byte, value []byte) {
	mm.m[string(key)] = value
}

func (mm *memoryMap) delete(key []byte) error {
	_, ok := mm.m[string(key)]
	if ok {
		delete(mm.m, string(key))
		return nil
	}
	return &InvalidKeyError{Key: key}
}

func (mm *memoryMap) reset() {
	mm.m = make(map[string][]byte)
}

func (mm *memoryMap) iter() <-chan memoryMapItem {
	c := make(chan memoryMapItem)
	go func() {
		mm.Lock()
		defer mm.Unlock()

		for k, v := range mm.m {
			c <- memoryMapItem{[]byte(k), v}
		}

		close(c)
	}()

	return c
}

func (mm *memoryMap) copy() *memoryMap {
	mm.Lock()
	defer mm.Unlock()

	m := make(map[string][]byte, len(mm.m))
	for k, v := range mm.m {
		m[k] = v
	}

	return &memoryMap{
		m: m,
	}
}
