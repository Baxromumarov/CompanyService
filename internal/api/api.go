package api

import (
	"github.com/baxromumarov/CompanyService/config"
	v1 "github.com/baxromumarov/CompanyService/internal/api/handlers/v1"
	"github.com/baxromumarov/CompanyService/internal/auth"
	"github.com/baxromumarov/CompanyService/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type RouterOptions struct {
	Log logger.Logger
	Cfg *config.Config
	Db  *sqlx.DB
}

func New(opt *RouterOptions) *gin.Engine {
	router := gin.Default()
	// companyHandler := handler.NewCompanyHandler(companyService)

	router.Use(CORSMiddleware())

	options := &v1.HandlerV1Options{
		Log: opt.Log,
		Cfg: opt.Cfg,
		Db:  opt.Db,
	}

	handlerV1 := v1.New(options)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1Group := router.Group("/api/v1")

	v1Group.Use(auth.Authentication(opt.Cfg))

	v1Group.POST("/companies", handlerV1.CreateCompany)
	v1Group.GET("/companies/:id", handlerV1.GetCompany)
	v1Group.PATCH("/companies/:id", handlerV1.UpdateCompany)
	v1Group.DELETE("/companies/:id", handlerV1.DeleteCompany)
	// v1Group.GET("/companies", handlerV1.GetCompanies)

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, HEAD, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
