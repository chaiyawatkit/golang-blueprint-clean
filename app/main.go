package main

import (
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang-blueprint-clean/app/database"
	"golang-blueprint-clean/app/env"
	backofficeRepo "golang-blueprint-clean/app/layers/repositories/back_office"
	backOfficeUseCase "golang-blueprint-clean/app/layers/usecases/back_office"

	_ "golang-blueprint-clean/app/docs"
	_healthcheck "golang-blueprint-clean/app/layers/deliveries/http/health_check"
	usersRepo "golang-blueprint-clean/app/layers/repositories/users"
	usersUseCase "golang-blueprint-clean/app/layers/usecases/users"

	backOfficeHandler "golang-blueprint-clean/app/layers/deliveries/http/back_office"
	usersHandler "golang-blueprint-clean/app/layers/deliveries/http/users"
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
	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//ginEngine.Use(ginney.LogWithCorrelationIdMiddleware(gin.DefaultWriter, constants.NoLoggedRoutes))
	//ginEngine.Use(ginney.MicroServiceCorrelationIdMiddleware())
	//ginEngine.Use(ginney.FromGinContextToContextMiddleware())

	dbConn := database.ConnectDB()
	defer func() {
		_ = dbConn.Close()
	}()

	//database.DBMigration()

	backofficeRepo := backofficeRepo.InitRepo(dbConn)
	usersRepo := usersRepo.InitRepo(dbConn)

	backOfficeUseCase := backOfficeUseCase.InitUseCase(backofficeRepo)
	usersUseCase := usersUseCase.InitUseCase(usersRepo)

	backOfficeHandler.NewEndpointHttpHandler(ginEngine, backOfficeUseCase)
	usersHandler.NewEndpointHttpHandler(ginEngine, usersUseCase)

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
