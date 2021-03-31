package apphandlers

import (
	"encoding/json"
	"net/http"
	"onecause/models"
	"onecause/utils"
	"time"

	"github.com/go-playground/validator/v10"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginReqData models.LoginReqModel
	decodeErr := json.NewDecoder(r.Body).Decode(&loginReqData)
	if decodeErr != nil {
		utils.RespondWithError(w, 401, decodeErr.Error())
		return
	}
	validate := validator.New()
	validationErr := validate.Struct(loginReqData)
	if validationErr != nil {
		validationErrors := validationErr.(validator.ValidationErrors)
		utils.RespondWithError(w, 401, validationErrors.Error())
		return
	}

	ottValidationErr := loginReqData.IsOTTValid()
	if ottValidationErr != nil {
		utils.RespondWithError(w, 401, ottValidationErr.Error())
		return
	}
	currentTime := time.Now()
	if loginReqData.Hour != uint8(currentTime.Hour()) ||
		loginReqData.Minute != uint8(currentTime.Minute()) ||
		loginReqData.Username != "c137@onecause.com" ||
		loginReqData.Password != "#th@nH@rm#y#r!$100%D0p#" {
		utils.RespondWithError(w, 401, "Your credentials does not match")
		return
	}

	utils.RespondWithJSON(w, 200, "Login successfull")
}
