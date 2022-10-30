package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/murilorscampos/go-rest-api/database"
	"github.com/murilorscampos/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Home Page")

}

func TodasAsPersonalidades(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")

	var p []models.Personalidade

	database.DB.Find(&p)

	json.NewEncoder(w).Encode(p)

}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade

	database.DB.First(&personalidade, id)

	json.NewEncoder(w).Encode(personalidade)

}

func CriaUmaNovaPersonalida(w http.ResponseWriter, r *http.Request) {

	var novaPersonalidade models.Personalidade

	json.NewDecoder(r.Body).Decode(&novaPersonalidade)

	database.DB.Create(&novaPersonalidade)

	json.NewEncoder(w).Encode(novaPersonalidade)

}

func DeletaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id := vars["id"]

	var personalidade models.Personalidade

	database.DB.Delete(&personalidade, id)

	json.NewEncoder(w).Encode(personalidade)

}

func EditaPersonsalidade(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id := vars["id"]

	var personalidade models.Personalidade

	database.DB.First(&personalidade, id)

	json.NewDecoder(r.Body).Decode(&personalidade)

	database.DB.Save(&personalidade)

	json.NewEncoder(w).Encode(&personalidade)

}
