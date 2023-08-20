package repository

import (
	"context"
	"database/sql"
	"fmt"
	"restful-api/domain/entity"
	"restful-api/domain/repository"
	"restful-api/internal/repository/model"
	"restful-api/pkg"
	"time"

	"github.com/rocketlaunchr/dbq/v2"
)

type itemRepository struct {
	db     *sql.DB
	config *pkg.Config
}

func ItemRepositoryImpl(db *sql.DB, config *pkg.Config) repository.ItemRepository {
	return &itemRepository{db: db, config: config}
}

var itemModel model.Item

func (r *itemRepository) Get(ctx context.Context) ([]*entity.Item, error) {

	duration, err := time.ParseDuration(r.config.CustomTimout)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()

	stmt := fmt.Sprintf("SELECT * FROM %s", itemModel.GetTableName())
	opts := &dbq.Options{SingleResult: false, ConcreteStruct: itemModel, DecoderConfig: dbq.StdTimeConversionConfig()}
	data := dbq.MustQ(ctx, r.db, stmt, opts)

	result, err := model.MappingManyItem(data.([]*model.Item))
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *itemRepository) GetByID(id uint, ctx context.Context) (*entity.Item, error) {

	duration, err := time.ParseDuration(r.config.CustomTimout)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()

	stmt := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", itemModel.GetTableName())
	opts := &dbq.Options{SingleResult: true, ConcreteStruct: itemModel, DecoderConfig: dbq.StdTimeConversionConfig()}
	data := dbq.MustQ(ctx, r.db, stmt, opts, id)

	result, err := model.MappingSingleItem(data.(*model.Item))
	if err == nil {
		return nil, err
	}

	return result, nil

}

func (r *itemRepository) Update(item entity.ItemDTO, id uint, ctx context.Context) error {
	duration, err := time.ParseDuration(r.config.CustomTimout)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()

	stmt := fmt.Sprintf("UPDATE %s SET name = ? WHERE id = ?", itemModel.GetTableName())
	opts := &dbq.Options{SingleResult: true, ConcreteStruct: item, DecoderConfig: dbq.StdTimeConversionConfig()}

	_, err = dbq.E(ctx, r.db, stmt, opts, item, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *itemRepository) Delete(id uint, ctx context.Context) error {
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
	stmt := fmt.Sprintf("UPDATE %s SET deleted_at = ? WHERE id = ?", itemModel.GetTableName())
	opts := &dbq.Options{SingleResult: true, ConcreteStruct: itemModel, DecoderConfig: dbq.StdTimeConversionConfig()}

	_, err = dbq.E(ctx, r.db, stmt, opts, deleted, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *itemRepository) Create(item entity.ItemDTO, ctx context.Context) error {

	duration, err := time.ParseDuration(r.config.CustomTimout)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()

	stmt := fmt.Sprintf("INSERT INTO %s (name, age, date_of_birth, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", student.GetTableName())
	opts := &dbq.Options{SingleResult: true, ConcreteStruct: item, DecoderConfig: dbq.StdTimeConversionConfig()}

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
