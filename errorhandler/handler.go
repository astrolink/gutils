package errorhandler

import (
	"encoding/json"
	"github.com/astrolink/gutils/language"
	"net/http"
)

/**
 ErrorHandler foi criado para padronizar e possibilitar a traducao de mensagens de erro
 a ideia eh que seja utilizada para retornos de APIs
 para passar uma mensagem clara tanto ao usuario como ao administrador
	Exemplo de response:
		{
			"error": {
				"message": "unsupported media type",
				"error_user_msg": "Phelipe, to access your account, click the link to reset your password"
			}
		}
**/

type ErrorObject struct {
	Error      error
	StatusCode int
	Lang       Lang
}

type Lang struct {
	Idiom        string
	Replacements []string
	MessageKey   string
	Context      string
	MessagesLang map[string]map[string]string
}

// HasLang valida se foi carregado as lang necessarias
func (l *Lang) HasLang() bool {
	if l.Idiom == "" || len(l.MessagesLang) == 0 || l.Context == "" {
		return false
	}

	return true
}

// LoadLanguages funcao principal necessaria para carregar as mensagens de erro que serao utilizadas
func LoadLanguages(messages map[string]map[string]string, context string) ErrorObject {
	return ErrorObject{
		Lang: Lang{
			MessagesLang: messages,
			Context:      context,
		},
	}
}

// BuildErrorMessage responsavel por atribuir os valores que serao usados para criar uma mensagem de erro internacionalizada
func (e *ErrorObject) BuildErrorMessage(idiom, messageKey string, replacements []string) *ErrorObject {
	e.Lang.Idiom = idiom
	e.Lang.MessageKey = messageKey
	e.Lang.Replacements = replacements
	return e
}

// GetErrorMessage retorna a mensagem de erro como string
func (e *ErrorObject) GetErrorMessage() string {
	if e.Error == nil {
		return ""
	}

	return e.Error.Error()
}

// SetError funcao generica para ser utilizada caso NAO seja necessario mensagem customizada traduzida
func (e *ErrorObject) SetError(err error, statusCode int) *ErrorObject {
	e.Error = err
	e.StatusCode = statusCode
	return e
}

type ErrorResponse struct {
	Error detail `json:"error"`
}

type detail struct {
	Message      string `json:"message"`
	ErrorUserMsg string `json:"error_user_msg"`
}

// getErrorResponse
// responsavel por fazer o parse dos parametros para gerar uma mensagem customizada
// devolver o objeto ErrorResponse que sera utilizada para devolver um erro estruturado na api
func (e *ErrorObject) getErrorResponse() ErrorResponse {
	if !e.Lang.HasLang() {
		return ErrorResponse{
			detail{
				Message: e.GetErrorMessage(),
			},
		}
	}

	language.LoadLang(e.Lang.MessagesLang, e.Lang.Context, e.Lang.Idiom)
	i18nMessageUser := language.Translate(e.Lang.MessageKey, e.Lang.Idiom, e.Lang.Replacements)

	handler := ErrorResponse{
		detail{
			Message:      e.GetErrorMessage(),
			ErrorUserMsg: i18nMessageUser,
		},
	}

	return handler
}

// CreateErrorResponse constroi a estrutura padrao de mensagem de erro e escreve a resposta na request
func (e *ErrorObject) CreateErrorResponse(w http.ResponseWriter, err error, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	e.Error = err
	e.StatusCode = statusCode
	errorResponse := e.getErrorResponse()

	mapB, _ := json.Marshal(errorResponse)
	w.Write(mapB)
}
