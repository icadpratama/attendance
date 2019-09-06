package handlers

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/icadpratama/attendance/internal/gql"
	"github.com/icadpratama/attendance/internal/orm"
	"github.com/icadpratama/attendance/internal/gql/resolvers"
)

func GraphqlHandler(orm *orm.ORM) gin.HandlerFunc {
	c := gql.Config{
		Resolvers: &resolvers.Resolver{
			ORM: orm,
		},
	}

	h := handler.GraphQL(gql.NewExecutableSchema(c))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func PlaygroundHandler(path string) gin.HandlerFunc {
	h := handler.Playground("Go GraphQL Server", path)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}