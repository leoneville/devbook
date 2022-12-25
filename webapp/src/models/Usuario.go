package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requisicoes"
)

// Usuario representa uma pessoa utilizando a rede social
type Usuario struct {
	ID          uint64       `json:"id"`
	Nome        string       `json:"nome"`
	Email       string       `json:"email"`
	Nick        string       `json:"nick"`
	CriadoEm    time.Time    `json:"criadoEm"`
	Seguidores  []Usuario    `json:"seguidores"`
	Seguindo    []Usuario    `json:"seguindo"`
	Publicacoes []Publicacao `json:"publicacoes"`
}

// BuscarUsuarioCompleto faz 4 requisições na API para montar o usuário
func BuscarUsuarioCompleto(usuarioID uint64, r *http.Request) (Usuario, error) {
	var err error

	canalUsuario := make(chan Usuario)
	canalSeguidores := make(chan []Usuario)
	canalSeguindo := make(chan []Usuario)
	canalPublicacoes := make(chan []Publicacao)

	go func() {
		err = BuscarDadosDoUsuario(canalUsuario, usuarioID, r)
	}()
	go func() {
		err = BuscarSeguidores(canalSeguidores, usuarioID, r)
	}()
	go func() {
		err = BuscarSeguindo(canalSeguindo, usuarioID, r)
	}()
	go func() {
		err = BuscarPublicacoes(canalPublicacoes, usuarioID, r)
	}()

	var (
		usuario     Usuario
		seguidores  []Usuario
		seguindo    []Usuario
		publicacoes []Publicacao
	)

	for i := 0; i < 4; i++ {
		select {
		case usuarioCarregado := <-canalUsuario:
			if usuarioCarregado.ID == 0 && err != nil {
				return Usuario{}, errors.New("erro ao buscar o usuário")
			}

			usuario = usuarioCarregado

		case seguidoresCarregados := <-canalSeguidores:
			if seguidoresCarregados == nil && err != nil {
				return Usuario{}, errors.New("erro ao buscar os seguidores")
			}

			seguidores = seguidoresCarregados

		case seguindoCarregados := <-canalSeguindo:
			if seguindoCarregados == nil && err != nil {
				return Usuario{}, errors.New("erro ao buscar quem o usuário está seguindo")
			}

			seguindo = seguindoCarregados

		case publicacoesCarregadas := <-canalPublicacoes:
			if publicacoesCarregadas == nil && err != nil {
				return Usuario{}, errors.New("erro ao buscar publicações")
			}

			publicacoes = publicacoesCarregadas
		}
	}

	usuario.Seguidores = seguidores
	usuario.Seguindo = seguindo
	usuario.Publicacoes = publicacoes

	return usuario, nil
}

// BuscarDadosDoUsuario chama a API para buscar os dados base do usuário
func BuscarDadosDoUsuario(canal chan<- Usuario, usuarioID uint64, r *http.Request) error {
	url := fmt.Sprintf("%s/usuarios/%d", config.APIURL, usuarioID)

	response, err := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if err != nil {
		canal <- Usuario{}
		return err
	}
	defer response.Body.Close()

	var usuario Usuario
	if err = json.NewDecoder(response.Body).Decode(&usuario); err != nil {
		canal <- Usuario{}
		return err
	}

	canal <- usuario

	return nil
}

// BuscarSeguidores chama a API para buscar os seguidores do usuário
func BuscarSeguidores(canal chan<- []Usuario, usuarioID uint64, r *http.Request) error {
	url := fmt.Sprintf("%s/usuarios/%d/seguidores", config.APIURL, usuarioID)

	response, err := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if err != nil {
		canal <- nil
		return err
	}
	defer response.Body.Close()

	var seguidores []Usuario

	if err = json.NewDecoder(response.Body).Decode(&seguidores); err != nil {
		canal <- nil
		return err
	}

	canal <- seguidores

	return nil
}

// BuscarSeguindo chama a API para buscar quem o usuário está seguindo
func BuscarSeguindo(canal chan<- []Usuario, usuarioID uint64, r *http.Request) error {
	url := fmt.Sprintf("%s/usuarios/%d/seguindo", config.APIURL, usuarioID)

	response, err := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if err != nil {
		canal <- nil
		return err
	}
	defer response.Body.Close()

	var seguindo []Usuario

	if err = json.NewDecoder(response.Body).Decode(&seguindo); err != nil {
		canal <- nil
		return err
	}

	canal <- seguindo

	return nil
}

// BuscarPublicacoes chama a API para buscar as publicações de um usuário
func BuscarPublicacoes(canal chan<- []Publicacao, usuarioID uint64, r *http.Request) error {
	url := fmt.Sprintf("%s/usuarios/%d/publicacoes", config.APIURL, usuarioID)

	response, err := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if err != nil {
		canal <- nil
		return err
	}
	defer response.Body.Close()

	var publicacoes []Publicacao

	if err = json.NewDecoder(response.Body).Decode(&publicacoes); err != nil {
		canal <- nil
		return err
	}

	canal <- publicacoes

	return nil
}
