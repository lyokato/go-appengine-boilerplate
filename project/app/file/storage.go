package file

import (
	"fmt"
	"io"

	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
	"google.golang.org/appengine/blobstore"
	"google.golang.org/appengine/image"
)

func UploadFileToGCS(c context.Context,
	bucketName, objectName string,
	contentType string, file io.Reader) error {

	client, err := storage.NewClient(c)
	if err != nil {
		return err
	}
	writer := client.Bucket(bucketName).Object(objectName).NewWriter(c)
	if err != nil {
		return err
	}
	writer.ContentType = contentType
	io.Copy(writer, file)
	return writer.Close()
}

func GetFileURL(c context.Context,
	bucketName, objectName string) (string, error) {
	path := fileBlobPath(bucketName, objectName)
	key, err := blobstore.BlobKeyForFile(c, path)
	opts := image.ServingURLOptions{Secure: true, Crop: true}
	url, err := image.ServingURL(c, key, &opts)
	if err != nil {
		return "", err
	}
	return url.String(), nil
}

func fileBlobPath(bucketName, objName string) string {
	return fmt.Sprintf("/gs/%s/%s", bucketName, objName)
}
