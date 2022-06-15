package repository

import tasks "github.com/MaidedPonce/TODO-LIST-GO/src/models"

type Repository interface {
	ReadTaskID(id string) tasks.Task
	ReadTasks() tasks.List
	CreateTask(task tasks.Task) (tasks.Task, error)
	DeleteTask(id string)
}
