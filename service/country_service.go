package service

import (
	"fmt"
	"os"

	"github.com/IstvanN/cashcalc-backend/database"
	"github.com/IstvanN/cashcalc-backend/model"
)

var (
	countriesCollectionName = os.Getenv("COUNTRIES_COLL")
)

// GetCountriesAirFromDB returns with a slice of all air elements of the Countries collection, or an error
func GetCountriesAirFromDB() ([]model.Country, error) {
	c, err := getCountriesFromDB()
	if err != nil {
		return nil, err
	}

	return c.CountriesAir, nil
}

// GetCountriesRoadFromDB returns with an array of all road elements of the Countries collection, or an error
func GetCountriesRoadFromDB() ([]model.Country, error) {
	c, err := getCountriesFromDB()
	if err != nil {
		return nil, err
	}

	return c.CountriesRoad, nil
}

func getCountriesFromDB() (model.Countries, error) {
	coll := database.GetCollectionByName(countriesCollectionName)

	var c model.Countries
	err := coll.Find(nil).One(&c)
	if err != nil {
		return model.Countries{}, fmt.Errorf("error while retrieving collection %v from database: %v", countriesCollectionName, err)
	}

	return c, nil
}
