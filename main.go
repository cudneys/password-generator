package main

import (
	"github.com/cudneys/password-generator/api"
	config "github.com/cudneys/password-generator/config"
	docs "github.com/cudneys/password-generator/docs"
	"github.com/cudneys/password-generator/models"
	monitoring "github.com/cudneys/password-generator/monitoring"
	version "github.com/cudneys/password-generator/version"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"
)

var (
	Version        = "dev"
	CommitHash     = "n/a"
	BuildTimestamp = "n/a"
)

// @title           Password Generator API
// @version         1.0
// @description     This is a simple password generator
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://github.com/cudneys
// @contact.email  password-generator@cudneys.net

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	bindHost := config.GetBindHost()
	logColorEnabled := config.GetLogColorEnabled()
	releaseMode := config.GetReleaseMode()

	r := gin.Default()

	if logColorEnabled == false {
		gin.DisableConsoleColor()
	}

	if releaseMode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	docs.SwaggerInfo.BasePath = "/v1"
	p := monitoring.NewPrometheus("password_generator")
	p.Use(r)

	v1 := r.Group("/v1")
	{
		v1.GET("/generate", api.PasswordGenerator)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/about", aboutHandler)
	r.GET("/liveness", livenessHandler)
	r.Run(bindHost)
}

func livenessHandler(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		models.Liveness{Status: "OK", Timestamp: time.Time.Format(time.Now(), "2006-01-02T15:04:05")},
	)
}

func aboutHandler(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		models.Version{
			Version:        version.Version,
			CommitHash:     version.CommitHash,
			BuildTimestamp: version.BuildTimestamp,
		})
}
