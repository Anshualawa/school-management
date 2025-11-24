package app

import (
	"github.com/Anshualawa/school-management/internal/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	G   *gin.Engine
	DB  *gorm.DB
	Cfg *config.Config
}

func NewServer(cfg *config.Config, db *gorm.DB) *Server {
	g := gin.New()

	RegisterRoutes(g, db, cfg)

	return &Server{G: g, DB: db, Cfg: cfg}
}

func (s *Server) Start(addr string) error {
	return s.G.Run(addr)
}
