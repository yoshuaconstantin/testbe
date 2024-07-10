package models

type Rekening struct {
    ID              int     
    NamaPemilik     string  
    NomorRekening   string  
    Saldo           float64 
    TanggalPembuatan string 
}