package handlers

import (
    "github.com/99designs/gqlgen/handler"
    "github.com/icadpratama/attendance/internal/gql"
    "github.com/icadpratama/attendance/internal/gql/resolvers"
    "github.com/gin-gonic/gin"
)

func GraphqlHandler() gin.HandlerFunc {
    c := gql.Config{
        Resolvers: &resolvers.Resolver{},
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