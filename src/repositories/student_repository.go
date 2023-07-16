package repository

import (
	"context"
	"database/sql"
	model "github-dbq/src/models"
	"log"
)

type studentRepository struct {
	db *sql.DB
}

func StudentRepositoryImpl(db *sql.DB) StudentRepository {
	return &studentRepository{db: db}
}

func (r *studentRepository) Get(ctx context.Context) (*[]model.Student, error) {

	query := "SELECT * FROM students"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []model.Student
	for rows.Next() {
		var student model.Student
		err := rows.Scan(
			&student.Id,
			&student.Name,
			&student.DateOfBirth,
			&student.CreatedAt,
			&student.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &students, nil
}

func (r *studentRepository) GetByID(id uint, ctx context.Context) (*model.Student, error) {

	query := "SELECT id, name, email FROM users WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var student model.Student
	err := row.Scan(
		&student.Id,
		&student.Name,
		&student.DateOfBirth,
		&student.CreatedAt,
		&student.DeletedAt,
	)
	if err != nil {
		log.Fatal(err)
	}
	return &student, nil

}
