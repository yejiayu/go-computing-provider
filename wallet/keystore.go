package wallet

import (
	"encoding/json"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"os"
	"sync"
)

var diskKeyStore *DiskKeyStore
var once sync.Once
var lock sync.Mutex

type DiskKeyStore struct {
	db *leveldb.DB
}

func OpenOrInitKeystore(p string) (*DiskKeyStore, error) {
	lock.Lock()
	defer lock.Unlock()
	_, err := os.Stat(p)
	if err != nil {
		if !os.IsNotExist(err) {
			return diskKeyStore, err
		} else {
			if err = os.Mkdir(p, 0700); err != nil {
				return diskKeyStore, err
			}
		}
	}

	db, err := leveldb.OpenFile(p, nil)
	if err != nil {
		return diskKeyStore, err
	}
	diskKeyStore = &DiskKeyStore{db}

	return diskKeyStore, err
}

// List lists all the keys stored in the KeyStore
func (dks *DiskKeyStore) List() ([]string, error) {
	defer dks.db.Close()
	var keys []string
	iter := dks.db.NewIterator(nil, nil)
	for iter.Next() {
		addr := string(iter.Key())
		keys = append(keys, addr)
	}
	iter.Release()
	return keys, nil
}

// Get gets a key out of keystore and returns KeyInfo coresponding to named key
func (dks *DiskKeyStore) Get(name string) (KeyInfo, error) {
	defer dks.db.Close()
	value, err := dks.db.Get([]byte(name), nil)
	if err != nil {
		if err != nil {
			return KeyInfo{}, fmt.Errorf("decoding key '%s': %w", name, err)
		}
	}
	var res KeyInfo
	if err = json.Unmarshal(value, &res); err != nil {
		return KeyInfo{}, err
	}
	return res, nil
}

// Put saves key info under given name
func (dks *DiskKeyStore) Put(key string, info KeyInfo) error {
	defer dks.db.Close()
	bytes, _ := json.Marshal(info)
	err := dks.db.Put([]byte(key), bytes, nil)
	if err != nil {
		return fmt.Errorf("writing key '%s': %w", key, err)
	}
	return nil
}

func (dks *DiskKeyStore) Delete(key string) error {
	defer dks.db.Close()
	err := dks.db.Delete([]byte(key), nil)
	if err != nil {
		return fmt.Errorf("deleting key '%s': %w", key, err)
	}
	return nil
}

// KeyInfo is used for storing keys in KeyStore
type KeyInfo struct {
	PrivateKey string
}

// KeyStore is used for storing secret keys
type KeyStore interface {
	// List lists all the keys stored in the KeyStore
	List() ([]string, error)
	// Get gets a key out of keystore and returns KeyInfo corresponding to named key
	Get(string) (KeyInfo, error)
	// Put saves a key info under given name
	Put(string, KeyInfo) error
	// Delete removes a key from keystore
	Delete(string) error
}
