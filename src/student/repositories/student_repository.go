package repository

import (
	"context"
	"database/sql"
	"fmt"
	model "github-dbq/src/student/models"
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

	query := "SELECT * FROM students WHERE id = ?"
	row := r.db.QueryRowContext(ctx, query, id)

	var student model.Student
	err := row.Scan(
		&student.Id,
		&student.Name,
		&student.DateOfBirth,
		&student.CreatedAt,
		&student.DeletedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("student not found")
		}
		return nil, err
	}
	return &student, nil

}
