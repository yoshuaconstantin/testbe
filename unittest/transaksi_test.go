package main

import (
	"database/sql"
	"testbe/module"
	"testbe/schemas/request"
	"testing"

	_ "github.com/lib/pq" // Import driver PostgreSQL
)

func TestTransaksi(t *testing.T) {
    // Konfigurasi koneksi ke database test
    connStr := "user=postgres password=123123123 dbname=testbecrud_test host=localhost sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        t.Fatal(err)
    }
    defer db.Close()

    // Buat rekening untuk keperluan testing transaksi
    newRekening := request.CreateRekeningRequest{
        NamaPemilik:     "Alice",
        NomorRekening:   1234567890,
        Saldo:           1000.0,
    
    }
    err = module.CreateRekening(newRekening)
    if err != nil {
        t.Fatal(err)
    }

    // Test CreateTransaksi
    newTransaksi := request.CreateTransaksi{
        NomorRekening:   "1234567890",
        JenisTransaksi:  "debit",
        JumlahTransaksi: 500.0,
    }
    err = module.CreateTransaksi(newTransaksi)
    if err != nil {
        t.Error("Error creating transaksi:", err)
    }

    // Test ReadTransaksi
    transaksiList, err := module.ReadTransaksi("1234567890") // Baca berdasarkan nomor rekening
    if err != nil {
        t.Error("Error reading transaksi:", err)
    } else if len(transaksiList) != 1 || transaksiList[0].JenisTransaksi != "debit" {
        t.Error("Unexpected transaksi data:", transaksiList)
    }

	updateTransaksi := request.UpdateTransaksi{
		Id: 1,
        NomorRekening:   "1234567890",
        JenisTransaksi:  "kredit",
        JumlahTransaksi: 501.0,
    }
    // Test UpdateTransaksi
    err = module.UpdateTransaksi(updateTransaksi)
    if err != nil {
        t.Error("Error updating transaksi:", err)
    }

    transaksiList, err = module.ReadTransaksi("1") // Baca berdasarkan ID
    if err != nil {
        t.Error("Error reading transaksi after update:", err)
    } else if len(transaksiList) != 1 || transaksiList[0].JenisTransaksi != "kredit" {
        t.Error("Transaksi not updated:", transaksiList)
    }

    // Test DeleteTransaksi
    err = module.DeleteTransaksi(updateTransaksi.Id)
    if err != nil {
        t.Error("Error deleting transaksi:", err)
    }

    transaksiList, err = module.ReadTransaksi("1")
    if err != nil {
        t.Error("Error reading transaksi after delete:", err)
    } else if len(transaksiList) != 0 {
        t.Error("Transaksi not deleted")
    }

    // Hapus rekening yang dibuat untuk testing
    err = module.DeleteRekening(updateTransaksi.Id)
    if err != nil {
        t.Fatal(err)
    }
}
