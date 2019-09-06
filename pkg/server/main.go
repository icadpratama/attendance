package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/icadpratama/attendance/internal/handlers"
	"github.com/icadpratama/attendance/pkg/utils"
)

var host, port, gqlPath, gqlPgPath string
var isPgEnabled bool

func init() {
	host = utils.MustGet("GQL_SERVER_HOST")
	port = utils.MustGet("GQL_SERVER_PORT")
	gqlPath = utils.MustGet("GQL_SERVER_GRAPHQL_PATH")
	gqlPgPath = utils.MustGet("GQL_SERVER_GRAPHQL_PLAYGROUND_PATH")
	isPgEnabled = utils.MustGetBool("GQL_SERVER_GRAPHQL_PLAYGROUND_ENABLED")
}

func Run() {
	endpoint := "http://" + host + ":" + port

	r := gin.Default()

	r.GET("/ping", handlers.Ping())

	if isPgEnabled {
		r.GET(gqlPgPath, handlers.PlaygroundHandler(gqlPath))
		log.Println("GraphQL Playground @ " + endpoint + gqlPgPath)
	}
	r.POST(gqlPath, handlers.GraphqlHandler())
	log.Println("GraphQL @ " + endpoint + gqlPath)

	log.Println("Running @ " + endpoint)
	log.Fatalln(r.Run(host + ":" + port))
}
