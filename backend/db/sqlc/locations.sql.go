// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: locations.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createCity = `-- name: CreateCity :one
insert into locations_cities
values ($1, $2, $3)
returning id, name, state
`

type CreateCityParams struct {
	Column1 pgtype.UUID `json:"column_1"`
	Column2 pgtype.Text `json:"column_2"`
	Column3 pgtype.Text `json:"column_3"`
}

func (q *Queries) CreateCity(ctx context.Context, arg CreateCityParams) (LocationsCity, error) {
	row := q.db.QueryRow(ctx, createCity, arg.Column1, arg.Column2, arg.Column3)
	var i LocationsCity
	err := row.Scan(&i.ID, &i.Name, &i.State)
	return i, err
}

const getAllCities = `-- name: GetAllCities :many
select id, name, state from locations_cities
`

func (q *Queries) GetAllCities(ctx context.Context) ([]LocationsCity, error) {
	rows, err := q.db.Query(ctx, getAllCities)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []LocationsCity{}
	for rows.Next() {
		var i LocationsCity
		if err := rows.Scan(&i.ID, &i.Name, &i.State); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}