package goodstore

import db "github.com/tendermint/tm-db"

type nodeStore struct {
	wb *writeBuffer
	uc *unsortedCache
}

func newNodeStore(db db.DB) *nodeStore {
	uc := newUnsortedCache()
	return &nodeStore{
		wb: newWriteBuffer(db, uc),
		uc: uc,
	}
}

func (ns *nodeStore) Get(key []byte) ([]byte, error) {
	return ns.uc.get(key)
}

func (ns *nodeStore) Set(key []byte, value []byte) error {
	ns.wb.set(key, value)
	return nil
}

func (ns *nodeStore) Delete(key []byte) error {
	ns.wb.delete(key)
	return nil
}

func (ns *nodeStore) commit() error {
	return ns.wb.flush()
}
