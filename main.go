package main

import (
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte(os.Getenv("APP_KEY")))
	r.LoadHTMLGlob("views/layouts/*.html")
	r.LoadHTMLGlob("views/*.html")

	r.Use(sessions.Sessions("gin-session", store))
	r.Use(csrf.Middleware(csrf.Options{
		Secret: os.Getenv("APP_KEY"),
		ErrorFunc: func(c *gin.Context) {
			c.String(400, "CSRF token mismatch")
			c.Abort()
		},
	}))

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{
			"message": "hello world",
		})
	})

	r.Run("0.0.0.0:8000")
}
