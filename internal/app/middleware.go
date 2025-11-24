package app

import (
	"net/http"
	"strings"

	"github.com/Anshualawa/school-management/internal/auth"
	"github.com/Anshualawa/school-management/internal/config"
	"github.com/Anshualawa/school-management/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InjectDB(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("db", db)
		ctx.Next()
	}
}

func JWTMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if utils.IsEmpty(authHeader) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization token"})
			ctx.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization"})
			ctx.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := auth.ValidateJWT(token)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token:" + err.Error()})
			ctx.Abort()
			return
		}

		ctx.Set("userID", claims.UserID)
		ctx.Set("email", claims.Email)
		ctx.Set("user_role", claims.Role)

		ctx.Next()
	}
}
