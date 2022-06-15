package taskrepo

import (
	"fmt"

	db "github.com/MaidedPonce/TODO-LIST-GO/src/database"
	tasks "github.com/MaidedPonce/TODO-LIST-GO/src/models"
)

type TaskRepo struct {
	db *db.GetDataBase
}

func NewTaskRepo(db *db.GetDataBase) *TaskRepo {
	return &TaskRepo{
		db: db,
	}
}

func (repo *TaskRepo) GetTasks() {

	var list tasks.List
	fmt.Println(list)
}
