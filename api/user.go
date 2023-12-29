package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"

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

	json, err := json.Marshal(map[string]any{
		"user": user,
		"type": "input",
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	go server.broadcast(json)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (server *Server) autoInputData(c *gin.Context) {
	wg := &sync.WaitGroup{}
	for i := 0; i < 25; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			user := &models.User{
				NoHandphone: "08" + fmt.Sprintf("%d", 123),
				Provider:    "telkomsel",
			}
			server.db.Create(user)
		}(wg)
	}
	wg.Wait()

	c.JSON(http.StatusOK, gin.H{"data": "success generate 25 data"})
}

func (server *Server) getData(c *gin.Context) {
	users := []models.User{}

	server.db.Find(&users)
	fmt.Println("lewat")
	oddUsers := []models.User{}
	evenUsers := []models.User{}
	for i := 0; i < len(users); i++ {
		intHandphone, err := strconv.Atoi(users[i].NoHandphone)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if intHandphone%2 == 1 {
			oddUsers = append(oddUsers, users[i])
		} else {
			evenUsers = append(evenUsers, users[i])
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": map[string]any{
		"oddUsers":  oddUsers,
		"evenUsers": evenUsers,
	}})
}

type EditDataRequest struct {
	ID          int    `json:"id,string" binding:"required"`
	NoHandphone string `json:"no_handphone" binding:"required"`
}

func (server *Server) editData(c *gin.Context) {
	var req EditDataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{}
	server.db.First(&user, req.ID)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "data not found"})
		return
	}

	user.NoHandphone = req.NoHandphone
	server.db.Save(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

type deleteDataRequest struct {
	ID int `uri:"id" binding:"required"`
}

func (server *Server) deleteData(c *gin.Context) {
	var req deleteDataRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{}
	server.db.First(&user, req.ID)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "data not found"})
		return
	}

	server.db.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": "success delete data"})
}
