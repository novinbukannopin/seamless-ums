package cmd

import (
	"log"
	"seamless-ums/helpers"
	"seamless-ums/internal/api"
	"seamless-ums/internal/interfaces"
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
	userV1.DELETE("/logout", di.MiddlewareValidateAuth, di.LogoutAPI.Logout)

	err := r.Run(":" + helpers.GetEnv("PORT", ""))
	if err != nil {
		log.Fatal(err)
	}
}

type DIContainer struct {
	UserRepository interfaces.IUserRepository
	HealthCheckAPI interfaces.IHealthcheckHandler
	RegisterAPI    interfaces.IRegisterHandler
	LoginAPI       interfaces.ILoginHandler
	LogoutAPI      interfaces.ILogoutHandler
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

	logoutSvc := &services.LogoutService{
		UserRepository: userRepository,
	}

	logoutAPI := &api.LogoutHandler{
		LogoutService: logoutSvc,
	}

	return DIContainer{
		UserRepository: userRepository,
		HealthCheckAPI: healthcheckAPI,
		RegisterAPI:    registerAPI,
		LoginAPI:       loginAPI,
		LogoutAPI:      logoutAPI,
	}
}
