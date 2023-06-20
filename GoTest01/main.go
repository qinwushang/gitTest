package main

import (
	"GoTest01/controller"
	"GoTest01/dao"
	"GoTest01/models"
	"github.com/gin-gonic/gin"
)

func main() {
	dao.InitMySQL()
	dao.Db.AutoMigrate(&models.Book{})
	router := gin.Default()
	router.LoadHTMLFiles("./resources/static/Klibrary.html")
	router.Static("/static", "./resources/static")
	router.MaxMultipartMemory = 300 << 20

	router.GET("/KLibrary", controller.InitHandler)

	bookgroup := router.Group("/books")
	{
		bookgroup.PUT("", controller.UpdateByRow)
		bookgroup.GET("", controller.GetAll)
		bookgroup.POST("/bookadd", controller.UploadFile)
		bookgroup.GET("/:id", controller.GetById)
		bookgroup.DELETE("/:id", controller.DeleteById)
		bookgroup.GET("/condition", controller.GetByCondition)
		bookgroup.GET("/downloadfiles", controller.DownloadFile)
	}

	router.POST("/mail", controller.MailSender)
	router.POST("/messages", controller.FeedbackReceiver)

	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Run("127.0.0.1:8080")
}
