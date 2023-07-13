package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	utils "github.com/taepunphu/go-rest-api-structure/utils/errors"
)

func main() {

	log.Info().Msg("Starting Server at http://localhost:4000")
	routes := gin.Default()

	routes.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "hello world!")
	})

	server := &http.Server{
		Addr: ":4000",
		Handler: routes,
	}

	err := server.ListenAndServe()
	utils.ErrorPanic(err)
}