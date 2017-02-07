package controller_helper

import (
	"app/file"

	"bytes"
	"io"
	"net/http"

	imageupload "github.com/olahol/go-imageupload"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func UploadImage(c *gin.Context,
	imageRequestKey, bucketName, objectName string,
	width, height int) (string, error) {

	ctx := CTX(c)

	data, err := createThumbnailFromRequest(c.Request,
		imageRequestKey, width, height)
	if err != nil {
		return "", err
	}

	err = file.UploadFileToGCS(ctx, bucketName, objectName,
		"image/png", data)
	if err != nil {
		return "", err
	}

	return file.GetFileURL(ctx, bucketName, objectName)
}

func createThumbnailFromRequest(r *http.Request,
	imageRequestKey string, width, height int) (io.Reader, error) {

	img, err := imageupload.Process(r, imageRequestKey)
	if err != nil {
		return nil, err
	}

	thum, err := imageupload.ThumbnailPNG(img, width, height)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(thum.Data), nil
}
