package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreateTask(*Task) error
	GetTasks() ([]*Task, error)
	GetTask(string) (*Task, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := os.Getenv("CONN_STR")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil

}

func (s *PostgresStore) Init() error {
	return s.CreateTasksTable()
}

func (s *PostgresStore) CreateTasksTable() error {
	query := `create table if not exists tasks(
		id serial primary key,
		taskName varchar(254) not null,
		taskDetail varchar(254) not null,
		created_at timestamp,
		updated_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateTask(task *Task) error {
	query := `insert into tasks
		(taskname, taskdetail, created_at, updated_at)
		values ($1, $2, $3, $4)
	`

	_, err := s.db.Query(
		query,
		task.TaskName,
		task.TaskDetail,
		task.CreatedAt,
		task.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) GetTasks() ([]*Task, error) {
	rows, err := s.db.Query(`select * from tasks;`)
	if err != nil {
		return nil, err
	}

	tasks := []*Task{}
	for rows.Next() {
		task, err := ScanIntoTasks(rows)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s *PostgresStore) GetTask(id string) (*Task, error) {
	query := `select * from tasks where id = $1`
	rows, err := s.db.Query(
		query,
		id,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		task, err := ScanIntoTasks(rows)
		if err != nil {
			return nil, err
		}

		return task, nil
	}

	notFoundErr := fmt.Errorf("Task não encontrada")
	return nil, notFoundErr
}

func ScanIntoTasks(rows *sql.Rows) (*Task, error) {
	task := &Task{}
	err := rows.Scan(
		&task.ID,
		&task.TaskName,
		&task.TaskDetail,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	return task, err
}
