package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErroAPI representa a respsta de erro da API
type ErroAPI struct {
	Erro string `json:"erro"`
}

// JSON  retorna uma resposta
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "Application/Json")
	w.WriteHeader(statusCode)

	if statusCode != http.StatusNoContent {
		if err := json.NewEncoder(w).Encode(dados); err != nil {
			log.Fatal(err)
		}
	}
}

// TrataStatusCodeDeErro trata as requisições com status code 400 ou superior
func TratarStatusCodeDeErro(w http.ResponseWriter, r *http.Response) {
	var err ErroAPI
	json.NewDecoder(r.Body).Decode(&err)
	JSON(w, r.StatusCode, err)
}
