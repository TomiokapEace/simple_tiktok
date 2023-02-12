package controller

import (
	"net/http"
	"simple_tiktok/dao/mysql"
	"simple_tiktok/logger"
	"simple_tiktok/middlewares"
	"simple_tiktok/service"
	"simple_tiktok/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

/**
 * @Author jiang
 * @Description 粉丝列表接口
 * @Date 12:00 2023/2/12
 **/
// 返回结构体
type FollowerListRespStruct struct {
	Code     int            `json:"status_code"`
	Msg      string         `json:"status_msg"`
	UserList []service.User `json:"user_list"`
}

// 传入参数返回
func FollowerListResp(c *gin.Context, code int, msg string, userList []service.User) {
	h := &FollowerListRespStruct{
		Code:     code,
		Msg:      msg,
		UserList: userList,
	}

	c.JSON(http.StatusOK, h)
}

// GetFollowerList
// @Summary 粉丝列表
// @Tags 社交接口
// @Param token query string true "token"
// @Param user_id query string true "用户id"
// @Success 200 {object} FollowerListRespStruct
// @Router /relation/follower/list/ [get]
func GetFollowerList(c *gin.Context) {
	// 1、获取参数
	token := c.DefaultQuery("token", "")
	userId := c.DefaultQuery("user_id", "")

	// 2、验证参数
	if token == "" || userId == "" {
		FollowerListResp(c, -1, "请求参数错误", []service.User{})
		return
	}

	user_id, _ := strconv.Atoi(userId)
	if user_id == 0 {
		FollowerListResp(c, -1, "请求参数错误", []service.User{})
		return
	}

	cnt, err := mysql.FindUserByIdentityCount(uint64(user_id))
	if err != nil {
		logger.SugarLogger.Error("FindUserByIdentityCount Error：", err.Error())
		FollowerListResp(c, -1, "请求参数错误", []service.User{})
		return
	}
	if cnt == 0 {
		FollowerListResp(c, -1, "非法用户", []service.User{})
		return
	}

	// 3、验证token
	t, _ := utils.GenerateToken(1, "test")
	_, err = middlewares.AuthUserCheck(t)
	if err != nil {
		FollowerListResp(c, -1, "无效token", []service.User{})
		return
	}

	// 4、将数据传入service层
	data, err := service.FollowerListService(c, uint64(user_id))
	if err != nil {
		FollowerListResp(c, -1, "获取粉丝列表失败", []service.User{})
		return
	}

	FollowerListResp(c, 0, "获取粉丝列表成功", data)
}
