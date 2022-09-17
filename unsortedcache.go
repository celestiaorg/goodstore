package goodstore

type unsortedCache struct {
	mm *memoryMap
}

func newUnsortedCache() *unsortedCache {
	return &unsortedCache{
		mm: newMemoryMap(),
	}
}

func (uc *unsortedCache) get(key []byte) ([]byte, error) {
	return uc.mm.get(key)
}

func (uc *unsortedCache) set(key []byte, value []byte) {
	uc.mm.set(key, value)
}

func (uc *unsortedCache) delete(key []byte) error {
	return uc.delete(key)
}
