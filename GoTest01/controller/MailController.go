package controller

import (
	"GoTest01/models"
	"github.com/gin-gonic/gin"
)

func MailSender(c *gin.Context) {
	//var json models.MailInfo
	to := c.Query("to")
	subject := c.Query("subject")
	context := c.Query("context")
	bookUrl := c.Query("bookUrl")
	models.MailSender(to, subject, context, bookUrl)
}
