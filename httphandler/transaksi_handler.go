package httphandler

import (
	"testbe/controller"
	"testbe/globalvariable/function"
	"net/http"
)


func TransaksiCreateHandler(w http.ResponseWriter, r *http.Request) {
    function.SetHeaders(function.CustomWR{w, r})
    controller.LoginAccountTest(function.CustomWR{w, r})
}

func TransaksiReadHandler(w http.ResponseWriter, r *http.Request) {
    function.SetHeaders(function.CustomWR{w, r})
    controller.LoginAccountTest(function.CustomWR{w, r})
}

func TransaksiUpdateHandler(w http.ResponseWriter, r *http.Request) {
    function.SetHeaders(function.CustomWR{w, r})
    controller.LoginAccountTest(function.CustomWR{w, r})
}

func TransaksiDeleteHandler(w http.ResponseWriter, r *http.Request) {
    function.SetHeaders(function.CustomWR{w, r})
    controller.LoginAccountTest(function.CustomWR{w, r})
}