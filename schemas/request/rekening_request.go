package request

type CreateRekeningRequest struct {
    NamaPemilik   string  `json:"namaPemilik"`
    NomorRekening int `json:"nomorRekening"`
    Saldo  float64  `json:"saldo"`
}

type UpdateRekeningRequest struct {
    ID              int  `json:"id"`
    NamaPemilik   string  `json:"namaPemilik"`
    NomorRekening int `json:"nomorRekening"`

}