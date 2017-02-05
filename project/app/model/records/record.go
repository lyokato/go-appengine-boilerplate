package records

import (
	"fmt"
	"io"

	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
	"google.golang.org/appengine/blobstore"
	"google.golang.org/appengine/image"
)

const (
	BucketName = ""
)

// "github.com/olahol/go-imageupload"
// img, err := imageupload.Process(c.Request, "upload")
// thum, err := imageupload.ThumbnailPNG(img, 300, 300)
// bytes.NewReader(thum.Data)
func UploadPNG(c context.Context,
	objectName string, file io.Reader) (string, error) {
	uploadPictureToGCS(c, objectName, file)
	return GetPictureURL(c, objectName)
}

func uploadPictureToGCS(c context.Context,
	objectName string, file io.Reader) error {

	client, err := storage.NewClient(c)
	if err != nil {
		return err
	}
	writer := client.Bucket(BucketName).Object(objectName).NewWriter(c)
	if err != nil {
		return err
	}
	writer.ContentType = "image/png"
	io.Copy(writer, file)
	return writer.Close()
}

func GetPictureURL(c context.Context, objectName string) (string, error) {
	key, err := blobstore.BlobKeyForFile(c, pictureBlobPath(objectName))
	opts := image.ServingURLOptions{Secure: true, Crop: true}
	url, err := image.ServingURL(c, key, &opts)
	if err != nil {
		return "", err
	}
	return url.String(), nil
}

func pictureBlobPath(objName string) string {
	return fmt.Sprintf("/gs/%s/%s", BucketName, objName)
}
