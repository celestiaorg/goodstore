package goodstore

import db "github.com/tendermint/tm-db"

type valueStore struct {
	wb *writeBuffer
	sc *sortedCache
}

func newValueStore(db db.DB) *valueStore {
	sc := newSortedCache()
	return &valueStore{
		wb: newWriteBuffer(db, sc),
		sc: sc,
	}
}

func (vs *valueStore) get(key []byte) ([]byte, error) {
	return vs.sc.get(key)
}

func (vs *valueStore) set(key []byte, value []byte) error {
	vs.sc.set(key, value)
	return nil
}

func (vs *valueStore) delete(key []byte) error {
	return vs.sc.delete(key)
}

func (vs *valueStore) commit() error {
	return vs.wb.flush()
}
