package uploadhttp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go01-airbnb/pkg/common"
	"net/http"
)

type uploadHandler struct {
}

func NewUploadHandler() *uploadHandler {
	return &uploadHandler{}
}

func (hdl *uploadHandler) Upload() gin.HandlerFunc {
	return func(context *gin.Context) {
		fileHeader, err := context.FormFile("file")
		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		//file, err := fileHeader.Open()
		//if err != nil {
		//	panic(common.ErrBadRequest(err))
		//}
		//defer file.Close()
		if err := context.SaveUploadedFile(fileHeader, fmt.Sprintf("static/%s", fileHeader.Filename)); err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, common.Response(common.Image{
			Url: "http://localhost:8080/static/" + fileHeader.Filename,
		}))
	}
}
