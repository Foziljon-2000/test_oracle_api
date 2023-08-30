package handlers

import (
	"net/http"

	"github.com/Foziljon-2000/test_oracle_api/internal/models"
	"github.com/gin-gonic/gin"
)

func JsonGan(c *gin.Context) {

	txt := &models.User{Name:"Youssef", ID: 1}

	c.JSON(http.StatusOK, txt)
}