package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple_tiktok/middlewares"
	"simple_tiktok/service"
	"simple_tiktok/utils"
)

func UserInfo(c *gin.Context) {
	tmp, _ := utils.GenerateToken(1, "merry")
	// 接受参数
	userId := c.DefaultQuery("user_id", "0")
	//token := c.DefaultQuery("token", "")

	// 验证token
	_, err := middlewares.AuthUserCheck(tmp)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": -1,
			"status_msg":  err.Error(),
		})
		return
	}

	// 布隆过滤器
	//if !utils.Filter.Check(userId) {
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"status_code": -1,
	//		"status_msg":  "no such user",
	//	})
	//	return
	//}

	// 把参数传给service层
	res, err := service.UserInfo(c, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": -1,
			"status_msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "string",
		"user":        res,
	})
}
