package models

import (
	"context"
	"database/sql"
	"fmt"
)

type Food struct {
	Name    string
	Cuisine string
}
type AddFoodInput struct {
	Name    string
	Cuisine string
}

type FoodDbService struct {
	*sql.DB
}

func GetFood(id string) (Food, error) {
	return Food{}, nil
}

func AddFood(ctx context.Context, input AddFoodInput) (Food, error) {

	db, _ := ctx.Value("dbKey").(*sql.DB)

	fdb := FoodDbService{DB: db}

	result, _ := fdb.AddFood(ctx, input)

	return result, nil
}

func (db *FoodDbService) AddFood(ctx context.Context, input AddFoodInput) (Food, error) {

	f := Food{
		Name:    input.Name,
		Cuisine: input.Cuisine,
	}

	const q = `INSERT INTO food (name, cuisine) VALUES ( $1, $2) RETURNING name`

	var name string
	row := db.QueryRowContext(ctx, q, f.Name, f.Cuisine)
	err := row.Scan(&name)
	if err != nil {
		return Food{}, fmt.Errorf("inserting user %w", err)
	}

	return f, nil

}
