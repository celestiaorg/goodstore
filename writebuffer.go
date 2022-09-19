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

	for k, v := range wb.writes.copy().m {
		wb.cache.set([]byte(k), v)
	}

	for k := range wb.deletes.copy().m{
		wb.cache.delete([]byte(k))
	}

	wb.writes.reset()
	wb.deletes.reset()
	return nil
}
