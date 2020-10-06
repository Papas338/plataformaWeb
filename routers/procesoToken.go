package routers

import (
	"errors"
	"strings"

	"github.com/proyLSIPAZUD/plataformaWeb/bd"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/proyLSIPAZUD/plataformaWeb/models"
)

/*Email valor de Email usado en todos los EndPoints */
var Email string

/*IDUsuario es el ID devuelto del modelo, que se usará en todos los EndPoints */
var IDUsuario string

/*RoleUsuario es el role del usuario usado en todos los EndPoints */
var RoleUsuario string

/*ProcesoToken proceso token para extraer sus valores */
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("proyectols2020")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
			RoleUsuario = claims.Role
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token Inválido")
	}

	return claims, false, string(""), err
}
