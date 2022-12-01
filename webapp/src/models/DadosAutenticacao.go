package models

// DadosAutenticacao contém o id e o token de usuário autenticado
type DadosAutenticacao struct {
	ID    string `json:"id"`
	Token string `json:"Token"`
}
