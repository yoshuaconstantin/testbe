package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {

    // // Konfigurasi koneksi awal ke database postgres
    // connStrInitial := "user=postgres password=123123123 dbname=postgres host=localhost sslmode=disable"
    // dbInitial, err := sql.Open("postgres", connStrInitial)
    // if err != nil {
    //     panic(err.Error())
    // }
    // defer dbInitial.Close()

    // // Buat database
    // _, err = dbInitial.Exec("CREATE DATABASE testbecrud")
    // if err != nil {
    //     panic(err.Error())
    // }

    // Tutup koneksi awal setelah membuat database
    // dbInitial.Close()

    // Konfigurasi koneksi baru ke database testbecrud
    connStr := "user=postgres password=123123123 dbname=testbecrud host=localhost sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

	// Buat tabel "rekening"
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS rekening (
            ID SERIAL PRIMARY KEY,
            NamaPemilik VARCHAR(255) NOT NULL,
            NomorRekening VARCHAR(50) UNIQUE NOT NULL,
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
