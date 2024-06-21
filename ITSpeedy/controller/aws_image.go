package controller

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"templategoapi/db"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func HandleUpload(resource *db.Resource) func(c *gin.Context) {

	return func(c *gin.Context) {
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
			return
		}
		path := c.Params.ByName("path")
		subpath := c.Request.URL.Query().Get("subpath")
		filename := header.Filename

		fileBytes, err := ioutil.ReadAll(file)
		pathurl, err := UploadFile(filename, fileBytes, path, subpath)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "อัพโหลดไม่สำเร็จ"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "สำเร็จ", "payload": pathurl})
	}
}

const (
	AWS_S3_REGION         = "1"
	AWS_S3_BUCKET         = "1"
	AWS_ACCESS_KEY_ID     = "1"
	AWS_SECRET_ACCESS_KEY = "1"
)

func UploadFile(fileName string, fileData []byte, service string, path string) (string, error) {
	session, err := session.NewSession(&aws.Config{
		Region:      aws.String(AWS_S3_REGION),
		Credentials: credentials.NewStaticCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, ""),
	})
	if err != nil {
		logrus.Fatal(err)
	}
	_, err = s3.New(session).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(AWS_S3_BUCKET),
		Key:                  aws.String(service + "/" + path + "/" + fileName),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(fileData),
		ContentType:          aws.String(http.DetectContentType(fileData)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	if err != nil {
		return "", err
	}
	urlpath := "https://itspeedy-image.s3.ap-southeast-1.amazonaws.com/" + service + "/" + path + "/" + fileName
	return urlpath, err
}
