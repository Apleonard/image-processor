package server

import (
	"image-processor/config"
	"image-processor/server/dependency"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	cfg *config.Config
}

func NewHttpServer() *Server {
	server := &Server{}
	server.cfg = config.Get()
	return server
}

func (s *Server) Run() {
	//setMode(s.cfg)

	router := gin.Default()
	dependency.InitializeRouter(router).Run(s.cfg.AppConfig.Port)
}

func setMode(cfg *config.Config) {
	appEnv := cfg.AppConfig.Environment
	if appEnv == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}
}
