package controllers

import (
	"fmt"
	"net/http"

	"github.com/IstvanN/cashcalc-backend/security"

	"github.com/gorilla/mux"
)

func registerLoginRoutes(router *mux.Router) {
	router.HandleFunc("/login", loginHandler).Methods(http.MethodGet)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	st, err := security.CreateToken("carrier")
	if err != nil {
		logErrorAndSendHTTPError(w, err, 500)
		return
	}

	fmt.Fprintln(w, st)
}
