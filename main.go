package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/arifimran5/db-sqlc-pg/db"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://root:secret@localhost:5432/dbSqlcPg?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	queries := db.New(conn)

	//! create a doctor

	queries.CreateDoctor(ctx, db.CreateDoctorParams{
		Name: "Dr. Brother",
		Speciality: pgtype.Text{
			String: "Optician",
			Valid:  true,
		},
	})

	//! create a patient

	queries.CreatePatient(ctx, db.CreatePatientParams{
		Name: "Bob Smith",
		Age: pgtype.Int4{
			Int32: 53,
			Valid: true,
		},
	})

	//! create an appointment

	queries.CreateAppointment(ctx, db.CreateAppointmentParams{
		DoctorID:  2,
		PatientID: 2,
		AppointmentDate: pgtype.Date{
			Time:  time.Now().Add(time.Hour * 24 * 5),
			Valid: true,
		},
	})

	//! get all appointments by doctor id

	data, err := queries.GetAppointmentsByDoctor(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}

	for _, appointment := range data {
		fmt.Printf("Doctor: %s\n", appointment.DoctorName)
		fmt.Printf("Patient: %s\n", appointment.PatientName)
		fmt.Printf("Appointment Date: %v\n", appointment.AppointmentDate)

		fmt.Println("------")
	}
}
