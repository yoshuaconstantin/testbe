package main

import (
	"database/sql"
	"fmt"
	"os"
	"testbe/module"
	"testbe/schemas/request"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Import driver PostgreSQL
)

func TestRekening(t *testing.T) {

	err := godotenv.Load(".env") // Atau gunakan jalur absolut jika perlu
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}
	// Konfigurasi koneksi ke database test (misalnya menggunakan database testbecrud_test)
	// connStr := "user=postgres password=123123123 dbname=testbecrud host=localhost sslmode=disable"
	// db, err := sql.Open("postgres", connStr)
	db, err := sql.Open("postgres", "user=posgres password=123123123 dbname=testbecrud host=localhost sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// Test CreateRekening
	newRekening := request.CreateRekeningRequest{
		NamaPemilik:   "Alice",
		NomorRekening: 123123123,
		Saldo:         1000.0,
	}

	newUpdateRekening := request.UpdateRekeningRequest{
		ID:            1,
		NamaPemilik:   "Alice2",
		NomorRekening: 123123123,
	}

	err = module.CreateRekening(newRekening)
	if err != nil {
		t.Error("Error creating rekening:", err)
	}

	// Test ReadRekening
	rekening, err := module.ReadRekening("1234567890")
	if err != nil {
		t.Error("Error reading rekening:", err)
	} else if rekening.NamaPemilik != "Alice" {
		t.Error("Unexpected rekening data:", rekening)
	}

	// Test UpdateRekening
	rekening.NamaPemilik = "Bob"
	err = module.UpdateRekening(newUpdateRekening)
	if err != nil {
		t.Error("Error updating rekening:", err)
	}

	rekening, err = module.ReadRekening("123123123")
	if err != nil {
		t.Error("Error reading rekening after update:", err)
	} else if rekening.NamaPemilik != "Bob" {
		t.Error("Rekening not updated:", rekening)
	}

	// Test DeleteRekening
	err = module.DeleteRekening(rekening.ID)
	if err != nil {
		t.Error("Error deleting rekening:", err)
	}

	_, err = module.ReadRekening("123123123")
	if err == nil {
		t.Error("Rekening not deleted")
	}

}
