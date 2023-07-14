package main

import (
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"github.com/taepunphu/go-rest-api-structure/configs"
	"github.com/taepunphu/go-rest-api-structure/controllers"
	"github.com/taepunphu/go-rest-api-structure/database"
	"github.com/taepunphu/go-rest-api-structure/repositories"
	"github.com/taepunphu/go-rest-api-structure/routes"
	"github.com/taepunphu/go-rest-api-structure/services"
	"github.com/taepunphu/go-rest-api-structure/utils"
)

func main() {

	log.Info().Msg("Starting Server at http://localhost:4000")

	configPath := utils.GetConfigPath(os.Getenv("config"))

	cfgFile, err := configs.LoadConfig(configPath)
	if err != nil {
		log.Err(err).Msgf("LoadConfig: %v", err)
	}

	cfg, err := configs.ParseConfig(cfgFile)
	if err != nil {
		log.Err(err).Msgf("ParseConfig: %v", err)
	}

	// Database
	psqlDB, err := database.NewPsqlDB(cfg)
	if err != nil {
		log.Err(err).Msgf("Postgresql init: %s", err)
	} else {
		log.Err(err).Msgf("Postgres connected, Status: %#v", psqlDB)
	}

	validate := validator.New()

	//psqlDB.DB().Table("tb_ms_products").AutoMigrate(&entities.ProductEntity)

	// Repository
	productRepository := repositories.NewProductRepository(psqlDB)

	// Service
	productService := services.NewProductservice(productRepository, validate)

	// Controller
	productController := controllers.NewProductController(productService)

	// Router
	routes := routes.NewProductRoute(productController)

	server := &http.Server{
		Addr:    ":4000",
		Handler: routes,
	}

	errList := server.ListenAndServe()
	utils.ErrorPanic(errList)
}
