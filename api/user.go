package api

import (
	"fmt"
	"net/http"

	"github.com/alifanza259/jwt-techtest/models"
	"github.com/gin-gonic/gin"
)

type InputDataRequest struct {
	NoHandphone string `json:"no_handphone" binding:"required"`
	Provider    string `json:"provider" binding:"required"`
}

func (server *Server) inputData(c *gin.Context) {
	var req InputDataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{
		NoHandphone: req.NoHandphone,
		Provider:    req.Provider,
	}
	server.db.Create(user)

	c.JSON(200, gin.H{"data": user})
}

func (server *Server) autoInputData(c *gin.Context) {
	// var wg sync.WaitGroup
	for i := 0; i < 25; i++ {
		go func() {
			user := &models.User{
				NoHandphone: "08" + fmt.Sprintf("%d", i),
				Provider:    "telkomsel",
			}
			server.db.Create(user)
		}()
	}

	c.JSON(200, gin.H{"data": "success"})
}
