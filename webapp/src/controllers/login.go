package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"webapp/src/respostas"
)

// FazerLogin utiliza o e-mail e senha do usuário para autenticar na aplicação
func FazerLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, err := json.Marshal(map[string]string{
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})
	if err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: err.Error()})
		return
	}

	response, err := http.Post("http://localhost:5000/login", "application/json", bytes.NewBuffer(usuario))
	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: err.Error()})
		return
	}
	defer response.Body.Close()

	token, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode, string(token))
}
