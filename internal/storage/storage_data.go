package storage

import (
//	"fmt"
	"errors"
	"sync"
	"go18_api/internal/models"
)

type StorageData struct{ 
}

var mu sync.RWMutex
var data []models.Task = make([]models.Task, 0, 10)

// Конструктор
func NewStorageData() StorageData {
	d := StorageData{}
	return d;
}

func getIndex(id int) int {
	result := -1
	
	for i := 0; i < len(data); i++ {
		if data[i].ID == id {
			result = i
			break
		}
	}

	return result
}

// Реализуем интерфейс Storage
func (d StorageData) List() []models.Task {
	mu.RLock()
	defer mu.RUnlock()

	return data
}

func (d StorageData) Get(id int) (models.Task, bool) {
	mu.RLock()
	defer mu.RUnlock()

	index := getIndex(id)
	
	if index != -1 {
		return data[index], true
	} else {
		return models.Task{}, false
	}
}

func (d StorageData) Create(task models.Task) (models.Task, error) {
	mu.Lock()
	defer mu.Unlock()

	task.ID = len(data) + 1
	data = append(data, task)

	return task, nil
}

func (d StorageData) Delete(id int) error {
	mu.Lock()
	defer mu.Unlock()

	index := getIndex(id)
	
	if index == -1 {
		return errors.New("Элемент с таким ID не существует!")
	} else {
		data = append(data[:index], data[index+1:]...)
		return nil
	}
}

func (d StorageData) Update(id int, value models.Task) (models.Task, error) {
	mu.Lock()
	defer mu.Unlock()

	index := getIndex(id)
	
	if index == -1 {
		return models.Task{}, errors.New("Элемент с таким ID не существует!")
	} else {
		data[index].Title = value.Title
		data[index].Done = value.Done
		return value, nil
	}
}