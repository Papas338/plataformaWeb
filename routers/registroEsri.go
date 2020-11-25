package routers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/proyLSIPAZUD/plataformaWeb/bd"
	"github.com/proyLSIPAZUD/plataformaWeb/models"
)

/*RegistroEsri es la función para crear en la bd el registro de lider */
func RegistroEsri(t models.LiderGeneral) {

	var s []*models.Municipio

	s, _ = bd.BuscarMunicipio(t.Departamento, t.Municipio)

	url := "https://services.arcgis.com/deQSb0Gn7gDPf3uV/arcgis/rest/services/lideres_sociales_indigenas/FeatureServer/0/applyEdits"
	method := "POST"

	payload := strings.NewReader("f=json&token=CQuxYHA5U6u2iOVKr44eNC8jcyxuxAMlpkyUavlvcpmh2enydJ5BRUZozPq2MrEk76JFPcqKGgwgd7VnN3WjS9td8c-Mrp8xQ9biFDGaBt-GzGTXn3WmQuS_iCijRZIk66tpm3vl4RPiQ4NkK0K6BUS2r4Oe0UvfNhgN4CdkyBrvtKd--TjOsX-R0CMUmctcMb6EWHUB3FBrFJUp05U44-bDqtVIs95ncUuEScj4jhE.&adds=%5B%0A%20%20%7B%0A%20%20%20%20%22geometry%22%20%3A%20%7B%0A%20%20%20%20%20%20%22x%22%3A%20" + FloatToString(s[0].X) + "%2C%0A%20%20%20%20%20%20%22y%22%3A%20" + FloatToString(s[0].Y) + "%2C%0A%20%20%20%20%20%20%22spatialReference%22%3A%20%7B%0A%20%20%20%20%20%20%20%20%22wkid%22%3A%204326%0A%20%20%20%20%20%20%7D%0A%20%20%20%20%7D%2C%0A%20%20%20%20%22attributes%22%20%3A%20%7B%0A%20%20%20%20%20%20%22Nombre%22%20%3A%20%22" + t.Nombre + "%22%2C%0A%20%20%20%20%20%20%22TipoLiderazgo%22%20%3A%20%22" + t.TipoLiderazgo + "%22%2C%0A%20%20%20%20%20%20%22Territorio%22%20%3A%20%22" + t.Territorio + "%22%2C%0A%20%20%20%20%20%20%22Palabra1%22%20%3A%20%22" + t.Palabra1 + "%22%2C%0A%20%20%20%20%20%20%22Palabra2%22%20%3A%20%22" + t.Palabra2 + "%22%2C%0A%20%20%20%20%20%20%22Palabra3%22%20%3A%20%22" + t.Palabra3 + "%22%2C%0A%20%20%20%20%20%20%22Palabra4%22%20%3A%20%22" + t.Palabra4 + "%22%2C%0A%20%20%20%20%20%20%22Palabra5%22%20%3A%20%22" + t.Palabra5 + "%22%2C%0A%20%20%20%20%20%20%22Palabra6%22%20%3A%20%22" + t.Palabra6 + "%22%2C%0A%20%20%20%20%20%20%22Palabra7%22%20%3A%20%22" + t.Palabra7 + "%22%2C%0A%20%20%20%20%20%20%22Palabra8%22%20%3A%20%22" + t.Palabra8 + "%22%2C%0A%20%20%20%20%20%20%22Palabra9%22%20%3A%20%22" + t.Palabra9 + "%22%2C%0A%20%20%20%20%20%20%22Palabra10%22%20%3A%20%22" + t.Palabra10 + "%22%2C%0A%20%20%20%20%20%20%22Departamento%22%20%3A%20%22" + t.Departamento + "%22%2C%0A%20%20%20%20%20%20%22Municipio%22%20%3A%20%22" + t.Municipio + "%22%0A%20%20%20%20%7D%0A%20%20%7D%0A%5D")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

/*FloatToString convierte los float64 a string*/
func FloatToString(inputNum float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(inputNum, 'f', 6, 64)
}
