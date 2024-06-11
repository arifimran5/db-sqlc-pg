-- name: CreateDoctor :exec
INSERT INTO doctors (name, speciality) VALUES ($1, $2);

-- name: GetDoctor :one
SELECT * FROM doctors WHERE id = $1;

-- name: GetDoctors :many
SELECT * FROM doctors;

-- name: UpdateDoctor :exec
UPDATE doctors SET name = $1, speciality = $2 WHERE id = $3;

-- name: DeleteDoctor :exec
DELETE FROM doctors WHERE id = $1;