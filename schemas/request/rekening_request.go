package request

type CreateRekeningRequest struct {
    NamaPemilik   string  `json:"namapemilik"`
    NomorRekening int `json:"nomorrekening"`
    Saldo  int  `json:"saldo"`
}

type UpdateRekeningRequest struct {
    ID              int  `json:"id"`
    NamaPemilik   string  `json:"namapemilik"`
    NomorRekening int `json:"nomorrekening"`

}