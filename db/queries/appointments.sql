-- name: CreateAppointment :exec
INSERT INTO
    appointments (
        doctor_id,
        patient_id,
        appointment_date
    )
VALUES ($1, $2, $3);

-- name: GetAppointmentsByDoctor :many
SELECT
    d.id AS doctor_id,
    d.name AS doctor_name,
    p.id AS patitent_id,
    p.name AS patient_name,
    p.age AS patient_age,
    a.appointment_date
FROM
    doctors d
    JOIN appointments a ON d.id = a.doctor_id
    JOIN patients p ON a.patient_id = p.id
WHERE
    d.id = $1
ORDER BY a.appointment_date;

-- name: GetAppointmentsByPatient :many
SELECT
    d.id AS doctor_id,
    d.name AS doctor_name,
    p.id AS patitent_id,
    p.name AS patient_name,
    p.age AS patient_age,
    a.appointment_date
FROM
    doctors d
    JOIN appointments a ON d.id = a.doctor_id
    JOIN patients p ON a.patient_id = p.id
WHERE
    p.id = $1
ORDER BY a.appointment_date;