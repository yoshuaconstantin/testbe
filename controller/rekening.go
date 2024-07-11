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

func CreateRekenings(cr function.CustomWR) {

	var reqCreate request.CreateRekeningRequest

	if err := request.DynamicUnmarshalFromReader(cr.R.Body, &reqCreate); err != nil {
		function.Ehandler(cr, http.StatusBadRequest, err.Error(), "User facing this error: "+err.Error())
		return
	}

	if err := module.CreateRekening(reqCreate); err != nil {
		function.Ehandler(cr, http.StatusInternalServerError, err.Error(), "User facing this error: "+err.Error())
		return
	}

	var dataResponse response.GeneralResponseNoData
	dataResponse.Message = "ok"
	dataResponse.Status = http.StatusOK
	
	function.Rsp(cr.W).StatusCode(http.StatusOK).Log("berhasil").JSON(response.GeneralResponseNoData(dataResponse))
	
}

func ReadRekenings(cr function.CustomWR) {
	
	queryParams := cr.R.URL.Query()

	norekParam := queryParams.Get("nomorRekening")

	var responseRead models.Rekening

	responseRead, err := module.ReadRekening(norekParam)
	
	if err != nil {
		function.Ehandler(cr, http.StatusInternalServerError, err.Error(), "User facing this error: "+err.Error())
		return
	}
	
	function.Rsp(cr.W).StatusCode(http.StatusOK).Log("berhasil").JSON(responseRead)
	
}

func UpdateRekenings(cr function.CustomWR) {

	var reqUpdate request.UpdateRekeningRequest

	if err := request.DynamicUnmarshalFromReader(cr.R.Body, &reqUpdate); err != nil {
		function.Ehandler(cr, http.StatusBadRequest, err.Error(), "User facing this error: "+err.Error())
		return
	}

	if err := module.UpdateRekening(reqUpdate); err != nil {
		function.Ehandler(cr, http.StatusInternalServerError, err.Error(), "User facing this error: "+err.Error())
		return
	}

	var dataResponse response.GeneralResponseNoData
	dataResponse.Message = "ok"
	dataResponse.Status = http.StatusOK
	
	function.Rsp(cr.W).StatusCode(http.StatusOK).Log("berhasil").JSON(response.GeneralResponseNoData(dataResponse))
	
}

func DeleteRekenings(cr function.CustomWR) {
	
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
	


