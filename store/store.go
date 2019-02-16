package store

import (
	"flag"
	"github.com/dgraph-io/badger"
)

var dataStr, valueDir string

var db *badger.DB

func init() {
	flag.StringVar(&dataStr, "dataDir", "/tmp", "数据存储目录")
	flag.StringVar(&valueDir, "valueDir", "/tmp", "值LOG存储目录")
	flag.Parse()

	opts := badger.DefaultOptions
	opts.Dir = dataStr
	opts.ValueDir = valueDir

	var err error
	db, err = badger.Open(opts)
	if err != nil {
		panic(err)
	}
}

func Close() {
	if db != nil {
		db.Close()
	}
}

func View(key string) ([]byte, error) {
	var value []byte
	err := db.View(func(txn *badger.Txn) error {
		it, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		// value, err = it.ValueCopy(nil)
		err = it.Value(func(val []byte) error {
			value = append([]byte{}, val...)
			return nil
		});
		return err
	})
	return value, err
}

func Store(key string, value []byte) error {
	return db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), value)
	})
}


func Remove(key string) error {
	return db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})
}