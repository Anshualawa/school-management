package main

import (
	"log"
	"net/http"

	"github.com/Anshualawa/school-management/internal/auth"
	"github.com/gin-gonic/gin"
)

var books = []string{"Hindi", "English", "Social Science", "Science", "Mathematics"}

func main() {
	g := gin.Default()

	// Load HTML templates
	g.LoadHTMLGlob("templates/*.html")

	g.GET("/login", func(ctx *gin.Context) { ctx.HTML(http.StatusOK, "login.html", nil) })
	g.POST("/login", auth.Login)

	g.GET("/home", func(ctx *gin.Context) {

		// read jwt from cookie
		token, err := ctx.Cookie("access_token")
		if err != nil {
			ctx.Redirect(http.StatusFound, "/signup")
			return
		}

		claims, err := auth.ParseToken(token)
		if err != nil {
			ctx.Redirect(http.StatusFound, "/signup")
			return
		}

		data := gin.H{
			"name":  claims.Name,
			"email": claims.Email,
			"role":  claims.Role}

		ctx.HTML(http.StatusOK, "home.html", data)
	})

	g.GET("/signup", func(ctx *gin.Context) { ctx.HTML(http.StatusOK, "signup.html", nil) })
	g.POST("/signup", auth.Signup)

	g.GET("/book-list", func(ctx *gin.Context) { ctx.HTML(http.StatusOK, "index.html", books) })

	log.Fatal(g.Run(":8080"))
}
