package module

import (
	"database/sql"
	"testbe/schemas/models"
	"testbe/schemas/request"
	"time"

	_ "github.com/lib/pq" // Import driver PostgreSQL
)

// Fungsi untuk membuat transaksi baru
func CreateTransaksi(transaksi request.CreateTransaksi) error {
    
    connStr := "user=postgres password=123123123 dbname=testbecrud host=localhost sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    currentTime := time.Now()

    var tanggal = currentTime.Format("2006.01.02")

    _, err = db.Exec("INSERT INTO transaksi (nomorrekening, jenistransaksi, jumlahtransaksi, tanggaltransaksi) VALUES ($1, $2, $3, $4)",
        transaksi.NomorRekening, transaksi.JenisTransaksi, transaksi.JumlahTransaksi, tanggal)
    return err
}

// Fungsi untuk membaca informasi transaksi berdasarkan ID atau nomor rekening
func ReadTransaksi(identifier string) ([]models.Transaksi, error) {

    connStr := "user=postgres password=123123123 dbname=testbecrud host=localhost sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    rows, err := db.Query("SELECT * FROM transaksi WHERE nomorrekening = $1", identifier)
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
    connStr := "user=postgres password=123123123 dbname=testbecrud host=localhost sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    currentTime := time.Now()

    var tanggal = currentTime.Format("2006.01.02")

    _, err = db.Exec("UPDATE transaksi SET nomorrekening = $1, jenistransaksi = $2, jumlahtransaksi = $3, tanggaltransaksi = $4 WHERE id = $5",
        transaksi.NomorRekening, transaksi.JenisTransaksi, transaksi.JumlahTransaksi, tanggal, transaksi.Id)
    return err
}

// Fungsi untuk menghapus transaksi berdasarkan ID
func DeleteTransaksi(id int) error {
    connStr := "user=postgres password=123123123 dbname=testbecrud host=localhost sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
    
    _, err = db.Exec("DELETE FROM transaksi WHERE id = $1", id)
    return err
}
