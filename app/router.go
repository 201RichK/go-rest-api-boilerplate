package app

import (
	"github.com/aro/controllers"
	"github.com/gin-contrib/cors"
	ginI18n "github.com/gin-contrib/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/aro/docs"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.LoadHTMLGlob("templates/*")

	// apply i18n middleware
	router.Use(
		ginI18n.Localize(ginI18n.WithBundle(
			&ginI18n.BundleCfg{
				DefaultLanguage:  language.English,
				FormatBundleFile: "yaml",
				AcceptLanguage:   []language.Tag{language.English, language.French},
				RootPath:         "i18n/",
				UnmarshalFunc:    yaml.Unmarshal,
			},
		)),
	)

	// Repositories

	// Controllers

	// cors config
	config := cors.Config{
		AllowHeaders: []string{
			"Signature",
			"CloudSecret",
			"ApiKey",
		},
		AllowOrigins: []string{
			"http://localhost:8080",
			"https://cloud.id30.ci",
		},
		AllowMethods: []string{
			"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS",
		},
	}
	router.Use(cors.New(config))

	g := router.Group("/")
	g.GET("/healphcheck", controllers.Healphcheck)

	return router
}
