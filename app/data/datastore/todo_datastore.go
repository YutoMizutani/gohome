package datastore

import (
	"errors"

	"github.com/YutoMizutani/gohome/app/domain/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type TodoDataStore struct {
}

func (dataStore *TodoDataStore) open() (db *gorm.DB, err error) {
	db, err = gorm.Open("sqlite3", "tmp/gorm.db")
	if err != nil {
		err = errors.New("failed to open DB: sqlite3")
		return nil, err
	}
	db.LogMode(true)
	return db, nil
}

func (dataStore *TodoDataStore) Migrate() error {
	db, err := dataStore.open()
	if err != nil {
		return err
	}
	defer db.Close()
	db.AutoMigrate(&entity.Todo{})
	return nil
}

func (dataStore *TodoDataStore) FindAll() (*entity.TodoList, error) {
	db, err := dataStore.open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var list entity.TodoList
	db.Find(&list)
	return &list, nil
}

func (dataStore *TodoDataStore) Create(entity *entity.Todo) error {
	db, err := dataStore.open()
	if err != nil {
		return err
	}
	defer db.Close()

	db.NewRecord(&entity)
	db.Create(&entity)
	return nil
}
