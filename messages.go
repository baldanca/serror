package serror

var (
	// Default message
	Default Message = "Houve algum problema com a sua solicitação, tente novamente."

	// InvalidID message
	InvalidID Message = "ID inválido!"

	// InvalidLink message
	InvalidLink Message = "Link inválido!"

	// AuthorizationNotFound message
	AuthorizationNotFound Message = "Parametro Authorization não enviado."

	// TokenNotFound message
	TokenNotFound Message = "Parametro token não enviado."

	// EmptyParam message
	EmptyParam Message = "Parâmetro obrigatório não enviado."

	// ServiceUnavailable message
	ServiceUnavailable Message = "Serviço indisponível."
)
