package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Хранилище данных.
type Storage struct {
	db *pgxpool.Pool
}

// Конструктор, принимает строку подключения к БД.
func New(constr string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), constr)
	if err != nil {
		return nil, err
	}
	s := Storage{
		db: db,
	}
	return &s, nil
}

// Задача.
type Task struct {
	ID         int
	Opened     int64
	Closed     int64
	AuthorID   int
	AssignedID int
	Title      string
	Content    string
}

// Tasks возвращает список задач из БД.
func (s *Storage) Tasks(taskID, authorID int) ([]Task, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT 
			id,
			opened,
			closed,
			author_id,
			assigned_id,
			title,
			content
		FROM tasks
		WHERE
			($1 = 0 OR id = $1) AND
			($2 = 0 OR author_id = $2)
		ORDER BY id;
	`,
		taskID,
		authorID,
	)
	if err != nil {
		return nil, err
	}
	var tasks []Task
	// итерирование по результату выполнения запроса
	// и сканирование каждой строки в переменную
	for rows.Next() {
		var t Task
		err = rows.Scan(
			&t.ID,
			&t.Opened,
			&t.Closed,
			&t.AuthorID,
			&t.AssignedID,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		// добавление переменной в массив результатов
		tasks = append(tasks, t)

	}
	// ВАЖНО не забыть проверить rows.Err()
	return tasks, rows.Err()
}

// NewTask создаёт новую задачу и возвращает её id.
func (s *Storage) NewTask(t Task) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO tasks (title, content)
		VALUES ($1, $2) RETURNING id;
		`,
		t.Title,
		t.Content,
	).Scan(&id)
	return id, err
}

// TaskByLabel возвращает список задач из БД по метке.
func (s *Storage) TaskByLabel(label string) ([]Task, error) {
	rows, err := s.db.Query(context.Background(), `
			SELECT 
				id,
				opened,
				closed,
				author_id,
				assigned_id,
				title,
				content
			FROM tasks
			WHERE
				id=(SELECT task_id FROM tasks_labels
					WHERE 
						label_id=(SELECT id FROM labels WHERE name=$1)) ORDER BY id;
	`,
		label,
	)
	if err != nil {
		return nil, err
	}
	var tasks []Task
	// итерирование по результату выполнения запроса
	// и сканирование каждой строки в переменную
	for rows.Next() {
		var t Task
		err = rows.Scan(
			&t.ID,
			&t.Opened,
			&t.Closed,
			&t.AuthorID,
			&t.AssignedID,
			&t.Title,
			&t.Content,
		)
		if err != nil {
			return nil, err
		}
		// добавление переменной в массив результатов
		tasks = append(tasks, t)

	}
	// ВАЖНО не забыть проверить rows.Err()
	return tasks, rows.Err()
}

// Обновление задачи по id (Assigned, Title, Content).
func (s *Storage) UpdateTask(t Task) error {
	row := s.db.QueryRow(context.Background(), `
		UPDATE 
			tasks
		SET 
			assigned_id=$1,
			title=$2,
			content=$3
		WHERE 
			id=$4
			RETURNING id;
		`,
		t.AssignedID,
		t.Title,
		t.Content,
		t.ID,
	)
	err := row.Scan(
		&t.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

// Удаление задачи по id.
func (s *Storage) DeleteTask(t Task) error {
	_, err := s.db.Exec(context.Background(), `DELETE FROM tasks WHERE id=$1;`, t.ID)
	if err != nil {
		return err
	}
	return nil
}
