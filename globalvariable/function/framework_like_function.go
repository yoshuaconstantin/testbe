package function

import (
    "encoding/json"
    "net/http"
	Lg "testbe/logging"
)

// customWR holds the http.ResponseWriter and *http.Request together.
type CustomWR struct {
    W http.ResponseWriter
    R *http.Request
}

// StatusCode sets the response status code and returns the responseWriter for chaining.
func (rw *CustomWR) StatusCode(statusCode int) *CustomWR {
    rw.W.WriteHeader(statusCode)
    return rw
}

func (rw *CustomWR) Log(message string) *CustomWR {
	Lg.Info(message)
    return rw
}

func (rw *CustomWR) LogE(message string) *CustomWR {
	Lg.Error(message,"")
    return rw
}

// Data encodes the given response data as JSON and sends it.
func (rw *CustomWR) JSON(response interface{}) error {
    return rw.Response(response)
}

// Response is a helper method for encoding and sending a JSON response.
func (rw *CustomWR) Response(response interface{}) error {
    return json.NewEncoder(rw.W).Encode(response)
}

// NewResponseWriter creates a new responseWriter for a given http.ResponseWriter.
func Rsp(w http.ResponseWriter) *CustomWR {
    return &CustomWR{W: w}
}

// setHeaders sets the desired headers on the ResponseWriter.
func SetHeaders(cr CustomWR) {
    cr.W.Header().Set("Content-Type", "application/json; charset=UTF-8")
    cr.W.Header().Set("Access-Control-Allow-Origin", "*") // Adjust for production CORS
}


/// Handle Http Error with simple custom function
func Ehandler(cr CustomWR, statusCode int, ErrMessage, LogMessage string) {
    http.Error(cr.W, ErrMessage, statusCode)
    Lg.Error(LogMessage,"")
}