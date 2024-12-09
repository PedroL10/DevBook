package modelos

// Senha representa o formatdo da requisicao de alteração de senha
type Senha struct {
	Nova  string `json:"nova"`
	Atual string `json:"atual"`
}
