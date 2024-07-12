package models

type Rekening struct {
    ID              int     
    NamaPemilik     string  
    NomorRekening   int  
    Saldo           int 
    TanggalPembuatan string 
}