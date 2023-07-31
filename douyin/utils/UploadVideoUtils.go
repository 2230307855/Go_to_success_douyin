package utils

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"os"
	"strings"
)

// 上传文件到本地的中间件
// 返回上传文件的文件绝对路径和文件名
func SingleFileUploadMidWare(c *gin.Context) (filePath string, fileName string) {
	projectPath, _ := os.Getwd()
	//直接从formfile中获得文件
	file, _ := c.FormFile("data")
	filename := file.Filename
	extename := filename[strings.Index(filename, "."):]
	uuid := uuid.New()
	filename = uuid.String() + extename
	log.Printf(filename)
	//"E:/webServer/golang/file_test/upload"+"/images/"绝对路径测试
	c.SaveUploadedFile(file, projectPath+"/upload/videos/"+filename)
	return projectPath + "/upload/videos/" + filename, filename
}

// 上传视频到服务器并返回视频的地址
// 最终返回服务器可访问的地址
func UploadVideo(filePath string, filename string) string {
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	client, err := oss.New(EndPoint, AccessKeyId, AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// 填写存储空间名称，例如examplebucket。
	bucket, err := client.Bucket(BurketName)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = bucket.PutObjectFromFile(OssFilePath+filename, filePath)
	DeleteFile()
	// 依次填写Object的完整路径（例如exampledir/exampleobject.txt）和本地文件的完整路径（例如D:\\localpath\\examplefile.txt）。
	if err != nil {
		fmt.Println("Error:", err)
	}
	return OssVisitPath + filename
}
