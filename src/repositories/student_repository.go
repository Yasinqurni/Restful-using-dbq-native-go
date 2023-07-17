package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github-dbq/config"
	model "github-dbq/src/models"
	"time"

	"github.com/rocketlaunchr/dbq/v2"
)

type studentRepository struct {
	db     *sql.DB
	config *config.Config
}

func StudentRepositoryImpl(db *sql.DB, config *config.Config) StudentRepository {
	return &studentRepository{db: db, config: config}
}

var student model.Student

func (r *studentRepository) Get(ctx context.Context) ([]*model.Student, error) {

	duration, err := time.ParseDuration(r.config.CustomTimout)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()

	stmt := fmt.Sprintf("SELECT * FROM %s", student.GetTableName())
	opts := &dbq.Options{SingleResult: false, ConcreteStruct: student, DecoderConfig: dbq.StdTimeConversionConfig()}
	data := dbq.MustQ(ctx, r.db, stmt, opts)

	result := data.([]*model.Student)

	if len(result) == 0 || result == nil {
		return nil, errors.New("students not found")
	}

	return result, nil
}

func (r *studentRepository) GetByID(id uint, ctx context.Context) (*model.Student, error) {

	query := "SELECT * FROM students WHERE id = ?"
	row := r.db.QueryRowContext(ctx, query, id)

	var student model.Student
	err := row.Scan(
		&student.Id,
		&student.Name,
		&student.Age,
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
