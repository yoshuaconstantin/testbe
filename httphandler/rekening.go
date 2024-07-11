package httphandler

import (
	"testbe/controller"
	"testbe/globalvariable/function"
	"net/http"
)


func RekeningCreateHandler(w http.ResponseWriter, r *http.Request) {
    function.SetHeaders(function.CustomWR{w, r})
    controller.CreateRekenings(function.CustomWR{w, r})
}

func RekeningReadHandler(w http.ResponseWriter, r *http.Request) {
    function.SetHeaders(function.CustomWR{w, r})
    controller.ReadRekenings(function.CustomWR{w, r})
}

func RekeningUpdateHandler(w http.ResponseWriter, r *http.Request) {
    function.SetHeaders(function.CustomWR{w, r})
    controller.UpdateRekenings(function.CustomWR{w, r})
}

func RekeningDeleteHandler(w http.ResponseWriter, r *http.Request) {
    function.SetHeaders(function.CustomWR{w, r})
    controller.DeleteRekenings(function.CustomWR{w, r})
}