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
	healthcheckSvc := &services.Healthcheck{}
	healthcheckAPI := &api.Healthcheck{
		HealthcheckServices: healthcheckSvc,
	}

	registerRepo := &repository.RegisterRepository{
		DB: helpers.DB,
	}

	registerSvc := &services.RegisterService{
		RegisterRepo: registerRepo,
	}

	registerAPI := &api.RegisterHandler{
		RegisterService: registerSvc,
	}

	r := gin.Default()

	r.GET("/health", healthcheckAPI.HealthcheckHandlerHTTP)

	userV1 := r.Group("/user/v1")
	userV1.POST("/register", registerAPI.Register)

	err := r.Run(":" + helpers.GetEnv("PORT", ""))
	if err != nil {
		log.Fatal(err)
	}
}
