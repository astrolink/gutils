package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	ContentTypeHeaderKey = "Content-Type"
	ApplicationJsonHeaderValue = "application/json"
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

// CreateBadRequestResponse cria uma resposta com HTTP status code 400
// e envia o erro no atributo body
func CreateBadRequestResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
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

// CreateNoContentResponse cria uma resposta com HTTP status code 203
// e não envia dados no atributo body
func CreateNoContentResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

// CreateCreatedResponse cria uma resposta com HTTP status code 201
// e envia dados no atributo body
func CreateCreatedResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set(ContentTypeHeaderKey, ApplicationJsonHeaderValue)

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

	w.WriteHeader(http.StatusCreated)
	w.Write(bytes)
}

// CreateCustomStatusCodeResponse cria uma resposta com o código de status http recebido
// e envia os erros no atributo body
func CreateCustomStatusCodeResponse(statusCode int, w http.ResponseWriter, err error) {
	w.WriteHeader(statusCode)
	responseErrorData(w, err)
}

// CreateTooManyRequestsResponse creates a response with 429 HTTP status code
func CreateTooManyRequestsResponse(w http.ResponseWriter, err error) {
	w.Header().Set(ContentTypeHeaderKey, ApplicationJsonHeaderValue)
	w.WriteHeader(http.StatusTooManyRequests)
	responseErrorData(w, err)
}