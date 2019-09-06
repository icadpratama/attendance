package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/icadpratama/attendance/internal/handlers"
	"github.com/icadpratama/attendance/pkg/utils"
)

var host, port string

func init() {
	host = utils.MustGet("GQL_SERVER_HOST")
	port = utils.MustGet("GQL_SERVER_PORT")
}

func Run() {
	r := gin.Default()
	r.GET("/ping", handlers.Ping())

	log.Println("Running @ http://" + host + ":" + port)
	log.Fatalln(r.Run(host + ":" + port))
}
