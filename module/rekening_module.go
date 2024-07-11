package module

import (
	"testbe/config"
	"testbe/schemas/models"
	"testbe/schemas/request"
	"time"
)

// Fungsi untuk membuat rekening baru
func CreateRekening(rekening request.CreateRekeningRequest) error {
    db := config.CreateConnection()

	defer db.Close()


    currentTime := time.Now()

    var tanggal = currentTime.Format("2006.01.02")


    _, err := db.Exec("INSERT INTO rekening (NamaPemilik, NomorRekening, Saldo, TanggalPembuatan) VALUES (?, ?, ?, ?)",
        rekening.NamaPemilik, rekening.NomorRekening, rekening.Saldo, tanggal)
    return err
}

// Fungsi untuk membaca informasi rekening berdasarkan ID atau nomor rekening
func ReadRekening( identifier string) (models.Rekening, error) {

    db := config.CreateConnection()

	defer db.Close()

    var rekening models.Rekening
    err := db.QueryRow("SELECT * FROM rekening WHERE ID = ? OR NomorRekening = ?", identifier, identifier).Scan(
        &rekening.ID, &rekening.NamaPemilik, &rekening.NomorRekening, &rekening.Saldo, &rekening.TanggalPembuatan)
    return rekening, err
}

// Fungsi untuk memperbarui informasi rekening (kecuali saldo)
func UpdateRekening( rekening request.UpdateRekeningRequest) error {

    db := config.CreateConnection()

	defer db.Close()

    currentTime := time.Now()

    var tanggal = currentTime.Format("2006.01.02")

    _, err := db.Exec("UPDATE rekening SET NamaPemilik = ?, NomorRekening = ?, TanggalPembuatan = ? WHERE ID = ?",
        rekening.NamaPemilik, rekening.NomorRekening, tanggal, rekening.ID)
    return err
}

// Fungsi untuk menghapus rekening berdasarkan ID
func DeleteRekening( id int) error {

    db := config.CreateConnection()

	defer db.Close()

    _, err := db.Exec("DELETE FROM rekening WHERE ID = ?", id)
    return err
}