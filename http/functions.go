package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	ContentTypeHeaderKey                = "Content-Type"
	ApplicationJsonHeaderValue          = "application/json"
	CustomRealIpHeaderKey               = "X-Real-Ip"
	CustomForwardedForKey               = "X-Forwarded-For"
	ErrorParsingDataToBytesErrorMessage = "error parsing data interface to bytes array, %s"
)

func responseErrorData(w http.ResponseWriter, err error) {
	if err != nil {
		message := err.Error()

		mapB, err := json.Marshal(map[string]string{"error": message})
		if err != nil {
			log.Println(err)
		}

		w.Write(mapB)
	}
}

func responseData(w http.ResponseWriter, data interface{}) {
	var bytes []byte
	var err error

	if data != nil {
		bytes, err = json.Marshal(data)

		if err != nil {
			err = fmt.Errorf(ErrorParsingDataToBytesErrorMessage, err.Error())
			log.Println(err)
			return
		}
	}

	w.Write(bytes)
}

// CreateBadRequestResponse cria uma resposta com HTTP status code 400
// e envia o erro no atributo body
func CreateBadRequestResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	responseErrorData(w, err)
}

// CreateNotFoundResponse cria uma resposta com HTTP status code 404
// e envia o erro no atributo body
func CreateNotFoundResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusNotFound)
	responseErrorData(w, err)
}

// CreateInternalServerErrorResponse cria uma resposta com HTTP status code 500
// e envia o erro no atributo body
func CreateInternalServerErrorResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	responseErrorData(w, err)
}

// CreateSuccessResponse cria uma resposta com HTTP status code 500
// e envia os dados no atributo body
func CreateSuccessResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set(ContentTypeHeaderKey, ApplicationJsonHeaderValue)
	responseData(w, data)
}

// CreateNoContentResponse cria uma resposta com HTTP status code 203
// e não envia dados no atributo body
func CreateNoContentResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

// CreateCreatedResponse cria uma resposta com HTTP status code 201
// e envia dados no atributo body
func CreateCreatedResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set(ContentTypeHeaderKey, ApplicationJsonHeaderValue)
	w.WriteHeader(http.StatusCreated)
	responseData(w, data)
}

// CreateCustomStatusCodeResponse cria uma resposta com o código de status http recebido
// e envia os dados no atributo body
func CreateCustomStatusCodeResponse(statusCode int, w http.ResponseWriter, data interface{}) {
	w.WriteHeader(statusCode)
	responseData(w, data)
}

// CreateTooManyRequestsResponse creates a response with 429 HTTP status code
func CreateTooManyRequestsResponse(w http.ResponseWriter, err error) {
	w.Header().Set(ContentTypeHeaderKey, ApplicationJsonHeaderValue)
	w.WriteHeader(http.StatusTooManyRequests)
	responseErrorData(w, err)
}
