package uploadhttp

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"go01-airbnb/pkg/common"
	"go01-airbnb/pkg/upload"
	"image"
	"net/http"
)

type uploadHandler struct {
	s3Provider upload.UploadProvider
}

func NewUploadHandler(s3Provider upload.UploadProvider) *uploadHandler {
	return &uploadHandler{s3Provider}
}

func (hdl *uploadHandler) Upload() gin.HandlerFunc {
	return func(context *gin.Context) {
		fileHeader, err := context.FormFile("file")
		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		//Save file trực tiếp vào server của mình
		//if err := context.SaveUploadedFile(fileHeader, fmt.Sprintf("static/%s", fileHeader.Filename)); err != nil {
		//	panic(err)
		//}
		//context.JSON(http.StatusOK, common.Response(common.Image{
		//	Url: "http://localhost:8080/static/" + fileHeader.Filename,
		//}))

		folder := context.DefaultPostForm("folder", "img")
		fileName := fileHeader.Filename // image.png
		//fileExt := filepath.Ext(fileName) // png

		file, err := fileHeader.Open()
		defer file.Close()
		if err != nil {
			panic(common.ErrBadRequest(err))
		}

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrBadRequest(err))
		}

		//// Get width and height of image
		//fileBytes := bytes.NewBuffer(dataBytes)
		//
		//w, h, err := getImageDimension(fileBytes)
		//if err != nil {
		//	panic(common.ErrBadRequest(err))
		//}

		img, err := hdl.s3Provider.UploadFile(context.Request.Context(),
			dataBytes,
			fmt.Sprintf("%s/%s", folder, fileName))
		if err != nil {
			panic(common.ErrBadRequest(err))
		}
		//img.Width = w
		//img.Height = h

		context.JSON(http.StatusOK, common.Response(img))
	}
}

func getImageDimension(file *bytes.Buffer) (int, int, error) {
	img, _, err := image.DecodeConfig(file)
	if err != nil {
		return 0, 0, err
	}
	return img.Width, img.Height, nil
}
