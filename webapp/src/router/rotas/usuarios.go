package rotas

import "net/http"

var rotasUsuarios = []Rota{
	Rota{
		URI:                "/criar-usuario",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeCadastroDeUsuario,
		RequerAutenticacao: false,
	},
}
