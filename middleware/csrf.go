package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
	adapter "github.com/gwatts/gin-adapter"
)

func CSRFMiddleware(router *gin.Engine) {
	router.Use(CSRF())
	router.Use(CSRFToken())
}

func CSRF() gin.HandlerFunc {
	csrfMiddleware := csrf.Protect([]byte("32-byte-long-auth-key"))
	return adapter.Wrap(csrfMiddleware)
}

func CSRFToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-CSRF-Token", csrf.Token(c.Request))
	}
}
