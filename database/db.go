package db

import (
	"context"
	"database/sql"
	"log"

	tasks "github.com/MaidedPonce/TODO-LIST-GO/models"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewDBRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	return &PostgresRepository{db}, nil
}

func (repo *PostgresRepository) CreateTask(ctx context.Context, task *tasks.Task) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO tasks (id, text) VALUES ($1, $2)", task.Id, task.Text)

	return err
}

func (repo *PostgresRepository) ReadTaskID(ctx context.Context, id string) (*tasks.Task, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, task FROM tasks WHERE id = $1", id)

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var t = tasks.Task{}

	for rows.Next() {
		if err = rows.Scan(&t.Id, &t.Text); err == nil {
			return &t, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &t, nil
}

func (repo *PostgresRepository) ReadTasks(ctx context.Context) ([]*tasks.Task, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT * FROM tasks")

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var alltasks []*tasks.Task

	for rows.Next() {
		var t = tasks.Task{}
		if err = rows.Scan(&t.Id, &t.Text); err == nil {
			alltasks = append(alltasks, &t)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return alltasks, nil
}

func (repo *PostgresRepository) Close() error {
	return repo.db.Close()
}
