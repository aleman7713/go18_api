package storage

import (
	"errors"
	"sync"
	"go18_api/internal/models"
)

type StorageData struct{ 
}

var mu sync.RWMutex
var data []models.Task = make([]models.Task, 0, 10)
var seq_task_id int = 0;

// Конструктор
func NewStorageData() Storage {
	d := StorageData{}
	return &d;
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

	data2 := make([]models.Task, len(data)) 
	copy(data2, data)

	return data2
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

	seq_task_id++
	task.ID = seq_task_id
	data = append(data, task)

	return data[len(data) - 1], nil
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
		return data[index], nil
	}
}