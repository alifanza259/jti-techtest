package api

import (
	"github.com/alifanza259/jwt-techtest/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db = make(map[string]string)

type Server struct {
	config util.Config
	router *gin.Engine
	db     *gorm.DB
}

func NewServer(config util.Config, db *gorm.DB) *Server {
	server := &Server{
		config: config,
		db:     db,
	}
	server.setupRouter()

	return server
}

func (server *Server) setupRouter() {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.POST("/user", server.inputData)
	r.POST("/user/auto", server.autoInputData)
	// Ping test
	// r.GET("/ping", func(c *gin.Context) {
	// 	var user models.User
	// 	server.db.First(&user)
	// 	c.JSON(http.StatusOK, user)
	// })

	// // Get user value
	// r.GET("/user/:name", func(c *gin.Context) {
	// 	user := c.Params.ByName("name")
	// 	value, ok := db[user]
	// 	if ok {
	// 		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	// 	} else {
	// 		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	// 	}
	// })

	// // Authorized group (uses gin.BasicAuth() middleware)
	// // Same than:
	// // authorized := r.Group("/")
	// // authorized.Use(gin.BasicAuth(gin.Credentials{
	// // 	"foo":  "bar",
	// // 	"manu": "123",
	// // }))
	// authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
	// 	"foo":  "bar", // user:foo password:bar
	// 	"manu": "123", // user:manu password:123
	// }))

	// /* example curl for /admin with basicauth header
	//    Zm9vOmJhcg== is base64("foo:bar")

	// 	curl -X POST \
	//   	http://localhost:8080/admin \
	//   	-H 'authorization: Basic Zm9vOmJhcg==' \
	//   	-H 'content-type: application/json' \
	//   	-d '{"value":"bar"}'
	// */
	// authorized.POST("admin", func(c *gin.Context) {
	// 	user := c.MustGet(gin.AuthUserKey).(string)

	// 	// Parse JSON
	// 	var json struct {
	// 		Value string `json:"value" binding:"required"`
	// 	}

	// 	if c.Bind(&json) == nil {
	// 		db[user] = json.Value
	// 		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	// 	}
	// })

	server.router = r
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
