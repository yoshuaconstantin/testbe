package models

type Transaksi struct {
    ID              int     
    NomorRekening   string  
    JenisTransaksi  string  
    JumlahTransaksi float64 
    TanggalTransaksi string 
}