package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/proyLSIPAZUD/plataformaWeb/models"
)

/*GeneroJWT genera el encriptado con JWT */
func GeneroJWT(t models.Usuario) (string, error) {

	miClave := []byte("proyectols2020")

	payload := jwt.MapClaims{
		"email":      t.Email,
		"role":       t.Role,
		"isvalidate": t.IsValidate,
		"_id":        t.ID.Hex(),
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
