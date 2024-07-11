package module

import (
	"testbe/config"
	"testbe/schemas/models"
	"testbe/schemas/request"
	"time"

	_ "github.com/lib/pq" // Import driver PostgreSQL
)

// Fungsi untuk membuat transaksi baru
func CreateTransaksi(transaksi request.CreateTransaksi) error {
    
    db := config.CreateConnection()

	defer db.Close()

    currentTime := time.Now()

    var tanggal = currentTime.Format("2006.01.02")

    _, err := db.Exec("INSERT INTO transaksi (NomorRekening, JenisTransaksi, JumlahTransaksi, TanggalTransaksi) VALUES (?, ?, ?, ?)",
        transaksi.NomorRekening, transaksi.JenisTransaksi, transaksi.JumlahTransaksi, tanggal)
    return err
}

// Fungsi untuk membaca informasi transaksi berdasarkan ID atau nomor rekening
func ReadTransaksi(identifier string) ([]models.Transaksi, error) {

    db := config.CreateConnection()

	defer db.Close()

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
func UpdateTransaksi(transaksi request.UpdateTransaksi) error {
    db := config.CreateConnection()

	defer db.Close()

    currentTime := time.Now()

    var tanggal = currentTime.Format("2006.01.02")

    _, err := db.Exec("UPDATE transaksi SET NomorRekening = ?, JenisTransaksi = ?, JumlahTransaksi = ?, TanggalTransaksi = ? WHERE ID = ?",
        transaksi.NomorRekening, transaksi.JenisTransaksi, transaksi.JumlahTransaksi, tanggal, transaksi.Id)
    return err
}

// Fungsi untuk menghapus transaksi berdasarkan ID
func DeleteTransaksi(id int) error {

    db := config.CreateConnection()

	defer db.Close()
    
    _, err := db.Exec("DELETE FROM transaksi WHERE ID = ?", id)
    return err
}
