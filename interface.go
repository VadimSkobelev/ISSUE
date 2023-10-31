package storageInterface

import postgres "Issue/pkg/storage"

type Interface interface {
	Tasks(int, int) ([]postgres.Task, error)
	TaskByLabel(string) ([]postgres.Task, error)
	UpdateTask(postgres.Task) error
	DeleteTask(postgres.Task) error
	NewTask(postgres.Task) (int, error)
}
