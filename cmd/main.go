package main

import (
	"Training/go-crud-with-oracle/docs"
	"Training/go-crud-with-oracle/infrastructure/persistance/usecase"
	"Training/go-crud-with-oracle/interface/handler"
	"Training/go-crud-with-oracle/technical_service/config"
	"Training/go-crud-with-oracle/technical_service/database"
	"Training/go-crud-with-oracle/technical_service/log"
	"Training/go-crud-with-oracle/technical_service/security/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	config.SetupConfiguration()
}

// @termsOfService http://swagger.io/terms/
// @contact.name Muhammad Rais Adlani
// @contact.url https://gitlab.com/mraisadlani
// @contact.email mraisadlani@gmail.com

// @license.name MIT
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	docs.SwaggerInfo.Title = "Go Oracle"
	docs.SwaggerInfo.Description = "Gin CRUD with Database Oracle"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%v", config.C.App.HOST, config.C.App.PORT)
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	logs := log.SetupLog()

	db, err := database.InitDB()
	defer db.Close()

	if err != nil {
		logs.Errorf("Database Error : %s", err.Error())
		return
	}

	ping := db.Ping()

	if ping != nil {
		logs.Errorf( "Request Timeout : %s", err.Error())
		return
	} else {
		logs.Info("Connected Database")
	}

	port := config.C.App.PORT

	if port == 7000 {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	BaseConfiguration(logs, db)
}

func BaseConfiguration(logs *logrus.Logger, db *database.Database) {
	userCase := usecase.BuildUserUsecase(db.UserRepo)

	r := gin.New()

	r.Use(log.Logger(logs), gin.Recovery())
	r.Use(middleware.SetupCorsMiddleware())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PONG",
		})
	})

	api := r.Group("/api/v1")

	handler.BuildHandler(api, userCase)

	logs.Info("Listening on port ", config.C.App.PORT)

	err := r.Run(fmt.Sprintf(":%v", config.C.App.PORT))

	if err != nil {
		logs.Errorf("Error listening port server : %v", err)
		return
	}
}