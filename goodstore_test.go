package goodstore

import (
	"crypto/sha256"
	"math/rand"
	"strconv"
	"testing"

	dbm "github.com/tendermint/tm-db"
)

func TestGoodStore(t *testing.T) {
	db, err := dbm.NewGoLevelDB("test", t.TempDir())
	if err != nil {
		t.Error(err)
	}

	_ = NewGoodStore(db, sha256.New())
}

func BenchmarkGet(b *testing.B) {
	db, _ := dbm.NewGoLevelDB("test", b.TempDir())
	store := NewGoodStore(db, sha256.New())

	for i := 0; i < 1000000; i++ {
		s := strconv.Itoa(i)
		_ = store.Set([]byte(s), []byte(s))
	}
	store.Commit()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s := strconv.Itoa(rand.Intn(1000000))
		_, err := store.Get([]byte(s))
		if err != nil {
			b.Error(err)
		}
	}
}
