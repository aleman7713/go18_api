package storage

import "go18_api/internal/models"

type Storage interface {
	List() []models.Task
	Get(id int) (models.Task, bool)

	Create(task models.Task) (models.Task, error)
	Update(id int, value models.Task) (models.Task, error)
	Delete(id int) error
}