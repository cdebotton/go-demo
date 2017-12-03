package main

import (
	"log"

	"github.com/cdebotton/obscura/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/graphql", handler.Query)
	r.GET("/graphiql", handler.GraphiQL)

	log.Fatal(r.Run())
}
