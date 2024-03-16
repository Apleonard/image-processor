package router

import (
	"image-processor/config"
	"log"

	"image-processor/server/middleware"

	imagehandler "image-processor/api/image_processor"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
	cfg    *config.Config
	img    imagehandler.ImgHandler
}

func NewRouter(
	router *gin.Engine,
	cfg *config.Config,
	imagehandler imagehandler.ImgHandler,

) *Router {
	return &Router{
		router: router,
		cfg:    cfg,
		img:    imagehandler,
	}
}

func (r *Router) Run(port string) {
	r.setRoutes()
	r.run(port)
}

func (r *Router) setRoutes() {
	r.router.Use(middleware.CorsMiddleware())
	v1Router := r.router.Group("/api/v1")
	v1Router.Use(middleware.GinRecoveryMiddleware())

	v1Router.POST("/img", r.img.ImageProcess)

}

func (r *Router) run(port string) {
	if port == "" {
		port = "8080"
	}
	log.Printf("running on port ::[:%v]", port)
	if err := endless.ListenAndServe(":"+port, r.router); err != nil {
		log.Fatalf("failed to run on port [::%v]", port)
	}
}
