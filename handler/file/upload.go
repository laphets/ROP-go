package file

import (
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go"
	"github.com/spf13/viper"
	. "rop/handler"
	"rop/pkg/errno"
	"fmt"
	"github.com/satori/go.uuid"
	"path"
	"github.com/lexkong/log"
)

func UploadImage(c *gin.Context) {
	file, _ := c.FormFile("file")

	contentType := file.Header["Content-Type"][0]
	if contentType != "image/png" && contentType != "image/jpeg" {
		SendResponse(c, errno.ErrFileType, "Need *.png or *.jpg")
		return
	}


	log.Debugf("%d", file.Size)

	fileSize := file.Size
	if fileSize > (1 << 20) {
		SendResponse(c, errno.ErrFileSize, "Maxium file size is 1MB")
		return
	}

	f, err := file.Open()
	if err != nil {
		SendResponse(c, errno.ErrFileRead, err.Error())
		return
	}

	UUID, err := uuid.NewV4()

	fileName := fmt.Sprintf("QSC-IMG-%s%s", UUID.String(), path.Ext(file.Filename))
	minioClient, err := minio.New(viper.GetString("minio.endpoint"), viper.GetString("minio.accessKeyID"), viper.GetString("minio.secretAccessKey"), viper.GetBool("minio.useSSL"))
	_, err = minioClient.PutObject(viper.GetString("minio.bucket"), fileName, f, file.Size,  minio.PutObjectOptions{ContentType: file.Header["Content-Type"][0]})
	if err != nil {
		SendResponse(c, errno.ErrOSS, err.Error())
		return
	}

	//log.Debugf("%d ", n)
	//log.Debugf("%d %s ... %+v", file.Size, file.Header["Content-Type"][0], file)
	SendResponse(c, nil, fmt.Sprintf("http://%s/%s/%s", viper.GetString("minio.endpoint"), viper.GetString("minio.bucket"), fileName))
}