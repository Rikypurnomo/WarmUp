package config

import (
	"github.com/gin-contrib/cors"
)

const (
	Development = "DEV"
	Production  = "PRODUTION"
	Staging     = "STAGING"

	// Log Type
	LogTypeJson = "json"
	LogTypeText = "text"

	// Error Message Default
	MessageErrorDefault   = "Something went wrong"
	MessageSuccessDefault = "Success"
)

var Cors = cors.Config{
	AllowMethods:     appConfig.Cors.Methods,
	AllowHeaders:     appConfig.Cors.Headers,
	ExposeHeaders:    appConfig.Cors.ExposeHeader,
	AllowCredentials: appConfig.Cors.AllowCredentials,
	AllowOrigins:     appConfig.Cors.Origins,
	AllowAllOrigins:  true,
	MaxAge:           GetMaxAgeDuration(),
}
