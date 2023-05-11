package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jcesarcr/twitter/bd"
	"github.com/jcesarcr/twitter/models"
)

/*Registro es la funcion para crear el registro de usuario en la bd*/
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "el email de usuario es requerido ", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "el password tiene que ser mayor a 6 caracteres ", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "ya existe un usuario registrado con ese email", 400)
		return
	}
	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "no se ha logrado insertar el registro de usuario ", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
