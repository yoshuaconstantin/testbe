package module

import (
	"database/sql"
	"testbe/schemas/models"

	_ "github.com/lib/pq" // Import driver PostgreSQL
)

// Fungsi untuk membuat transaksi baru
func CreateTransaksi(db *sql.DB, transaksi models.Transaksi) error {
    _, err := db.Exec("INSERT INTO transaksi (NomorRekening, JenisTransaksi, JumlahTransaksi, TanggalTransaksi) VALUES (?, ?, ?, ?)",
        transaksi.NomorRekening, transaksi.JenisTransaksi, transaksi.JumlahTransaksi, transaksi.TanggalTransaksi)
    return err
}

// Fungsi untuk membaca informasi transaksi berdasarkan ID atau nomor rekening
func ReadTransaksi(db *sql.DB, identifier string) ([]models.Transaksi, error) {
    rows, err := db.Query("SELECT * FROM transaksi WHERE ID = ? OR NomorRekening = ?", identifier, identifier)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var transaksiList []models.Transaksi
    for rows.Next() {
        var transaksi models.Transaksi
        err := rows.Scan(&transaksi.ID, &transaksi.NomorRekening, &transaksi.JenisTransaksi, &transaksi.JumlahTransaksi, &transaksi.TanggalTransaksi)
        if err != nil {
            return nil, err
        }
        transaksiList = append(transaksiList, transaksi)
    }
    return transaksiList, nil
}

// Fungsi untuk memperbarui informasi transaksi
func UpdateTransaksi(db *sql.DB, transaksi models.Transaksi) error {
    _, err := db.Exec("UPDATE transaksi SET NomorRekening = ?, JenisTransaksi = ?, JumlahTransaksi = ?, TanggalTransaksi = ? WHERE ID = ?",
        transaksi.NomorRekening, transaksi.JenisTransaksi, transaksi.JumlahTransaksi, transaksi.TanggalTransaksi, transaksi.ID)
    return err
}

// Fungsi untuk menghapus transaksi berdasarkan ID
func DeleteTransaksi(db *sql.DB, id int) error {
    _, err := db.Exec("DELETE FROM transaksi WHERE ID = ?", id)
    return err
}
