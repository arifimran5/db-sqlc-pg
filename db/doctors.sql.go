// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: doctors.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createDoctor = `-- name: CreateDoctor :exec
INSERT INTO doctors (name, speciality) VALUES ($1, $2)
`

type CreateDoctorParams struct {
	Name       string
	Speciality pgtype.Text
}

func (q *Queries) CreateDoctor(ctx context.Context, arg CreateDoctorParams) error {
	_, err := q.db.Exec(ctx, createDoctor, arg.Name, arg.Speciality)
	return err
}

const deleteDoctor = `-- name: DeleteDoctor :exec
DELETE FROM doctors WHERE id = $1
`

func (q *Queries) DeleteDoctor(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteDoctor, id)
	return err
}

const getDoctor = `-- name: GetDoctor :one
SELECT id, name, speciality FROM doctors WHERE id = $1
`

func (q *Queries) GetDoctor(ctx context.Context, id int32) (Doctor, error) {
	row := q.db.QueryRow(ctx, getDoctor, id)
	var i Doctor
	err := row.Scan(&i.ID, &i.Name, &i.Speciality)
	return i, err
}

const getDoctors = `-- name: GetDoctors :many
SELECT id, name, speciality FROM doctors
`

func (q *Queries) GetDoctors(ctx context.Context) ([]Doctor, error) {
	rows, err := q.db.Query(ctx, getDoctors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Doctor
	for rows.Next() {
		var i Doctor
		if err := rows.Scan(&i.ID, &i.Name, &i.Speciality); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateDoctor = `-- name: UpdateDoctor :exec
UPDATE doctors SET name = $1, speciality = $2 WHERE id = $3
`

type UpdateDoctorParams struct {
	Name       string
	Speciality pgtype.Text
	ID         int32
}

func (q *Queries) UpdateDoctor(ctx context.Context, arg UpdateDoctorParams) error {
	_, err := q.db.Exec(ctx, updateDoctor, arg.Name, arg.Speciality, arg.ID)
	return err
}
