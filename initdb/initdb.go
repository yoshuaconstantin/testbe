package initdb

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

func InitDB() {
    
    connStr := "user=your_user password=your_password dbname=postgres host=localhost sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    // Buat database
    _, err = db.Exec("CREATE DATABASE IF NOT EXISTS testbecrud")
    if err != nil {
        panic(err.Error())
    }

    // Pilih database yang baru dibuat
    _, err = db.Exec("USE testbecrud")
    if err != nil {
        panic(err.Error())
    }

    // Buat tabel "rekening"
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS rekening (
            ID SERIAL PRIMARY KEY,
            NamaPemilik VARCHAR(255) NOT NULL,
            NomorRekening VARCHAR(20) UNIQUE NOT NULL,
            Saldo REAL NOT NULL,
            TanggalPembuatan TIMESTAMP NOT NULL
        )
    `)
    if err != nil {
        panic(err.Error())
    }

    // Buat tabel "transaksi"
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS transaksi (
            ID SERIAL PRIMARY KEY,
            NomorRekening VARCHAR(20),
            JenisTransaksi VARCHAR(6) CHECK (JenisTransaksi IN ('debit', 'kredit')) NOT NULL,
            JumlahTransaksi REAL NOT NULL,
            TanggalTransaksi TIMESTAMP NOT NULL,
            FOREIGN KEY (NomorRekening) REFERENCES rekening(NomorRekening)
        )
    `)
    if err != nil {
        panic(err.Error())
    }

    fmt.Println("Database dan tabel berhasil dibuat!")
}
