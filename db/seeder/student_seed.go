package seeder

import (
	"database/sql"
	"fmt"
	model "github-dbq/src/models"
	"log"
	"time"
)

func StudentSeed(db *sql.DB) {

	students := make([]model.Student, 3)

	for i := 0; i < 3; i++ {
		student := model.Student{
			Name:        fmt.Sprintf("aku ke-%d", i+1),
			Age:         15,
			DateOfBirth: time.Now().AddDate(1995, 12, 12),
			CreatedAt:   time.Now(),
			DeletedAt:   time.Now(),
		}
		students[i] = student
	}

	// Memulai transaksi
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// Buat prepared statement
	stmt, err := tx.Prepare("INSERT INTO students (id, name, age, date_of_birth, created_at, deleted_at) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Eksekusi prepared statement untuk setiap data
	for _, student := range students {
		_, err = stmt.Exec(student.Id, student.Name, student.Age, student.DateOfBirth, student.CreatedAt, student.DeletedAt)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Commit transaksi
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Bulk create berhasil")
}
