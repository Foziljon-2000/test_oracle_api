package handlers

import (
	"net/http"

	"github.com/Foziljon-2000/test_oracle_api/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
)

func JsonGan(c *gin.Context) {
	var txt []models.User

	for i:=1;i<=50;i++{
		name := faker.Name()
		
		txt = append(txt, models.User{Name: name, ID: i})
	}

	c.JSON(http.StatusOK, txt)
}