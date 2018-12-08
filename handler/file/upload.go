package file

import (
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go"
	"github.com/spf13/viper"
	. "git.zjuqsc.com/rop/ROP-go/handler"
	"git.zjuqsc.com/rop/ROP-go/pkg/errno"
	"fmt"
	"github.com/satori/go.uuid"
	"path"
)

func UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		SendResponse(c, errno.ErrFileRead, "请上传有效的文件~")
		return
	}

	contentType := file.Header["Content-Type"][0]
	if contentType != "image/png" && contentType != "image/jpeg" {
		SendResponse(c, errno.ErrFileType, "需要以.png或者.jpg类型的图片~")
		return
	}


	//log.Debugf("%d", file.Size)

	fileSize := file.Size
	if fileSize > (1 << 20) {
		SendResponse(c, errno.ErrFileSize, "图片的最大尺寸为1MB哦~~")
		return
	}

	f, err := file.Open()
	if err != nil {
		SendResponse(c, errno.ErrFileRead, err.Error())
		return
	}

	UUID := uuid.NewV4()

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