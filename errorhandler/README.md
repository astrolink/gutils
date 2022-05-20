# ErrorHandler

ErrorHandler foi criado para padronizar e possibilitar a tradução de mensagens de erro.
A ideia principal é que seja utilizado em retornos de APIs feitas em GO, enviando uma  mensagem clara tanto ao usuário como ao administrador do sistema.

### Exemplo de resposta

```
{
	"error": {
		"message": "",
		"error_user_msg": "Facebook didn't send us your account email address. This could be because the email was missing, invalid, or unconfirmed. Please review your Facebook email address or using another methods available to register"
	}
}
```

## Como usar

1 - Crie uma var em global level scope (main.go) carregue o contexto e suas mensagens para tradução utilizando `errorhandler.LoadLanguages()`

### Método 1:
Tenho uma mensagem para ser traduzida para o usuário final

**1 -** Tenha uma instancia de `ErrorMessage` em seus métodos. Você pode optar por:
* Instância-la utilizando NewErrorMessage(<statusCode>)
* Contruí-lo utilizando o BuildErrorMessage

**2 -** Utilize a `var global` instanciada anteriormente para construir o objeto `ErrorObject`
**3 -** Crie a resposta chamando o método `CreateErrorResponseV2`

### Método 2:
Só quero transmitir um erro! Sem mensagem traduzida para o usuário

**1 -** Tenha uma instancia de `ErrorMessage` em seus métodos. Você pode optar por:
* Instância-la utilizando `NewErrorMessage(<statusCode>)`
* Contruí-lo utilizando o `BuildErrorMessage`

**2 -** Utilize a `var global` para construir o objeto `ErrorObject`
**3 -** Crie a resposta chamando o método `CreateErrorResponseV2`