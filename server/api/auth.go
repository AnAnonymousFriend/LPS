package api

import "C"
import (
	"LNPS/server/models"
	"LNPS/server/pkg/app"
	"LNPS/server/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)



// @Summary 登录
// @Produce  json
// @Param userName path int true "userName"
// @Param passWord path int true "password"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/login [Post]
func GetAuth(ctx *gin.Context){
		userName := ctx.PostForm("userName")
		passWord := ctx.PostForm("passWord")
		islogin,err := models.Login(userName,passWord)
		if err != nil {
			println(err)
		}else{


		println(islogin)
		}

}


// @Summary 获取用户列表
// @Produce  json
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/user [Post]
func AddUser(ctx *gin.Context){
	g := app.Gin{ctx}
	userName := ctx.PostForm("userName")
	passWord := ctx.PostForm("passWord")
	isCreate := models.AddUser(userName,passWord)

	if isCreate {
		g.Response(http.StatusOK,e.SUCCESS,nil)
	}


	/*if isCreate {
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":   e.SUCCESS,
			"data": nil,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":   e.SUCCESS,
		"data": nil,
	})*/


}


