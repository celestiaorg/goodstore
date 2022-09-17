package goodstore

import (
	"github.com/tidwall/btree"
)

type sortedCache struct {
	btree btree.Map[string, []byte]
}

func newSortedCache() *sortedCache {
	return &sortedCache{}
}

func (sc *sortedCache) get(key []byte) ([]byte, error) {
	value, found := sc.btree.Get(string(key))
	if !found {
		return nil, &InvalidKeyError{Key: key}
	}

	return value, nil
}

func (sc *sortedCache) set(key []byte, value []byte) {
	sc.btree.Set(string(key), value)
}

func (sc *sortedCache) delete(key []byte) error {
	sc.btree.Delete(string(key))
	return nil
}
