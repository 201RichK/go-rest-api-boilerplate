package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	config "github.com/rest-api/configs"
	"github.com/rest-api/domain"
	"github.com/rest-api/infra/logger"
)

// @title Rest-API API
// @version 1.0
// @description This is Go Rest API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email contact@amospascal.me

// @license.name MIT
// @license.url https://github.com/201RichK/go-rest-api-boilerplate/LICENSE

// @host localhost:8080
// @BasePath /api/

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func Run() {
	config.Setup() // Setup env variable
	logger.Setup() // Setup app logger

	gin.SetMode(config.ServerSetting.RunMode)
	gin.ForceConsoleColor()

	db, err := config.New(config.GetDbConfig())

	// Auto-migrate database models
	if err != nil {
		logger.Errorf("Failed to connect")
	} else {
		if db == nil {
			logger.Error("db is nil")
		} else {
			if migrate := domain.MigrateModels(db); !migrate {
				panic("migration failed")
			}
		}
	}

	endPoint := fmt.Sprintf(":%s", config.ServerSetting.HttpPort)
	router := InitRouter()
	server := &http.Server{
		Addr:         endPoint,
		Handler:      router,
		ReadTimeout:  config.ServerSetting.ReadTimeout,
		WriteTimeout: config.ServerSetting.WriteTimeout,
	}

	logger.Infof("[info] start http server listening %s", endPoint)
	server.ListenAndServe()
}
