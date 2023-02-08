// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: place.sql

package db

import (
	"context"
	"database/sql"
)

const countPlaceReviews = `-- name: CountPlaceReviews :one
SELECT COUNT(id)
FROM reviews
WHERE place_id = $1
`

func (q *Queries) CountPlaceReviews(ctx context.Context, placeID int32) (int64, error) {
	row := q.db.QueryRowContext(ctx, countPlaceReviews, placeID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createPlace = `-- name: CreatePlace :one
INSERT INTO places (
  name,
  address,
  city,
  state,
  country,
  classification, 
  average_rate
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING id, name, address, city, state, country, average_rate, classification, created_at
`

type CreatePlaceParams struct {
	Name           string         `json:"name"`
	Address        sql.NullString `json:"address"`
	City           string         `json:"city"`
	State          string         `json:"state"`
	Country        sql.NullString `json:"country"`
	Classification sql.NullString `json:"classification"`
	AverageRate    int32          `json:"average_rate"`
}

func (q *Queries) CreatePlace(ctx context.Context, arg CreatePlaceParams) (Place, error) {
	row := q.db.QueryRowContext(ctx, createPlace,
		arg.Name,
		arg.Address,
		arg.City,
		arg.State,
		arg.Country,
		arg.Classification,
		arg.AverageRate,
	)
	var i Place
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Address,
		&i.City,
		&i.State,
		&i.Country,
		&i.AverageRate,
		&i.Classification,
		&i.CreatedAt,
	)
	return i, err
}

const deletePlace = `-- name: DeletePlace :exec
DELETE FROM places
WHERE id = $1
`

func (q *Queries) DeletePlace(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletePlace, id)
	return err
}

const getPlace = `-- name: GetPlace :one
SELECT id, name, address, city, state, country, average_rate, classification, created_at FROM places
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPlace(ctx context.Context, id int64) (Place, error) {
	row := q.db.QueryRowContext(ctx, getPlace, id)
	var i Place
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Address,
		&i.City,
		&i.State,
		&i.Country,
		&i.AverageRate,
		&i.Classification,
		&i.CreatedAt,
	)
	return i, err
}

const listPlaces = `-- name: ListPlaces :many
SELECT id, name, address, city, state, country, average_rate, classification, created_at FROM places
ORDER BY id
`

func (q *Queries) ListPlaces(ctx context.Context) ([]Place, error) {
	rows, err := q.db.QueryContext(ctx, listPlaces)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Place
	for rows.Next() {
		var i Place
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Address,
			&i.City,
			&i.State,
			&i.Country,
			&i.AverageRate,
			&i.Classification,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePlace = `-- name: UpdatePlace :exec
UPDATE places
  set name = $2,
  address = $3,
  city = $4,
  state = $5,
  country = $6,
  classification = $7
WHERE id = $1
`

type UpdatePlaceParams struct {
	ID             int64          `json:"id"`
	Name           string         `json:"name"`
	Address        sql.NullString `json:"address"`
	City           string         `json:"city"`
	State          string         `json:"state"`
	Country        sql.NullString `json:"country"`
	Classification sql.NullString `json:"classification"`
}

func (q *Queries) UpdatePlace(ctx context.Context, arg UpdatePlaceParams) error {
	_, err := q.db.ExecContext(ctx, updatePlace,
		arg.ID,
		arg.Name,
		arg.Address,
		arg.City,
		arg.State,
		arg.Country,
		arg.Classification,
	)
	return err
}