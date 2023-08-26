package routers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"go-gin-example/conf"
	"go-gin-example/core"
	"go-gin-example/models"

	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	arr, err := core.FindData()
	if err != nil {
		fmt.Print(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "查询学生信息成功",
		"data":    arr,
	})
}

func TaskByIdentifierHandler(c *gin.Context) {
	identifier := c.Param("identifier")
	data, err := core.GetByIdentifier(identifier)
	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "成功",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "查询学生信息成功",
		"data":    data,
	})
}

func TestUploadToMinIO(c *gin.Context) {
	objectKey := "富邦财险_新核心系统实施项目_反洗钱系统设计说明书_v0.2_20230228.doc"
	filePath := "C:/Home/File/huangqi/work/file/富邦财险_新核心系统实施项目_反洗钱系统设计说明书_v0.2_20230228.doc"
	core.UploadToMinIO(objectKey, filePath)
	c.JSON(http.StatusOK, gin.H{
		"message": "上传完成文件成功",
		"code":    "000",
	})
}

func GetStsUploadCredentials(c *gin.Context) {
	cre, err := core.GetStsUploadCredentials()
	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取临时凭证失败",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "获取临时凭证成功",
		"code":    "0000",
		"data":    cre,
	})
}

func GeneratePresignedURL(c *gin.Context) {
	expiration := time.Hour
	objectKey := c.Query("objectKey")
	url, err := conf.GeneratePresignedURL(objectKey, expiration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "获取临时签名成功",
		"code":    200000,
		"data":    url,
	})
}

func InitTask(c *gin.Context) {
	var param models.InitTaskParam

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("%+v\n", param)
	data, err := core.InitTask(param)
	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "失败",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "成功",
		"code":    200000,
		"data":    data,
	})
}

func GenPreSignUploadURL(c *gin.Context) {
	identifier := c.Param("identifier")
	task, err := core.GetByIdentifier(identifier)
	acc := models.SysUploadTask{}
	if task == acc || err != nil {
		fmt.Print(task)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "分片任务不存在",
			"error":   200001,
		})
		return
	}
	partNumber := c.Param("partNumber")
	number, err := strconv.ParseInt(partNumber, 10, 64)
	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "失败",
			"error":   err.Error(),
		})
		return
	}

	data, err := core.GenPreSignUploadURL(task.ObjectKey, task.UploadID, number)
	// 使用ShouldBindJSON方法将请求体中的JSON数据绑定到param变量
	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "失败",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "成功",
		"code":    200000,
		"data":    data,
	})
}

func GetTaskInfo(c *gin.Context) {
	identifier := c.Param("identifier")
	data := core.GetTaskInfo(identifier)
	// 使用ShouldBindJSON方法将请求体中的JSON数据绑定到param变量

	c.JSON(http.StatusOK, gin.H{
		"message": "成功",
		"code":    200000,
		"data":    data,
	})
}
func Merge(c *gin.Context) {
	identifier := c.Param("identifier")
	core.Merge(identifier)
	// 使用ShouldBindJSON方法将请求体中的JSON数据绑定到param变量

	c.JSON(http.StatusOK, gin.H{
		"message": "成功",
		"code":    200000,
		"data":    nil,
	})
}
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Next()

	})
	// router.LoadHTMLGlob("./templates/*")
	// router.GET("/demo", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", gin.H{})
	// })
	router.GET("/GetTasks", PingHandler)
	router.GET("test_upload", TestUploadToMinIO)
	router.GET("getStsUploadCredentials", GetStsUploadCredentials)
	router.GET("generatePresignedURL", GeneratePresignedURL)
	router.POST("/v1/minio/tasks", InitTask)
	router.GET("/v1/minio/tasks/:identifier/:partNumber", GenPreSignUploadURL)
	router.GET("/v1/minio/tasks/:identifier", GetTaskInfo)
	router.POST("/v1/minio/tasks/merge/:identifier", Merge)

	return router
}
