package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/jcesarcr/twitter/bd"
)

/*ObtenerAvatar envia el avatar al http*/
func ObtenerAvatar(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/avatars" + perfil.Avatar)
	if err != nil {
		http.Error(w, "imagen no encontrada", http.StatusBadRequest)
		return
	}
	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error al copiar la imagen", http.StatusBadRequest)
		return
	}

}
