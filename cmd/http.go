package cmd

import (
	"log"
	"seamless-ums/helpers"
	"seamless-ums/internal/api"
	"seamless-ums/internal/repository"
	"seamless-ums/internal/services"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	di := DI()

	r := gin.Default()

	r.GET("/health", di.HealthCheckAPI.HealthcheckHandlerHTTP)

	userV1 := r.Group("/user/v1")
	userV1.POST("/register", di.RegisterAPI.Register)
	userV1.POST("/login", di.LoginAPI.Login)

	err := r.Run(":" + helpers.GetEnv("PORT", ""))
	if err != nil {
		log.Fatal(err)
	}
}

type DIContainer struct {
	HealthCheckAPI *api.Healthcheck
	RegisterAPI    *api.RegisterHandler
	LoginAPI       *api.LoginHandler
}

func DI() DIContainer {
	healthcheckSvc := &services.Healthcheck{}
	healthcheckAPI := &api.Healthcheck{
		HealthcheckServices: healthcheckSvc,
	}

	userRepository := &repository.UserRepository{
		DB: helpers.DB,
	}

	registerSvc := &services.RegisterService{
		UserRepository: userRepository,
	}

	registerAPI := &api.RegisterHandler{
		RegisterService: registerSvc,
	}

	loginSvc := &services.LoginService{
		UserRepository: userRepository,
	}

	loginAPI := &api.LoginHandler{
		LoginService: loginSvc,
	}

	return DIContainer{
		HealthCheckAPI: healthcheckAPI,
		RegisterAPI:    registerAPI,
		LoginAPI:       loginAPI,
	}
}
