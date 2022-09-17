package goodstore

import (
	"hash"

	"github.com/celestiaorg/smt"
	dbm "github.com/tendermint/tm-db"
)

type GoodStore struct {
	smt smt.SparseMerkleTree
	vs  *valueStore
	ns  *nodeStore
}

func NewGoodStore(db dbm.DB, hasher hash.Hash) *GoodStore {
	ns := newNodeStore(db)
	smt := smt.NewSMT(ns, hasher)
	gs := GoodStore{
		smt: smt,
		ns:  ns,
		vs:  newValueStore(db),
	}
	return &gs
}

func (gs *GoodStore) Get(key []byte) ([]byte, error) {
	return gs.vs.get(key)
}

func (gs *GoodStore) Set(key []byte, value []byte) error {
	var err error

	err = gs.vs.set(key, value)
	if err != nil {
		return err
	}

	err = gs.smt.Update(key, value)
	if err != nil {
		// TODO: if this happens, the valueStore gets corrupted. Should we undo
		// the valueStore, or treat it as fatal?
		return err
	}

	return nil
}

func (gs *GoodStore) Delete(key []byte) error {
	var err error

	err = gs.vs.delete(key)
	if err != nil {
		return err
	}

	err = gs.smt.Delete(key)
	if err != nil {
		// TODO: if this happens, the valueStore gets corrupted. Should we undo
		// the valueStore, or treat it as fatal?
		return err
	}

	return nil
}

func (gs *GoodStore) Commit() error {
	var err error

	err = gs.vs.commit()
	if err != nil {
		return err
	}

	gs.smt.Commit()

	err = gs.ns.commit()
	if err != nil {
		// TODO: this is fatal.
		return err
	}

	return nil
}
