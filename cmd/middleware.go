package cmd

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"seamless-ums/helpers"
	"time"
)

func (d *DIContainer) MiddlewareValidateAuth(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		log.Println("Authorization header is missing")
		helpers.SendResponseHTTP(c, http.StatusUnauthorized, "Authorization header is missing", nil)
		c.Abort()
		return
	}

	if d.UserRepository == nil {
		log.Println("UserRepository is nil")
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, "internal server error", nil)
		c.Abort()
		return
	}

	_, err := d.UserRepository.GetUserSessionByToken(c.Request.Context(), auth)
	if err != nil {
		log.Println("Invalid token:", err)
		helpers.SendResponseHTTP(c, http.StatusUnauthorized, "Invalid token", nil)
		c.Abort()
		return
	}

	claim, err := helpers.ValidateToken(c.Request.Context(), auth)
	if err != nil {
		log.Println("Token validation failed:", err)
		helpers.SendResponseHTTP(c, http.StatusUnauthorized, "Token validation failed", nil)
		c.Abort()
		return
	}

	if time.Now().Unix() > claim.ExpiresAt.Unix() {
		log.Println("Token has expired")
		helpers.SendResponseHTTP(c, http.StatusUnauthorized, "Token has expired", nil)
		c.Abort()
		return
	}

	c.Set("token", claim)

	c.Next()
}
