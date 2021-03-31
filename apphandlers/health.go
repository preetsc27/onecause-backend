package apphandlers

import (
	"net/http"
	"onecause/utils"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {

	health := 100
	// Some login to calculate health.
	// 1.Db connection check
	// 2.Other microservice reachable chek
	// 3.Cache connection check
	res := map[string]int{"health": health}

	utils.RespondWithJSON(w, 200, res)
}
