-- name: CreatePatient :exec
INSERT INTO patients (name, age) VALUES ($1, $2);

-- name: GetPatient :one
SELECT * FROM patients WHERE id = $1;

-- name: GetPatients :many
SELECT * FROM patients;

-- name: UpdatePatient :exec
UPDATE patients SET name = $1, age = $2 WHERE id = $3;

-- name: DeletePatient :exec
DELETE FROM patients WHERE id = $1;