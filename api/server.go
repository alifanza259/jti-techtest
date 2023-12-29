package api

import (
	"fmt"
	"net/http"

	"github.com/alifanza259/jwt-techtest/util"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

type Server struct {
	config util.Config
	router *gin.Engine
	db     *gorm.DB
	conns  map[*websocket.Conn]bool
}

func NewServer(config util.Config, db *gorm.DB) *Server {
	server := &Server{
		config: config,
		db:     db,
		conns:  make(map[*websocket.Conn]bool),
	}
	server.setupRouter()

	return server
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (server *Server) broadcast(b []byte) {
	for ws := range server.conns {
		go func(ws *websocket.Conn) {
			if err := ws.WriteMessage(websocket.TextMessage, b); err != nil {
				fmt.Println(err)
				delete(server.conns, ws)
			}
		}(ws)
	}
}

func (server *Server) setupRouter() {
	r := gin.Default()

	r.Use(CORSMiddleware())

	r.GET("/ws", func(c *gin.Context) {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		server.conns[conn] = true

	})
	r.GET("/homepage", func(c *gin.Context) {
		c.File("homepage.html")
	})
	r.GET("/input", func(c *gin.Context) {
		c.File("input.html")
	})
	r.GET("/output", func(c *gin.Context) {
		c.File("output.html")
	})

	r.GET("/user", server.getData)
	r.POST("/user", server.inputData)
	r.POST("/user/auto", server.autoInputData)
	r.PATCH("/user", server.editData)
	r.DELETE("/user/:id", server.deleteData)

	server.router = r
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
