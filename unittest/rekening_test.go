package main

import (
	"database/sql"
	"strconv"

	"testbe/module"
	"testbe/schemas/request"
	"testing"

	_ "github.com/lib/pq" // Import driver PostgreSQL
)

func TestRekening(t *testing.T) {

    connStrInitial := "user=postgres password=123123123 dbname=postgres host=localhost sslmode=disable"
    dbInitial, err := sql.Open("postgres", connStrInitial)
    if err != nil {
        panic(err.Error())
    }
    defer dbInitial.Close()

    //Tutup koneksi awal setelah membuat database
    dbInitial.Close()

    // Konfigurasi koneksi baru ke database testbecrud
    connStr := "user=postgres password=123123123 dbname=testbecrud host=localhost sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
	

	// Test CreateRekening
	newRekening := request.CreateRekeningRequest{
		NamaPemilik:   "Alice",
		NomorRekening: 123123123,
		Saldo:         10000,
	}

	err = module.CreateRekening(newRekening)
	if err != nil {
		t.Error("Error creating rekening:", err)
	}

	// Test ReadRekening
	rekening, err := module.ReadRekening("123123123")
	if err != nil {
		t.Error("Error reading rekening:", err)
	} else if rekening.NamaPemilik != "Alice" {
		t.Error("Unexpected rekening data:", rekening)
	}

	var updaterek = request.UpdateRekeningRequest{
	NamaPemilik: "bob",
	ID: rekening.ID,
	NomorRekening: 123123123,
	}
	// Test UpdateRekening
	
	err = module.UpdateRekening(updaterek)
	if err != nil {
		t.Error("Error updating rekening:", err)
	}

	rekening, err = module.ReadRekening(strconv.Itoa(updaterek.NomorRekening))
	if err != nil {
		t.Error("Error reading rekening after update:", err)
	} else if rekening.NamaPemilik != "bob" {
		t.Error("Rekening not updated:", rekening)
	}

	// // Test DeleteRekening
	// err = module.DeleteRekening(rekening.ID)
	// if err != nil {
	// 	t.Error("Error deleting rekening:", err)
	// }

	// _, err = module.ReadRekening("123123123")
	// if err == nil {
	// 	t.Error("Rekening not deleted")
	// }

}
