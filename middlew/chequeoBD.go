package middlew

import (
	"net/http"

	"github.com/jhoonj/twittergojj/bd"
)

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConexion() == 0 {
			http.Error(rw, "conexion perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(rw, r)
	}

}
