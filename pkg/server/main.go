package server

import (
	"github.com/gin-gonic/gin"
	"github.com/icadpratama/attendance/internal/handlers"
	log "github.com/icadpratama/attendance/internal/logger"
	"github.com/icadpratama/attendance/internal/orm"
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

func Run(orm *orm.ORM) {
	log.Info("GORM_CONNECTION_DSN: ", utils.MustGet("GORM_CONNECTION_DSN"))

	endpoint := "http://" + host + ":" + port

	r := gin.Default()

	r.GET("/ping", handlers.Ping())

	if isPgEnabled {
		r.GET(gqlPgPath, handlers.PlaygroundHandler(gqlPath))
		log.Info("GraphQL Playground @ " + endpoint + gqlPgPath)
	}

	r.POST(gqlPath, handlers.GraphqlHandler(orm))
	log.Info("GraphQL @ " + endpoint + gqlPath)

	log.Info("Running @ " + endpoint)
	log.Fatal(r.Run(host + ":" + port))
}
