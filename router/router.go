package router

import (
	"github.com/gorilla/mux"
	"testbe/httphandler"
	//"Comnfo/websocketstruct"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	// User_Login API
	// router.HandleFunc("/api/users", controller.GetAllUsr).Methods("GET", "OPTIONS")
	// router.HandleFunc("/api/user/{id}", controller.GetSnglUsr).Methods("GET", "OPTIONS")

	router.HandleFunc("/api/transaksi/create", httphandler.TransaksiCreateHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/transaksi/read", httphandler.TransaksiReadHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/transaksi/update", httphandler.TransaksiUpdateHandler).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/transaksi/delete", httphandler.TransaksiDeleteHandler).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/rekening/create", httphandler.RekeningCreateHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/rekening/read", httphandler.RekeningReadHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/rekening/update", httphandler.RekeningUpdateHandler).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/rekening/delete", httphandler.RekeningDeleteHandler).Methods("DELETE", "OPTIONS")

	return router

}
