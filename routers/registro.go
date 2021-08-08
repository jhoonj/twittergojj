package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jhoonj/twittergojj/bd"
	"github.com/jhoonj/twittergojj/models"
)

//Registro es la funcion para crear en la base de datos el registro de usuarios
func Registro(rw http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t) // el body es un streams una vez lo leo se destruye
	if err != nil {
		http.Error(rw, "error en los datos recibidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(rw, "error el mail es obligatorio", 400)
		return
	}

	if len(t.PassWord) < 6 {
		http.Error(rw, "el password tiene una longitud minima", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(rw, "ya existe un usuario registrado con ese email", 400)
		return
	}

	_, status, err := bd.InsertarRegistro(t)

	if err != nil {
		http.Error(rw, "ocurrio un error al almacenar el registro"+err.Error(), 400)
	}

	if status == false {
		http.Error(rw, "no se ha logrado insertar el registro del usuario", 400)
	}

	rw.WriteHeader(http.StatusCreated)

}
