package controllers

import (
	"net/http"
	"strings"

	"github.com/JustSomeHack/one-oauth2-server/models"
	"github.com/JustSomeHack/one-oauth2-server/services/authenticationservice"

	"github.com/gin-gonic/gin"
)

var authenticationService authenticationservice.AuthenticationService

// AuthenticationPOST user login POST route
func AuthenticationPOST(c *gin.Context) {
	var user models.User
	if c.ShouldBind(&user) != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "bad request body",
		})
		return
	}
	if user.Username == "" || user.Password == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "bad request body",
		})
		return
	}

	authenticationService = authenticationservice.NewAuthenticationService()
	userMap, err := authenticationService.Authorize(user.Username, user.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "user not authorized",
		})
		return
	}

	if userMap == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "user not authorized",
		})
		return
	}
	c.JSON(http.StatusOK, userMap)
}

// ValidatePOST checks authorization token
func ValidatePOST(c *gin.Context) {
	bearerToken := make([]string, 0)
	authorizationHeader := c.Request.Header.Get("authorization")
	if authorizationHeader == "" {
		if c.Query("token") != "" {
			bearerToken = []string{"Bearer", c.Query("token")}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "An authorization header is required",
			})
			return
		}
	} else {
		bearerToken = strings.Split(authorizationHeader, " ")
	}

	if len(bearerToken) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid authorization token",
		})
	}

	authenticationService = authenticationservice.NewAuthenticationService()
	claims, err := authenticationService.Validate(bearerToken[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "User not authorized",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"decoded": claims,
	})
}
