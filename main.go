/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/IstvanN/cashcalc-backend/controllers"
	"github.com/IstvanN/cashcalc-backend/database"
	"github.com/IstvanN/cashcalc-backend/properties"
)

var port = ":" + os.Getenv("PORT")

func main() {
	properties.InitProperties()
	mongo := database.StartupMongo()
	database.StartupPostgres()
	defer mongo.Close()

	router := controllers.StartupRouter()

	log.Println("cashcalc-backend is up and running on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
