package auth

import (
	"net/http"

	"github.com/Anshualawa/school-management/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Signup(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.String(http.StatusBadRequest, "Invalid form data")
		return
	}

	user.ID = uuid.New()
	user.Role = "user"

	// create JWT
	token, err := CreateToken(user.Name, user.Email, user.Role)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Token error")
		return
	}

	// set jwt in cookie
	ctx.SetCookie("access_token", token, 3600*2, "/", "", false, true)

	ctx.Redirect(http.StatusFound, "/home")
}

// temp user
var demoUser = models.User{
	Name:     "Priyanshu Alawa",
	Email:    "alawa@admin.com",
	Password: "1234",
	Role:     "user",
}

func Login(ctx *gin.Context) {
	var form models.User

	if err := ctx.ShouldBind(&form); err != nil {
		ctx.String(http.StatusBadRequest, "invalid form data")
		return
	}

	// verify credentials
	if form.Email != demoUser.Email || form.Password != demoUser.Password {
		ctx.String(http.StatusUnauthorized, "invalid email or passsword")
		return
	}

	// create jwt token
	token, err := CreateToken(demoUser.Name, demoUser.Email, demoUser.Role)

	if err != nil {
		ctx.String(http.StatusInternalServerError, "could not create token")
		return
	}

	ctx.SetCookie("access_token", token, 3600*2, "/", "", false, true)

	ctx.Redirect(http.StatusFound, "/home")

}
