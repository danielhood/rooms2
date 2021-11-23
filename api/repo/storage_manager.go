package repo

import (
	"git.mills.io/prologic/bitcask"
)

// StorageManager defines interface
type StorageManager interface {
	Load(key string) ([]byte, error)
	Store(key string, data []byte) error
}

type storageManager struct {
	db *bitcask.Bitcask
}

// NewStorageManager returns a new UserRepo instance
func NewStorageManager(db *bitcask.Bitcask) StorageManager {
	return &storageManager{
		db: db,
	}
}

// Store saves data to redis
func (s *storageManager) Store(key string, data []byte) error {
	return s.db.Put([]byte(key), data)
}

// Load retrieves data from redis
func (s *storageManager) Load(key string) ([]byte, error) {
	data, err := s.db.Get([]byte(key))

	if err != nil {
		return nil, err
	}

	return data, nil
}
