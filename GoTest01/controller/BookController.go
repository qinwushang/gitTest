package controller

import (
	"GoTest01/dao"
	"GoTest01/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"strconv"
)

func InitHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "Klibrary.html", nil)
}

func UploadFile(c *gin.Context) { //文件新建上传
	var book models.Book
	c.ShouldBind(&book)
	f, _ := c.FormFile("file")
	c.JSON(200, gin.H{
		"name": book.Name,
		"url":  book.Url,
		"type": book.Type,
		"des":  book.Description,
	})
	dao.Db.Create(&book)
	println(f.Filename)
	err := c.SaveUploadedFile(f, "./bookRealUrl/"+f.Filename)
	if err != nil {
		panic(err)
		return
	}

	//Type:        c.PostForm("type"),
	//Name:        c.PostForm("name"),
	//Description: c.PostForm("description"),
	//Url:         c.PostForm("url"),
}

func GetAll(c *gin.Context) {
	var book = []models.Book{}
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	var total int64
	dao.Db.Model(&models.Book{}).Count(&total)
	dao.Db.Offset((page - 1) * size).Limit(size).Find(&book)

	c.JSON(http.StatusOK, gin.H{
		"data": map[string]interface{}{
			"records": book,
			"current": page,
			"size":    size,
			"total":   total,
		},
	})

}

func DeleteById(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	err := dao.Db.Where("id = ?", id).Delete(&book)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"flag": true,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"flag": false,
		})
	}
}

func GetById(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	err := dao.Db.Where("id = ?", id).Find(&book)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"flag": true,
			"data": book,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"flag": false,
			"data": nil,
		})
	}
}

func UpdateByRow(c *gin.Context) {
	var book models.Book
	c.ShouldBind(&book)
	err := dao.Db.Save(&book)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"flag": true,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"flag": false,
		})
	}
}

func GetByCondition(c *gin.Context) { //还没能处理undefined和“”的处理
	var book = []models.Book{}
	booktype := c.DefaultQuery("type", "")
	bookname := c.DefaultQuery("name", "")
	bookdes := c.DefaultQuery("description", "")
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	var total int64
	dao.Db.Model(&models.Book{}).Count(&total)
	dao.Db.Offset((page - 1) * size).Limit(size).Where(fmt.Sprintf(" type like '%%%s%%' ", booktype)).Where(fmt.Sprintf(" name like '%%%s%%' ", bookname)).Where(fmt.Sprintf(" description like '%%%s%%' ", bookdes)).Find(&book)

	c.JSON(http.StatusOK, gin.H{
		"data": map[string]interface{}{
			"records": book,
			"current": page,
			"size":    size,
			"total":   total,
		},
	})
}

func DownloadFile(c *gin.Context) { //文件下载
	bookurl := c.Query("url")

	//获取文件的名称
	fileName := path.Base(bookurl)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")
	c.File(bookurl)

}
