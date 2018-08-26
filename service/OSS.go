package service

//import (
//	"github.com/minio/minio-go"
//	"github.com/spf13/viper"
//)
//
//func SaveFileOSS(objectName, filePath string) (int64, error) {
//	minioClient, err := minio.New(viper.GetString("minio.endpoint"), viper.GetString("minio.accessKeyID"), viper.GetString("minio.secretAccessKey"), viper.GetBool("minio.useSSL"))
//	n, err := minioClient.FPutObject(viper.GetString("minio.bucket"), objectName, filePath, minio.PutObjectOptions{})
//	minioClient.PutObject()
//	if err != nil {
//		return 0, err
//	}
//	return n, nil
//}