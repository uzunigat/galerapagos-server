package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var corsConfig = cors.Config{
	AllowOrigins:     []string{"*"},
	AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	AllowHeaders:     []string{"*"},
	AllowCredentials: true,
}

func Cors() gin.HandlerFunc {
	return cors.New(corsConfig)
}
