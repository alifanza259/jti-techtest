package api

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"sync"

	"github.com/alifanza259/jwt-techtest/models"
	"github.com/alifanza259/jwt-techtest/util"
	"github.com/gin-gonic/gin"
)

type createHandphoneRequest struct {
	NoHandphone string `json:"no_handphone" binding:"required" validate:"regexp=^08[1-9][0-9]{6,9}$"`
	Provider    string `json:"provider" binding:"required"`
}

func (server *Server) createHandphone(c *gin.Context) {
	var req createHandphoneRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	decryptedNoHandphone, err := util.Decrypt(req.NoHandphone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	decryptedProvider, err := util.Decrypt(req.Provider)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	handphone := &models.Handphone{
		NoHandphone: decryptedNoHandphone,
		Provider:    decryptedProvider,
	}
	server.db.Create(handphone)

	json, err := json.Marshal(map[string]any{
		"handphone": handphone,
		"type":      "input",
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	go server.broadcast(json)

	c.JSON(http.StatusOK, gin.H{"data": handphone})
}

func (server *Server) generateHandphone(c *gin.Context) {
	availableProviders := []string{"Telkom", "XL", "Smartfren", "Tri"}

	wg := &sync.WaitGroup{}
	for i := 0; i < 25; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			randInt := rand.Intn(4)
			randNoHandphone := 10000000 + rand.Intn(99999999999-10000000)
			handphone := &models.Handphone{
				NoHandphone: "08" + fmt.Sprintf("%d", randNoHandphone),
				Provider:    availableProviders[randInt],
			}
			server.db.Create(handphone)
		}(wg)
	}
	wg.Wait()

	c.JSON(http.StatusOK, gin.H{"data": "success generate 25 data"})
}

func (server *Server) getHandphoneList(c *gin.Context) {
	users := []models.Handphone{}

	server.db.Find(&users)
	oddUsers := []models.Handphone{}
	evenUsers := []models.Handphone{}
	for i := 0; i < len(users); i++ {
		intHandphone, err := strconv.Atoi(users[i].NoHandphone)
		if err != nil {
			fmt.Println(err.Error())
			continue
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

type editHandphoneRequest struct {
	ID          int    `json:"id,string" binding:"required"`
	NoHandphone string `json:"no_handphone" binding:"required" validate:"regexp=^08[1-9][0-9]{6,9}$"`
}

func (server *Server) editHandphone(c *gin.Context) {
	var req editHandphoneRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	handphone := &models.Handphone{}
	server.db.First(&handphone, req.ID)
	if handphone.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "data not found"})
		return
	}

	handphone.NoHandphone = req.NoHandphone
	server.db.Save(&handphone)

	c.JSON(http.StatusOK, gin.H{"data": handphone})
}

type deleteHandphoneRequest struct {
	ID int `uri:"id" binding:"required"`
}

func (server *Server) deleteHandphone(c *gin.Context) {
	var req deleteHandphoneRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	handphone := &models.Handphone{}
	server.db.First(&handphone, req.ID)
	if handphone.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "data not found"})
		return
	}

	server.db.Delete(&handphone)

	c.JSON(http.StatusOK, gin.H{"data": "success delete data"})
}
