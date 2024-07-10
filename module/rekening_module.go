package module

import (
	"database/sql"
	"testbe/schemas/models"
)

// Fungsi untuk membuat rekening baru
func CreateRekening(db *sql.DB, rekening models.Rekening) error {
    _, err := db.Exec("INSERT INTO rekening (NamaPemilik, NomorRekening, Saldo, TanggalPembuatan) VALUES (?, ?, ?, ?)",
        rekening.NamaPemilik, rekening.NomorRekening, rekening.Saldo, rekening.TanggalPembuatan)
    return err
}

// Fungsi untuk membaca informasi rekening berdasarkan ID atau nomor rekening
func ReadRekening(db *sql.DB, identifier string) (models.Rekening, error) {
    var rekening models.Rekening
    err := db.QueryRow("SELECT * FROM rekening WHERE ID = ? OR NomorRekening = ?", identifier, identifier).Scan(
        &rekening.ID, &rekening.NamaPemilik, &rekening.NomorRekening, &rekening.Saldo, &rekening.TanggalPembuatan)
    return rekening, err
}

// Fungsi untuk memperbarui informasi rekening (kecuali saldo)
func UpdateRekening(db *sql.DB, rekening models.Rekening) error {
    _, err := db.Exec("UPDATE rekening SET NamaPemilik = ?, NomorRekening = ?, TanggalPembuatan = ? WHERE ID = ?",
        rekening.NamaPemilik, rekening.NomorRekening, rekening.TanggalPembuatan, rekening.ID)
    return err
}

// Fungsi untuk menghapus rekening berdasarkan ID
func DeleteRekening(db *sql.DB, id int) error {
    _, err := db.Exec("DELETE FROM rekening WHERE ID = ?", id)
    return err
}