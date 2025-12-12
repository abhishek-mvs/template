package db

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

type DB struct {
	data map[string]map[uuid.UUID]interface{}
	mu   sync.RWMutex
}

func NewDB() *DB {
	return &DB{
		data: make(map[string]map[uuid.UUID]interface{}),
	}
}

// Save inserts or updates a record in a table.
func (db *DB) Save(table string,  value interface{}) (uuid.UUID, error) 	{
	db.mu.Lock()
	defer db.mu.Unlock()
	id := uuid.New()
	if _, ok := db.data[table]; !ok {
		db.data[table] = make(map[uuid.UUID]interface{})
	}
	db.data[table][id] = value
	return id, nil
}


func (db *DB) Update(table string, id uuid.UUID, value interface{}) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	if _, ok := db.data[table]; !ok {
		return errors.New("table not found")
	}
	if _, ok := db.data[table][id]; !ok {
		return errors.New("record not found")
	}
	db.data[table][id] = value
	return nil
}

// Get returns a record by table + id.
func (db *DB) Get(table string, id uuid.UUID) (interface{}, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	if tableData, ok := db.data[table]; ok {
		val, exists := tableData[id]
		return val, exists
	}
	return nil, false
}

// GetAll returns a copy of all data in a table.
func (db *DB) GetAll(table string) map[uuid.UUID]interface{} {
	db.mu.RLock()
	defer db.mu.RUnlock()

	result := make(map[uuid.UUID]interface{})
	if tableData, ok := db.data[table]; ok {
		for k, v := range tableData {
			result[k] = v
		}
	}
	return result
}

// Delete removes a record from a table.
func (db *DB) Delete(table string, id uuid.UUID) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if tableData, ok := db.data[table]; ok {
		delete(tableData, id)
	}
}