package controller

import (
	"net/http"
	"strconv"
	"testbe/globalvariable/function"
	"testbe/module"
	"testbe/schemas/models"
	"testbe/schemas/request"
	"testbe/schemas/response"
)

func CreateTransaksis(cr function.CustomWR) {

	var reqCreate request.CreateTransaksi

	if err := request.DynamicUnmarshalFromReader(cr.R.Body, &reqCreate); err != nil {
		function.Ehandler(cr, http.StatusBadRequest, err.Error(), "User facing this error: "+err.Error())
		return
	}

	if err := module.CreateTransaksi(reqCreate); err != nil {
		function.Ehandler(cr, http.StatusInternalServerError, err.Error(), "User facing this error: "+err.Error())
		return
	}

	var dataResponse response.GeneralResponseNoData
	dataResponse.Message = "ok"
	dataResponse.Status = http.StatusOK
	
	function.Rsp(cr.W).StatusCode(http.StatusOK).Log("berhasil").JSON(response.GeneralResponseNoData(dataResponse))
	
}

func ReadTransaksis(cr function.CustomWR) {
	
	queryParams := cr.R.URL.Query()

	norekParam := queryParams.Get("nomorRekening")

	var responseRead []models.Transaksi

	responseRead, err := module.ReadTransaksi(norekParam)
	
	if err != nil {
		function.Ehandler(cr, http.StatusInternalServerError, err.Error(), "User facing this error: "+err.Error())
		return
	}
	
	function.Rsp(cr.W).StatusCode(http.StatusOK).Log("berhasil").JSON(responseRead)
	
}

func UpdateTransaksis(cr function.CustomWR) {

	var reqUpdate request.UpdateTransaksi

	if err := request.DynamicUnmarshalFromReader(cr.R.Body, &reqUpdate); err != nil {
		function.Ehandler(cr, http.StatusBadRequest, err.Error(), "User facing this error: "+err.Error())
		return
	}

	if err := module.UpdateTransaksi(reqUpdate); err != nil {
		function.Ehandler(cr, http.StatusInternalServerError, err.Error(), "User facing this error: "+err.Error())
		return
	}

	var dataResponse response.GeneralResponseNoData
	dataResponse.Message = "ok"
	dataResponse.Status = http.StatusOK
	
	function.Rsp(cr.W).StatusCode(http.StatusOK).Log("berhasil").JSON(response.GeneralResponseNoData(dataResponse))
	
}

func DeleteTransaksis(cr function.CustomWR) {
	
	queryParams := cr.R.URL.Query()

	idDelete := queryParams.Get("id")
	
	id, err := strconv.Atoi(idDelete)
	if err != nil {
		function.Ehandler(cr, http.StatusInternalServerError, err.Error(), "User facing this error: "+err.Error())
		return
	}
	

	if err := module.DeleteRekening(id); err != nil {
		function.Ehandler(cr, http.StatusInternalServerError, err.Error(), "User facing this error: "+err.Error())
		return
	}

	var dataResponse response.GeneralResponseNoData
	dataResponse.Message = "deleted"
	dataResponse.Status = http.StatusOK
	
	function.Rsp(cr.W).StatusCode(http.StatusOK).Log("berhasil").JSON(dataResponse)
	}
	


