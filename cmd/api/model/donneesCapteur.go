package model

import (
	"github.com/garyburd/redigo/redis"
	"github.com/my/repo/cmd/api/config"
	"log"
	"time"
)

type DonneesCapteur struct {
	Id			 int		`json:"id"`
	IdCapteur    int 		`json:"idCapteur"`
	IdAeroport   string		`json:"idAeroport"`
	TypeMesure   string		`json:"typeMesure"`
	ValeurMesure float32	`json:"valeurMesure"`
	Date         time.Time	`json:"date"`
}

type DonneesCapteurArray []DonneesCapteur

func NewDonneesCapteur(d *DonneesCapteur) {
	if d == nil {
		log.Fatal(d)
	}

	//_, err := config.Db().Do("HMSET", "donnee:" + string(d.Id), "idCapteur", d.IdCapteur, "idAeroport", d.IdAeroport, "typeMesure", d.TypeMesure, "valeurMesure", d.ValeurMesure, "date", d.Date)
	_, err := config.Db().Do("SET",  "idCapteur:" + string(d.Id), string(d.IdCapteur))
	_, err = config.Db().Do("SET",  "idAeroport:" + string(d.Id), d.IdAeroport)
	_, err = config.Db().Do("SET",  "typeMesure:" + string(d.Id), d.IdCapteur)
	_, err = config.Db().Do("SET",  "valeurMesure:" + string(d.Id), d.ValeurMesure)
	_, err = config.Db().Do("SET",  "date:" + string(d.Id), d.Date)
	//err := config.Db().QueryRow("INSERT INTO donneesCapteurs (manufacturer, design, style, doors, created_at, updated_at) VALUES ($1,$2,$3,$4,$5) RETURNING id;", d.IdCapteur, d.IdAeroport, d.TypeMesure, d.ValeurMesure, d.Date).Scan(&d.Id)

	if err != nil {
		log.Fatal(err)
	}
}

func FindDonneesCapteurById(id int) *DonneesCapteur {
	var donnees DonneesCapteur

	//row := config.Db().QueryRow("SELECT * FROM donneesCapteurs WHERE id = $1;", id)
	//err := row.Scan(&donnees.Id, donnees.IdCapteur, donnees.IdAeroport, donnees.TypeMesure, donnees.ValeurMesure, donnees.Date)

	value, err := redis.Values(config.Db().Do("HGETALL", "donnee:" + string(id)))

	if err != nil {
		log.Fatal(err)
	}

	err = redis.ScanStruct(value, &donnees)

	return &donnees
}

func AllDonneesCapteur() *DonneesCapteurArray {
	var donnees DonneesCapteurArray

	values, err := redis.Values(config.Db().Do("HGETALL", "donnee:5577006791947779410"))
	//values, err := redis.String(config.Db().Do("GETALL", "malcom:1"))

	if err != nil {
		log.Fatal(err)
	}

	err = redis.ScanStruct(values, &donnees)

	return &donnees
}

func DeleteDonneesCapteurById(id int) error {
	_, err := config.Db().Do("HDEL",  "idCapteur:" + string(id))
	_, err = config.Db().Do("HDEL",  "idAeroport:" + string(id))
	_, err = config.Db().Do("HDEL",  "typeMesure:" + string(id))
	_, err = config.Db().Do("HDEL",  "valeurMesure:" + string(id))
	_, err = config.Db().Do("HDEL",  "date:" + string(id))

	if err != nil {
		log.Fatal(err)
	}

	return err
}