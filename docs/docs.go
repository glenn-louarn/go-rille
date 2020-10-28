// Package classification Go-rille.
//
// Documentation de notre API
//
//     Schemes: http
//     BasePath: /
//     Version: 1.0.0
//     Host: localhost:8081
//
//     Consumes:
//     - application/json
//	   - application/xml
//	   - application/urlform
//	   - application/bin
//	   - application/multipartform
//	   - application/yaml
//	   - application/txt
//
//     Produces:
//     - application/json
//	   - application/xml
//	   - application/urlform
//	   - application/bin
//	   - application/multipartform
//	   - application/yaml
//	   - application/txt
//     Security:
//     - basic
//
//    SecurityDefinitions:
//    basic:
//      type: basic
//
// swagger:meta
package docs

import "github.com/my/repo/cmd/api/model"


type DonneesBetweenDatesRequest struct {
	//in:path
	Airport string `json:"idAirport"`
	//in:path
	Date1 string `json:"date1"`
	//in:path
	Date2 string `json:"date2"`
	//in:path
	SensorType string `json:"sensorType"`
}

type AverageValuesRequest struct {
	//in:path
	Airport string `json:"idAirport"`
	//in:path
	Date string `json:"date"`
}


// swagger:route GET /donnees/{idAirport}/{date} getAverageSensorByAirport idGet0
// Renvoie la moyenne des valeurs des trois senseurs pour un aéroport et une journée donnée.
// responses:
//   200: res

// Voici la reponse reçu.
// swagger:response res
type resWrapper struct {
	// in:body
	Body model.AverageValues
}



// swagger:route GET /donnees/{idAirport}/{date1}/{date2}/{sensorType} getSensorType idGet1
// Renvoie l'ensemble des valeurs du senseur pour l'aéroport passé en paramètre entre la date1 et la date2.
// responses:
//   200: foobarResponse

// Voici la reponse reçu.
// swagger:response foobarResponse
type foobarResponseWrapper struct {
	// in:body
	Body []model.DonneesBetweenDates
}
// swagger:parameters idGet1
type foobarParamsWrapper struct {
	// in:path
	DonneesBetweenDatesRequest
}
// swagger:parameters idGet0
type foobarParamsWrappers struct {
	// in:path
	AverageValuesRequest
}
