package controllers

import (
	"encoding/json"
	"mvc/utils"

	"mvc/services"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userIdParam := req.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam,10,64)
	if err != nil {
		userErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(userErr)
		// Return bad request
		resp.WriteHeader(userErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	user, appErr := services.GetUser(userId)
	if appErr != nil {
		// Handle the error
		jsonValue, err := json.Marshal(appErr)
		if err != nil {
			panic(err)
		}
		resp.WriteHeader(appErr.StatusCode)
		resp.Write([]byte(jsonValue))
		return
	}

	// return user to client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
