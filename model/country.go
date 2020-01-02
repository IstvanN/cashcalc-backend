package model

import (
	"context"
	"log"

	"github.com/IstvanN/cashcalc-backend/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	airCountriesCollectionName  = "countries_air"
	roadCountriesCollectionName = "countries_road"
)

// Country stores the countries with name and a zone number
type Country struct {
	Name       string `json:"name"`
	ZoneNumber int    `json:"zone_number"`
}

// GetAirCountriesFromDB returns with an array of all elements of the airCountries collection
func GetAirCountriesFromDB() (airCountries []Country) {
	coll := database.GetCollectionByName(airCountriesCollectionName)
	cur, err := coll.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		log.Printf("retrieving collection %v failed: %v\n", airCountriesCollectionName, err)
	}

	for cur.Next(context.TODO()) {
		var c Country
		err := cur.Decode(&c)
		if err != nil {
			log.Println("error while retrieving air countries, ", err)
		} else {
			airCountries = append(airCountries, c)
		}
	}
	cur.Close(context.TODO())
	return
}

// GetRoadCountriesFromDB returns with an array of all the elements of the roadCountries collection
func GetRoadCountriesFromDB() (roadCountries []Country) {
	coll := database.GetCollectionByName(roadCountriesCollectionName)
	cur, err := coll.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		log.Printf("retrieving collection %v failed: %v\n", airCountriesCollectionName, err)
	}

	for cur.Next(context.TODO()) {
		var c Country
		err := cur.Decode(&c)
		if err != nil {
			log.Println("error while retrieving air countries, ", err)
		} else {
			roadCountries = append(roadCountries, c)
		}
	}
	cur.Close(context.TODO())
	return
}
