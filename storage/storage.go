package storage

import (
	"time"

	"github.com/boltdb/bolt"

	"github.com/codecoins/codecoins/config"
	"github.com/codecoins/codecoins/log"
)

type Storage struct {
	db *bolt.DB
}

type Bucket struct {
	b string
	s *Storage
}

func Load() *Storage {
	dbPath, err := config.GetString("storage.path")
	log.DieFatal(err)
	db, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 5 * time.Second})
	log.DieFatal(err)
	return &Storage{db}
}

func (s *Storage) Bucket(b string) *Bucket {
	return &Bucket{
		b, s,
	}
}

func (b *Bucket) Get(k string) []byte {
	var out []byte
	err := b.s.db.View(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(b.b))
		if err == nil {
			out = bucket.Get([]byte(k))
		}
		return err
	})
	log.PrintError(err)
	return out
}

func (b *Bucket) Write(k string, v []byte) error {
	err := b.s.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(b.b))
		if err == nil {
			bucket.Put([]byte(k), v)
		}
		return err
	})
	return log.PrintError(err)
}

func (b *Bucket) WriteString(b, k, v string) error {
	return b.Write(k, []byte(v))
}
