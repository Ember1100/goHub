package core

import (
	"errors"
	"fmt"
	"go-gin-example/conf"
	"go-gin-example/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindData(c *gin.Context) ([]models.SysUploadTask, error) {
	db, ok := c.Get("db")
	if !ok {
		return nil, errors.New("database connection not found in context")
	}
	gormDB, ok := db.(*gorm.DB)
	if !ok {
		return nil, errors.New("invalid database connection type")
	}

	var tasks []models.SysUploadTask
	result := gormDB.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func GetByIdentifier(identifier string) (models.SysUploadTask, error) {
	db, err := conf.ConnectToDatabase()
	if err != nil {
		fmt.Print(err)
	}
	var task models.SysUploadTask
	result := db.Where("file_identifier = ?", identifier).First(&task)
	fmt.Println(result)
	if result == nil {
		return task, err
	}
	return task, nil
}

func InsertUploadTask(param models.SysUploadTask) (string, error) {
	db, err := conf.ConnectToDatabase()
	if err != nil {
		fmt.Print(err)
	}
	result := db.Create(&param)
	if result.Error != nil {
		fmt.Print(result.Error)
		return "", result.Error
	}
	return "ok", nil
}
