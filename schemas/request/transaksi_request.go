package request

type CreateTransaksi struct { 
    NomorRekening   string  
    JenisTransaksi  string  
    JumlahTransaksi float64 
}

type UpdateTransaksi struct { 
	Id int 
    NomorRekening   string  
    JenisTransaksi  string  
    JumlahTransaksi float64 
}