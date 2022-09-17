package goodstore

import db "github.com/tendermint/tm-db"

type writeBuffer struct {
	writes  *memoryMap
	deletes *memoryMap
	db      db.DB
	cache   cacheForWriteBuffer
}

type cacheForWriteBuffer interface {
	set([]byte, []byte)
	delete(key []byte) error
}

func newWriteBuffer(db db.DB, cache cacheForWriteBuffer) *writeBuffer {
	return &writeBuffer{
		writes:  newMemoryMap(),
		deletes: newMemoryMap(),
		db:      db,
		cache:   cache,
	}
}

func (wb *writeBuffer) set(key []byte, value []byte) {
	wb.writes.set(key, value)
}

func (wb *writeBuffer) delete(key []byte) {
	wb.deletes.set(key, nil)
}

func (wb *writeBuffer) flush() error {
	// TODO: flush to db

	for item := range wb.writes.iter() {
		wb.cache.set(item.key, item.value)
	}

	for item := range wb.deletes.iter() {
		wb.cache.delete(item.key)
	}

	wb.writes.reset()
	wb.deletes.reset()
	return nil
}
