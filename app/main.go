package main

import (
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang-blueprint-clean/app/env"
	backofficeRepo "golang-blueprint-clean/app/layers/repositories/back_office"
	backOfficeUseCase "golang-blueprint-clean/app/layers/usecases/back_office"

	_ "golang-blueprint-clean/app/docs"
	backOfficeHandler "golang-blueprint-clean/app/layers/deliveries/http/back_office"
	_healthcheck "golang-blueprint-clean/app/layers/deliveries/http/health_check"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pseidemann/finish"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Trainning Golang MicroService
// @version 0.1.0
// @description Trainning Golang MicroService.
// @schemes http https
// @BasePath /

func main() {
	env.Init()
	//boredom.InitL(env.LogLevel)
	ginEngine := gin.New()
	ginEngine.Use(gin.Recovery())
	_healthcheck.NewEndpointHTTPHandler(ginEngine)

	ginEngine.Use(CORSMiddleware())
	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//ginEngine.Use(ginney.LogWithCorrelationIdMiddleware(gin.DefaultWriter, constants.NoLoggedRoutes))
	//ginEngine.Use(ginney.MicroServiceCorrelationIdMiddleware())
	//ginEngine.Use(ginney.FromGinContextToContextMiddleware())

	backofficeRepo := backofficeRepo.InitRepo()
	backOfficeUseCase := backOfficeUseCase.InitUseCase(backofficeRepo)

	//middleware := _middlewareHttp.InitAuthMiddleware(nil)
	backOfficeHandler.NewEndpointHttpHandler(ginEngine, backOfficeUseCase)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: ginEngine,
	}

	timeOut, err := strconv.Atoi(os.Getenv("GRACEFUL_TIMEOUT"))
	if err != nil {
		timeOut = 30 // second
	}

	graceful := &finish.Finisher{Timeout: time.Duration(timeOut) * time.Second}
	graceful.Add(srv)

	go func() {
		err := srv.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	graceful.Wait()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-correlation-ID, x-correlation-id, x-finplus-auth")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
