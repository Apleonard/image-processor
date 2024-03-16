package imageprocessor

import (
	imageservice "image-processor/internal/service/image_service"
	appError "image-processor/pkg/error_util"
	appHttp "image-processor/pkg/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ImgHandler struct {
	imgService imageservice.ImageService
}

func NewImageHandler(imgService imageservice.ImageService) *ImgHandler {
	return &ImgHandler{imgService: imgService}
}

func (i *ImgHandler) ImageProcess(ctx *gin.Context) {

	err := ctx.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		appHttp.ResponseBadRequest(ctx, appError.HandlerError("",
			"ctx.ShouldBindJSON", err))
		return
	}

	// Retrieve file from form data
	desiredWidth, _ := strconv.Atoi(ctx.Request.FormValue("width"))
	desiredHeight, _ := strconv.Atoi(ctx.Request.FormValue("height"))
	file, _, err := ctx.Request.FormFile("file")
	if err != nil {
		appHttp.ResponseBadRequest(ctx, appError.HandlerError("",
			"ctx.ShouldBindJSON", err))
		return
	}
	defer file.Close()

	outputBuf, err := i.imgService.ImageProcessesService(file, desiredWidth, desiredHeight)
	if err != nil {
		appHttp.ResponseBadRequest(ctx, appError.HandlerError("",
			"ctx.ShouldBindJSON", err))
		return
	}

	ctx.Header("Content-Type", "image/jpeg")
	ctx.Data(200, "image/jpeg", outputBuf.Bytes())
}
