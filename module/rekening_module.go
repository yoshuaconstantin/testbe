package module

import (
	"database/sql"
	"testbe/schemas/models"
	"testbe/schemas/request"
	"time"
)

// Fungsi untuk membuat rekening baru
func CreateRekening(rekening request.CreateRekeningRequest) error {
    
    connStr := "user=postgres password=123123123 dbname=testbecrud host=localhost sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()


    currentTime := time.Now()

    var tanggal = currentTime.Format("2006.01.02")


    _, err = db.Exec("INSERT INTO rekening (namapemilik, nomorrekening, saldo, tanggalpembuatan) VALUES ($1, $2, $3, $4)",
        rekening.NamaPemilik, rekening.NomorRekening, rekening.Saldo, tanggal)
    return err
}

// Fungsi untuk membaca informasi rekening berdasarkan ID atau nomor rekening
func ReadRekening( identifier string) (models.Rekening, error) {

    connStr := "user=postgres password=123123123 dbname=testbecrud host=localhost sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    var rekening models.Rekening
    err = db.QueryRow("SELECT * FROM rekening WHERE nomorrekening = $1", identifier).Scan(
        &rekening.ID, &rekening.NamaPemilik, &rekening.NomorRekening, &rekening.Saldo, &rekening.TanggalPembuatan)
    return rekening, err
}

func ReadRekeningAll() (models.Rekening, error) {

    connStr := "user=postgres password=123123123 dbname=testbecrud host=localhost sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    var rekening models.Rekening
    err = db.QueryRow("SELECT * FROM rekening",).Scan(
        &rekening.ID, &rekening.NamaPemilik, &rekening.NomorRekening, &rekening.Saldo, &rekening.TanggalPembuatan)
    return rekening, err
}

// Fungsi untuk memperbarui informasi rekening (kecuali saldo)
func UpdateRekening( rekening request.UpdateRekeningRequest) error {

    connStr := "user=postgres password=123123123 dbname=testbecrud host=localhost sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    currentTime := time.Now()

    var tanggal = currentTime.Format("2006.01.02")

    _, err = db.Exec("UPDATE rekening SET namapemilik = $1, nomorrekening = $2, tanggalpembuatan = $3 WHERE id = $4",
        rekening.NamaPemilik, rekening.NomorRekening, tanggal, rekening.ID)
    return err
}

// Fungsi untuk menghapus rekening berdasarkan ID
func DeleteRekening( id int) error {

    connStr := "user=postgres password=123123123 dbname=testbecrud host=localhost sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    _, err = db.Exec("DELETE FROM rekening WHERE id = $1", id)
    return err
}

func DeleteRekening2( norek int) error {

    connStr := "user=postgres password=123123123 dbname=testbecrud host=localhost sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    _, err = db.Exec("DELETE FROM rekening WHERE nomorrekening = $1", norek)
    return err
}