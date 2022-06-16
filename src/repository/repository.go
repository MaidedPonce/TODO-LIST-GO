package repository

import (
	"context"

	tasks "github.com/MaidedPonce/TODO-LIST-GO/src/models"
)

type Repository interface {
	ReadTaskID(id string) tasks.Task
	ReadTasks() tasks.List
	CreateTask(task tasks.Task) (tasks.Task, error)
	DeleteTask(id string)
}

var implementation Repository

func SetRepository(repo Repository) {
	implementation = repo
}

func ReadTaskID(ctx context.Context, id string) (*tasks.Task, error) {
	return implementation.ReadTasks(ctx, id)
}
