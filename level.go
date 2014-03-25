package seamdb

import (
	"github.com/syndtr/goleveldb/leveldb"
)

type SL struct {
	Db *leveldb.DB
}

func (self *SL) Open() (err error) {
	self.Db, err = leveldb.OpenFile("/Users/peter/Dev/tmp/level", nil)
	if err != nil {
		return err
	}
	return nil
}

func (self *SL) Put(key string, value []byte) (err error) {
	err = self.Db.Put([]byte(key), value, nil)
	return
}

func (self *SL) Get(key string) (value []byte, err error) {
	value, err = self.Db.Get([]byte(key), nil)
	return
}

func (self *SL) Del(key string) error {
	return self.Db.Delete([]byte(key), nil)
}

func (self *SL) Close() error {
	return self.Db.Close()
}
