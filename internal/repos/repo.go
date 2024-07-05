package repos

import (
	"database/sql"
	"fmt"
	"task/internal/models"

	sq "github.com/Masterminds/squirrel"
)

type TaskRepo struct {
	DB *sql.DB
}

func NewTaskRepo(db *sql.DB) *TaskRepo {
	return &TaskRepo{DB: db}
}

func (t *TaskRepo) GetTasksFromDatabase() ([]models.Task, error) {
	var tasks = []models.Task{
		{
			Id:          1,
			Title:       "asdasdc",
			Description: "ascasdc",
			Status:      "casdcasdc",
		},
		{
			Id:          2,
			Title:       "asdasdc",
			Description: "asdc",
			Status:      "casdcasdc",
		},
	}

	return tasks, nil
}

func (t *TaskRepo) GetTaskByIdFromDatabase(id int) (*models.Task, error) {

	query, args, err := sq.
		Select("*").From("tasks").Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, fmt.Errorf("error: query select %v", err)
	}

	var task models.Task
	err = t.DB.QueryRow(query, args...).Scan(&task.Id, &task.Title, &task.Description, &task.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no task found")
		}
		return nil, fmt.Errorf("error: db queryrow  %v", err)
	}

	return &task, nil

}

func (t *TaskRepo) CreateTaskInDatabase(task models.Task) (string, error) {
	task.Status = "proccess"
	// query, args, err := sq.
	// 	Insert("tasks").Columns("title", "description", "status").Values(task.Title, task.Description, task.Status).PlaceholderFormat(sq.Dollar).ToSql()
	// if err != nil {
	// 	return "", fmt.Errorf("error: query insert %v", err)
	// }

	// _, err = t.DB.Exec(query, args...)
	// if err != nil {
	// 	return "", fmt.Errorf("error: db exec  %v", err)
	// }

	return "Created Succesfully", nil
}

func (t *TaskRepo) UpdateTaskInDatabase(task models.Task) (*models.Task, error) {
	query, args, err := sq.
		Update("tasks").SetMap(map[string]interface{}{
		"title":       task.Title,
		"description": task.Description,
		"status":      task.Status,
	}).Where(sq.Eq{"id": task.Id}).Suffix("RETURNING id, title, description, status").PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, fmt.Errorf("error: query update %v", err)
	}

	var newTask models.Task
	err = t.DB.QueryRow(query, args...).Scan(&newTask.Id, &newTask.Title, &newTask.Description, &newTask.Status)
	if err != nil {
		return nil, fmt.Errorf("error: db queryrow %v", err)
	}

	return &newTask, nil

}

func (t *TaskRepo) DeleteTaskFromDatabase(id int) (string, error) {

	query, args, err := sq.
		Delete("tasks").Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return "", fmt.Errorf("error: query delete %v", err)
	}

	n, err := t.DB.Exec(query, args...)
	if err != nil {
		return "", fmt.Errorf("error: db exec %v", err)
	}

	if result, _ := n.RowsAffected(); result == 0 {
		return "No Task Exists With this ID", nil
	}

	return "Deleted Succesfully", nil
}
