package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github-dbq/config"
	model "github-dbq/src/student/models"
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

	duration, err := time.ParseDuration(r.config.CustomTimout)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()

	stmt := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", student.GetTableName())
	opts := &dbq.Options{SingleResult: true, ConcreteStruct: student, DecoderConfig: dbq.StdTimeConversionConfig()}
	data := dbq.MustQ(ctx, r.db, stmt, opts, id)

	if data == nil {
		return nil, errors.New("students not found")
	}

	result := data.(*model.Student)

	return result, nil

}

func (r *studentRepository) Update(name string, id uint, ctx context.Context) error {
	duration, err := time.ParseDuration(r.config.CustomTimout)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()

	stmt := fmt.Sprintf("UPDATE %s SET name = ? WHERE id = ?", student.GetTableName())
	opts := &dbq.Options{SingleResult: true, ConcreteStruct: student, DecoderConfig: dbq.StdTimeConversionConfig()}

	_, err = dbq.E(ctx, r.db, stmt, opts, name, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *studentRepository) Delete(id uint, ctx context.Context) error {
	//soft delete
	duration, err := time.ParseDuration(r.config.CustomTimout)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()
	//hard delete
	//DELETE FROM %s WHERE id = ?
	deleted := time.Now()
	stmt := fmt.Sprintf("UPDATE %s SET deleted_at = ? WHERE id = ?", student.GetTableName())
	opts := &dbq.Options{SingleResult: true, ConcreteStruct: student, DecoderConfig: dbq.StdTimeConversionConfig()}

	_, err = dbq.E(ctx, r.db, stmt, opts, deleted, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *studentRepository) Create(student model.Student, ctx context.Context) error {

	duration, err := time.ParseDuration(r.config.CustomTimout)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()

	stmt := fmt.Sprintf("INSERT INTO %s (name, age, date_of_birth, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", student.GetTableName())
	opts := &dbq.Options{SingleResult: true, ConcreteStruct: student, DecoderConfig: dbq.StdTimeConversionConfig()}

	_, err = dbq.E(
		ctx,
		r.db,
		stmt,
		opts,
		student.Name,
		student.Age,
		student.DateOfBirth,
		student.CreatedAt,
		student.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
