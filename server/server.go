package server

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	ENV_STAGING    = "staging"
	ENV_PRODUCTION = "production"
	ENV_DEV        = "dev"
)

type Server struct {
	db     *gorm.DB
	router *gin.Engine
}

func New(db *gorm.DB, appEnv string) *Server {
	s := &Server{db: db}

	fmt.Printf("%+v\n", s)

	if appEnv == ENV_PRODUCTION {
		gin.SetMode(gin.ReleaseMode)
	}

	s.router = gin.Default()
	s.setRoutes()

	return s
}

func (s *Server) Run(port int) {
	s.router.Run(":" + strconv.Itoa(port))
}
