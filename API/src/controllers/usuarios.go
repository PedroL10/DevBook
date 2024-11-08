package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io"
	"net/http"
)

// CriarUsuario insere um usuario no banco de dados - Controller(lida com as funções http, requisicoes e repostas)
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
	}

	if erro = usuario.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
	}

	db, erro := banco.Conectar() // Abre a conexao com o banco
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db) // Cria o repositorio e passa o banco pra dentro desse repositorio
	usuario.ID, erro = repositorio.Criar(usuario)             // Cria o usuario que vem na requisicao
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	respostas.JSON(w, http.StatusCreated, usuario)

}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os ususarios"))
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuario"))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando o usuario"))
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuario"))
}
