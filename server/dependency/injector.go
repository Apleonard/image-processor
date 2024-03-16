package dependency

import (
	"image-processor/config"
	"image-processor/server/router"

	imagehandler "image-processor/api/image_processor"
	imageservice "image-processor/internal/service/image_service"

	"github.com/gin-gonic/gin"
)

func InitializeRouter(r *gin.Engine) *router.Router {

	var (
		cfg = config.Get()
	)

	// repo
	var ()

	//service
	var (
		imgSrv = imageservice.NewImageService()
	)

	// handler
	var (
		imgHndl = imagehandler.NewImageHandler(imgSrv)
	)

	routerRouter := router.NewRouter(r, cfg, *imgHndl)
	return routerRouter

}
