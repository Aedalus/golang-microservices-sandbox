package controllers

import (
	"github.com/gin-gonic/gin"
	"mvc/utils"

	"mvc/services"
	"net/http"
	"strconv"
)

func GetUser( c *gin.Context){
	userId, err := strconv.ParseInt(c.Param("user_id"),10,64)
	if err != nil {
		userErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		c.JSON(userErr.StatusCode, userErr)
		return
	}

	user, appErr := services.UsersService.GetUser(userId)
	if appErr != nil {
		// Handle the error
		c.JSON(appErr.StatusCode,appErr)
		return
	}

	c.JSON(http.StatusOK, user)
}
