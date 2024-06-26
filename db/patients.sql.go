// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: patients.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createPatient = `-- name: CreatePatient :exec
INSERT INTO patients (name, age) VALUES ($1, $2)
`

type CreatePatientParams struct {
	Name string
	Age  pgtype.Int4
}

func (q *Queries) CreatePatient(ctx context.Context, arg CreatePatientParams) error {
	_, err := q.db.Exec(ctx, createPatient, arg.Name, arg.Age)
	return err
}

const deletePatient = `-- name: DeletePatient :exec
DELETE FROM patients WHERE id = $1
`

func (q *Queries) DeletePatient(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deletePatient, id)
	return err
}

const getPatient = `-- name: GetPatient :one
SELECT id, name, age FROM patients WHERE id = $1
`

func (q *Queries) GetPatient(ctx context.Context, id int32) (Patient, error) {
	row := q.db.QueryRow(ctx, getPatient, id)
	var i Patient
	err := row.Scan(&i.ID, &i.Name, &i.Age)
	return i, err
}

const getPatients = `-- name: GetPatients :many
SELECT id, name, age FROM patients
`

func (q *Queries) GetPatients(ctx context.Context) ([]Patient, error) {
	rows, err := q.db.Query(ctx, getPatients)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Patient
	for rows.Next() {
		var i Patient
		if err := rows.Scan(&i.ID, &i.Name, &i.Age); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePatient = `-- name: UpdatePatient :exec
UPDATE patients SET name = $1, age = $2 WHERE id = $3
`

type UpdatePatientParams struct {
	Name string
	Age  pgtype.Int4
	ID   int32
}

func (q *Queries) UpdatePatient(ctx context.Context, arg UpdatePatientParams) error {
	_, err := q.db.Exec(ctx, updatePatient, arg.Name, arg.Age, arg.ID)
	return err
}
