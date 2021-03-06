package repository

import (
	"context"

	tasks "github.com/MaidedPonce/TODO-LIST-GO/models"
)

type Repository interface {
	ReadTaskID(ctx context.Context, id string) (*tasks.Task, error)
	ReadTasks(ctx context.Context) ([]*tasks.Task, error)
	CreateTask(ctx context.Context, task *tasks.Task) error
	UpdateTask(ctx context.Context, task *tasks.Task) error
	DeleteTask(ctx context.Context, id string) error
	Close() error
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func ReadTaskID(ctx context.Context, id string) (*tasks.Task, error) {
	return implementation.ReadTaskID(ctx, id)
}

func ReadTasks(ctx context.Context) ([]*tasks.Task, error) {
	return implementation.ReadTasks(ctx)
}

func CreateTask(ctx context.Context, task *tasks.Task) error {
	return implementation.CreateTask(ctx, task)
}

func UpdateTask(ctx context.Context, task *tasks.Task) error {
	return implementation.UpdateTask(ctx, task)
}

func DeleteTask(ctx context.Context, id string) error {
	return implementation.DeleteTask(ctx, id)
}

func Close() error {
	return implementation.Close()
}
