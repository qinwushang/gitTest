package controller

import (
	"GoTest01/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func FeedbackReceiver(c *gin.Context) {
	var feedback models.MsgFeedback
	c.ShouldBind(&feedback)
	fmt.Println("编号为：" + feedback.Id)
	fmt.Println("反馈为：" + feedback.Message)
}
