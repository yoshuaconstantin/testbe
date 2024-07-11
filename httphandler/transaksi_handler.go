package httphandler

import (
	"testbe/controller"
	"testbe/globalvariable/function"
	"net/http"
)


func TransaksiCreateHandler(w http.ResponseWriter, r *http.Request) {
    function.SetHeaders(function.CustomWR{w, r})
    controller.CreateTransaksis(function.CustomWR{w, r})
}

func TransaksiReadHandler(w http.ResponseWriter, r *http.Request) {
    function.SetHeaders(function.CustomWR{w, r})
    controller.ReadTransaksis(function.CustomWR{w, r})
}

func TransaksiUpdateHandler(w http.ResponseWriter, r *http.Request) {
    function.SetHeaders(function.CustomWR{w, r})
    controller.UpdateTransaksis(function.CustomWR{w, r})
}

func TransaksiDeleteHandler(w http.ResponseWriter, r *http.Request) {
    function.SetHeaders(function.CustomWR{w, r})
    controller.DeleteTransaksis(function.CustomWR{w, r})
}