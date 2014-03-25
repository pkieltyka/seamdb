package seamdb

import "github.com/boltdb/bolt"

type SB struct {
	Db *bolt.DB
}

func (self *SB) Open() (err error) {
	self.Db, err = bolt.Open("/Users/peter/Dev/tmp/bolt/main.db", 0666)
	if err != nil {
		return err
	}

	// make a default bucket..
	err = self.Db.Update(func(tx *bolt.Tx) error {
		return tx.CreateBucketIfNotExists("main")
	})

	return
}

func (self *SB) Put(bucket string, key string, value []byte) (err error) {
	err = self.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		return b.Put([]byte(key), value)
	})
	return
}

func (self *SB) Get(bucket string, key string) (value []byte, err error) {
	err = self.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		value = b.Get([]byte(key))
		return nil
	})
	return
}

func (self *SB) Del(bucket string, key string) error {
	return self.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		return b.Delete([]byte(key))
	})
}

func (self *SB) Update(fn func(*bolt.Tx) error) error {
	return self.Db.Update(fn)
}

func (self *SB) Close() error {
	return self.Db.Close()
}
