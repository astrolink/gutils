package errorhandler

/**
Exemplo de response:
	{
		"error": {
			"message": "unsupported media type",
			"error_user_msg": "Phelipe, to access your account, click the link to reset your password"
		}
	}
**/

// ErrorResponse estrutura padrao de respoosta de erro em APIS
type ErrorResponse struct {
	Error Detail `json:"error"`
}

// Detail armazena os valores detalhados para resposta ao usuário
type Detail struct {
	Message      string `json:"message"`
	ErrorUserMsg string `json:"error_user_msg"`
}

// BuildDetail constroi o objeto Detail
func (e *ErrorResponse) BuildDetail(message, messageForUser string) {
	e.Error.Message = message
	e.Error.ErrorUserMsg = messageForUser
}
