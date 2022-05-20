package errorhandler

// ErrorMessage estrutur responsavel por armazenar informacoes de erro que serao enviadas ao usuario
type ErrorMessage struct {
	Err          error
	StatusCode   int
	MessageKey   string
	Replacements []string
}

// getError retorna ErrorMessage.Err.Error() caso exista
func (e *ErrorMessage) getError() string {
	if e.Err == nil {
		return ""
	}

	return e.Err.Error()
}

// IsEmpty valida se ErrorMessage é vazio
func (e *ErrorMessage) IsEmpty() bool {
	if e == nil {
		return true
	}

	return false
}

// NewErrorMessage instancia um novo objeto ErrorMessage setando o statusCode utilizado na resposta do servidor
func NewErrorMessage(defaultStatusCode int) *ErrorMessage {
	return &ErrorMessage{
		StatusCode: defaultStatusCode,
	}
}

// BuildErrorMessage constroi um objeto ErrorMessage de forma completa
func BuildErrorMessage(err error, statusCode int, messageKey string, replacements []string) *ErrorMessage {
	return &ErrorMessage{
		Err:          err,
		StatusCode:   statusCode,
		MessageKey:   messageKey,
		Replacements: replacements,
	}
}
