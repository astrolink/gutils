package errorhandler

import (
	"encoding/json"
	ghttp "github.com/astrolink/gutils/http"
	"github.com/astrolink/gutils/language"
	"net/http"
)

// ErrorObject abstração responsável por armazenar as informações de erro
type ErrorObject struct {
	ErrorMessage *ErrorMessage
	Lang         Lang
}

// Lang estrutura responsável por armazenar o contexto e o array de mensagens para tradução
type Lang struct {
	Idiom        string
	Context      string
	MessagesLang map[string]map[string]string
}

// BuildMessage constroi o objeto ErrorObject utilizado pela aplicacao para montar a estrutura padrao de resposta
func (e *ErrorObject) BuildMessage(idiom string, errorMessage *ErrorMessage) *ErrorObject {
	e.ErrorMessage = errorMessage
	e.Lang.Idiom = idiom
	return e
}

// IsEmpty verifica se foram carregadas as lang necessárias
func (l *Lang) IsEmpty() bool {
	return l.Idiom == "" && len(l.MessagesLang) == 0 && l.Context == ""
}

// LoadLanguages função principal necessária para carregar as mensagens de erro que serão utilizadas
func LoadLanguages(messages map[string]map[string]string, context string) ErrorObject {
	return ErrorObject{
		Lang: Lang{
			MessagesLang: messages,
			Context:      context,
		},
	}
}

// BuildErrorMessage responsável por atribuir os valores que serão usados para criar uma mensagem de erro internacionalizada
func (e *ErrorObject) BuildErrorMessage(idiom, messageKey string, replacements []string) *ErrorObject {
	e.Lang.Idiom = idiom
	e.ErrorMessage = &ErrorMessage{
		MessageKey:   messageKey,
		Replacements: replacements,
	}
	return e
}

// GetError retorna a mensagem de erro no nível de desenvolvimento, ou seja a info de forma técnica
func (e *ErrorObject) GetError() string {
	return e.GetErrorMessage().getError()
}

// GetStatusCode retorna o statuscode caso nao exista retorna status 500
func (e *ErrorObject) GetStatusCode() int {
	return e.GetErrorMessage().StatusCode
}

// GetErrorMessage retorna uma instancia de ErrorMessage caso nao exista
func (e *ErrorObject) GetErrorMessage() *ErrorMessage {
	if e.ErrorMessage == nil {
		e.ErrorMessage = NewErrorMessage(http.StatusInternalServerError)
	}

	return e.ErrorMessage
}

// SetError função genérica para ser utilizada caso NÃO seja necessária mensagem customizada traduzida
func (e *ErrorObject) SetError(err error, statusCode int) *ErrorObject {
	errorMessage := e.GetErrorMessage()
	errorMessage.Err = err
	errorMessage.StatusCode = statusCode
	return e
}

// getErrorResponse responsável por fazer o parse dos parâmetros para gerar uma mensagem customizada
// devolver o objeto ErrorResponse que será utilizado para devolver um erro estruturado na api
func (e *ErrorObject) getErrorResponse() ErrorResponse {
	errorResponse := ErrorResponse{
		Detail{
			Message: e.GetError(),
		},
	}

	if e.Lang.IsEmpty() || e.ErrorMessage.IsEmpty() {
		return errorResponse
	}

	language.LoadLang(e.Lang.MessagesLang, e.Lang.Context, e.Lang.Idiom)
	i18nMessageUser := language.Translate(e.ErrorMessage.MessageKey, e.Lang.Idiom, e.ErrorMessage.Replacements)

	errorResponse.BuildDetail(e.GetError(), i18nMessageUser)
	return errorResponse
}

// CreateErrorResponseV2 constrói a estrutura padrão de mensagem de erro utilizando os valores setados anteriormente
// e escreve a resposta no response
// Necessário o objeto ErrorMessage exista e tenha sido construido utilizando o metodo ErrorObject.BuildMessage
func (e *ErrorObject) CreateErrorResponseV2(w http.ResponseWriter) {
	w.Header().Set(ghttp.ContentTypeHeaderKey, ghttp.ApplicationJsonHeaderValue)
	w.WriteHeader(e.GetStatusCode())
	errorResponse := e.getErrorResponse()

	mapB, _ := json.Marshal(errorResponse)
	w.Write(mapB)
}

// CreateErrorResponse constrói a estrutura padrão de mensagem de erro e escreve a resposta no response
func (e *ErrorObject) CreateErrorResponse(w http.ResponseWriter, err error, statusCode int) {
	w.Header().Set(ghttp.ContentTypeHeaderKey, ghttp.ApplicationJsonHeaderValue)
	w.WriteHeader(statusCode)

	e.ErrorMessage.Err = err
	e.ErrorMessage.StatusCode = statusCode
	errorResponse := e.getErrorResponse()

	mapB, _ := json.Marshal(errorResponse)
	w.Write(mapB)
}
