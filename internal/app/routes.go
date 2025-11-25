package app

import (
	"time"

	"github.com/Anshualawa/school-management/internal/config"
	"github.com/Anshualawa/school-management/internal/handlers"
	"github.com/Anshualawa/school-management/internal/repositories"
	"github.com/Anshualawa/school-management/internal/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB, cfg *config.Config) {
	r.Use(InjectDB(db))

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// repository or services call

	userRepo := repositories.NewUserRepository(db)
	userSvc := services.NewUserService(userRepo)
	authH := handlers.NewAuthorHandler(userSvc)

	
	r.POST("/api/v1/signup", authH.Signup)
	r.POST("/api/v1/login", authH.Login)


}
