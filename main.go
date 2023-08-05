package goose

import (
	"context"
	"database/sql"
)

type Goose struct {
	db *sql.DB
	ctx context.Context
}

func (goose *Goose) Insert(model interface{}) error {
	return goose.InsertSkipAndBind(model, []string{}, []string{})
}

func (goose *Goose) InsertSkipAndBind(model any, skip []string, bindingTemplates []string) error {
	query, params := goose.BuildInsertQueryWithParams()
	bindings, err := goose.BuildModelBindings(model, bindingTemplates)
	if err != nil {
		return err
	}

	return goose.db.QueryRowContext(goose.ctx, query, params).Scan(bindings...)
}

func (goose *Goose) BuildInsertQueryWithParams(model any, skip []string) (string, []string) {
	return "", []string{}
}

func (goose *Goose) BuildModelBindings(model any, bindingTemplates []string) ([]any, error) {
	return []any{}, nil
}

func New() *Goose {
	return &Goose{};
}
